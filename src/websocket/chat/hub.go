package websocket_chat

type Message struct {
	data []byte
	room string
}

type Subscription struct {
	conn *Connection
	room string
}

// Hub maintains the set of active Connections and broadcasts Messages to the
// Connections.
type Hub struct {
	// Registered Connections.
	rooms map[string]map[*Connection]bool

	// Inbound Messages from the Connections.
	broadcast chan Message

	// Register requests from the Connections.
	register chan Subscription

	// Unregister requests from Connections.
	unregister chan Subscription
}

var h = Hub{
	broadcast:  make(chan Message),
	register:   make(chan Subscription),
	unregister: make(chan Subscription),
	rooms:      make(map[string]map[*Connection]bool),
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan Message),
		register:   make(chan Subscription),
		unregister: make(chan Subscription),
		rooms:      make(map[string]map[*Connection]bool),
	}
}
func (h *Hub) Run() {
	for {
		select {
		case s := <-h.register:
			Connections := h.rooms[s.room]
			if Connections == nil {
				Connections = make(map[*Connection]bool)
				h.rooms[s.room] = Connections
			}
			h.rooms[s.room][s.conn] = true
		case s := <-h.unregister:
			Connections := h.rooms[s.room]
			if Connections != nil {
				if _, ok := Connections[s.conn]; ok {
					delete(Connections, s.conn)
					close(s.conn.send)
					if len(Connections) == 0 {
						delete(h.rooms, s.room)
					}
				}
			}
		case m := <-h.broadcast:
			Connections := h.rooms[m.room]
			for c := range Connections {
				select {
				case c.send <- m.data:
				default:
					close(c.send)
					delete(Connections, c)
					if len(Connections) == 0 {
						delete(h.rooms, m.room)
					}
				}
			}
		}
	}
}
