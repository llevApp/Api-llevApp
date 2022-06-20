// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package websocket_location

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"llevapp/src/models"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a msg to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong msg from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum msg size allowed from peer.
	maxmsgSize = 512
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// connection is an middleman between the websocket connection and the hub.
type Connection struct {
	// The websocket connection.
	ws *websocket.Conn

	// Buffered channel of outbound msgs.
	send chan []byte
}

// readPump pumps msgs from the websocket connection to the hub.
func (s Subscription) readPump(db *sql.DB, hub *Hub) {
	var (
		location models.LocationResponse
	)

	c := s.conn
	defer func() {
		hub.unregister <- s
		c.ws.Close()
	}()
	c.ws.SetReadLimit(maxmsgSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, msg, err := c.ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		msg = bytes.TrimSpace(bytes.Replace(msg, newline, space, -1))
		err = json.Unmarshal(msg, &location)

		if err != nil || location.Location.Latitude == 0 || location.Location.Longitude == 0 {
			if location.Location.Latitude == 0 || location.Location.Longitude == 0 {
				err = errors.New("location")
			}
			errorType := (strings.Split(err.Error(), ":"))[0]
			str := "Error: " + err.Error()
			switch errorType {
			case "pq":
				log.Printf("error: %v", err)

				str := "[SERVER]: Data base Error,please review the selected trip, it is possible that this is no longer available "
				errorMsg := []byte(str)
				m := Message{errorMsg, s.room}
				hub.broadcast <- m
			case "location":
				log.Printf("error: %v", err)

				str := "[LOCATION]: Location Error,please review the location or JSON struct, it is possible that contain errors "
				errorMsg := []byte(str)
				m := Message{errorMsg, s.room}
				hub.broadcast <- m
			default:

				errorMsg := []byte(str)
				m := Message{errorMsg, s.room}
				hub.broadcast <- m
			}

			/* if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break */
		} else {
			m := Message{msg, s.room}
			hub.broadcast <- m
		}

	}
}

// write writes a msg with the given msg type and payload.
func (c *Connection) write(mt int, payload []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(mt, payload)
}

// writePump pumps msgs from the hub to the websocket connection.
func (s *Subscription) writePump(db *sql.DB) {
	c := s.conn
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()
	for {
		select {
		case msg, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.write(websocket.TextMessage, msg); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func ServeWs(w http.ResponseWriter, r *http.Request, UserRoom string, db *sql.DB, hub *Hub) {

	ws, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	c := &Connection{send: make(chan []byte, 256), ws: ws}
	s := Subscription{c, UserRoom}
	hub.register <- s
	go s.writePump(db)
	go s.readPump(db, hub)
}
