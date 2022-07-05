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

	if err := c.ShouldBindJSON(&trip); err == nil {
		error_insert := newTripByDriver(db, trip)
		if error_insert == nil {
			c.JSON(200, "Trips save succesfully")
		} else {
			c.JSON(204, error_insert.Error())

		}
	}

}

func EndTrip(c *gin.Context, db *sql.DB) {

	tripId := c.Param("id")
	TripID, err := strconv.Atoi(tripId)
	if err == nil {
		error_update := UpdateTripStatus(db, TripID)
		err := InvalidTripsRequest(db, TripID)
		if error_update == nil && err == nil {
			c.JSON(200, "Trips save succesfully")
		} else {
			c.JSON(204, error_update.Error())

		}
	}

}

func ActiveTrips(c *gin.Context, db *sql.DB) {

	trips, err := GetActiveTrips(db)
	if err == nil {
		c.JSON(200, trips)
	} else {
		c.JSON(204, err.Error())

	}
}

func TripsDriver(c *gin.Context, db *sql.DB) {
	userId := c.Param("id")
	trips, err := GetTripsDriver(db, userId)
	for i := range trips.Trips {
		trips.Trips[i].TotalPassenger, _ = GetTotalPassenger(db, trips.Trips[i].Id)
		trips.Trips[i].TotalTip, _ = GetTotalTips(db, trips.Trips[i].Id)
	}
	if err == nil {
		c.JSON(200, trips)
	} else {
		c.JSON(204, err.Error())

	}
}

func ActiveTripDriver(c *gin.Context, db *sql.DB) {
	userId := c.Param("id")
	trips, err := GetActiveTripsDriver(db, userId)

	if err == nil {
		c.JSON(200, trips)
	} else {
		c.JSON(204, err.Error())

	}
}

func TripRequest(c *gin.Context, db *sql.DB) {
	var (
		trip models.TripRequest
	)

	if err := c.ShouldBindJSON(&trip); err == nil {
		error_insert := NewTripRequest(db, trip)
		if error_insert == nil {
			c.JSON(200, "Trips save succesfully")
		} else {
			c.JSON(204, error_insert.Error())

		}
	}

}

func TripRequestDriver(c *gin.Context, db *sql.DB) {

	response := c.Param("response")
	trip_id := c.Query("trip")
	passangerUserID := c.Query("user")
	TripId, err := strconv.Atoi(trip_id)
	PassangerUserID, err := strconv.Atoi(passangerUserID)

	if response == "accepted" {
		if err == nil {
			error_update, _ := AceptTripsRequest(db, TripId, PassangerUserID)
			if error_update == nil {
				c.JSON(200, "Request save succesfully")
			}
		}
	} else {
		if err == nil {
			error_update, _ := DeclineTripsRequest(db, TripId, PassangerUserID)
			if error_update == nil {
				c.JSON(200, "Request save succesfully")
			}
		}
	}

}

func RequestState(c *gin.Context, db *sql.DB) {
	trip_id := c.Query("trip")
	passangerUserID := c.Query("user")
	TripId, err := strconv.Atoi(trip_id)
	PassangerUserID, err := strconv.Atoi(passangerUserID)
	trips, err := GetRequestStatus(db, TripId, PassangerUserID)
	if err == nil {
		c.JSON(200, trips)
	} else {
		c.JSON(204, err.Error())

	}
}

func GetRequestDriver(c *gin.Context, db *sql.DB) {

	tripId := c.Query("trip")
	TripID, err := strconv.Atoi(tripId)
	if err == nil {
		getActiveRequest, err := GetNewRequestTrips(db, TripID)
		if err == nil {
			c.JSON(200, getActiveRequest)
		} else {
			c.JSON(204, err.Error())

		}
	}

}
