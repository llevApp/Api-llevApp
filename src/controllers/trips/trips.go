package controllers

import (
	"database/sql"
	"llevapp/src/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InsertNewTrip(c *gin.Context, db *sql.DB) {
	var (
		trip models.NewTripsRecords
	)

	error_insert := newTripByDriver(db, trip)
	if error_insert == nil {
		c.JSON(200, "Trips save succesfully")
	} else {
		c.JSON(204, "No data")

	}

}

func EndTrip(c *gin.Context, db *sql.DB) {

	userId := c.Param("id")
	UserId, err := strconv.Atoi(userId)
	if err == nil {
		error_update := UpdateTripStatus(db, UserId)
		if error_update == nil {
			c.JSON(200, "Trips save succesfully")
		} else {
			c.JSON(204, "No data")

		}
	}

}

func ActiveTrips(c *gin.Context, db *sql.DB) {

	trips, err := GetActiveTrips(db)
	if err == nil {
		c.JSON(200, trips)
	} else {
		c.JSON(204, "No data")

	}
}
