package controllers

import (
	"database/sql"
	"llevapp/src/models"
)

func newTripByDriver(db *sql.DB, trip models.NewTripsRecords) (err error) {
	insertDynStmt := `INSERT INTO llevapp.trips` +
		`(driver_user_id, init_longitude, init_latitude, init_time_utc, is_active)` +
		`VALUES($1, $2, $3, $4,TRUE)`

	_, err = db.Exec(insertDynStmt, trip.DriverId, trip.Longitude, trip.Latitude, trip.Time)
	if err != nil {
		return
	}
	return
}

func UpdateTripStatus(db *sql.DB, id int) (err error) {
	insertDynStmt := `UPDATE llevapp.trips ` +
		`SET is_active=false ` +
		`WHERE driver_user_id = $1`

	_, err = db.Exec(insertDynStmt, id)
	if err != nil {
		return
	}
	return
}

func GetActiveTrips(db *sql.DB) (ActiveTrips []models.TripsRecords, err error) {

	rows, err := db.Query(`SELECT u.name,c.name,t.init_longitude, t.init_latitude, t.init_time_utc ` +
		`FROM llevapp.trips as t ` +
		`INNER JOIN llevapp.users as u on u.id = t.driver_user_id ` +
		`INNER JOIN llevapp.career as c on c.id = u.career_id ` +
		`WHERE t.is_active = true `)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var Trips models.TripsRecords
		err = rows.Scan(&Trips.Driver, &Trips.DriverCareer, &Trips.Longitude, &Trips.Latitude, &Trips.InitTripTime)
		if err != nil {
			panic(err)
		}
		ActiveTrips = append(ActiveTrips, Trips)
	}

	return
}

func newTripRequest(db *sql.DB, trip models.TripRequest) (err error) {
	insertDynStmt := `INSERT INTO llevapp.trips_passenger ` +
		`(trip_id, passenger_user_id, longitude, latitude, contribution,is_valid,has_confirmation)` +
		`VALUES($1, $2, $3, $4,$5,TRUE,FALSE)`

	_, err = db.Exec(insertDynStmt, trip.TripID, trip.UserID, trip.Longitude, trip.Latitude, trip.Contribution)
	if err != nil {
		return
	}
	return
}

func AceptTripsRequest(db *sql.DB, tripId, passangerUserID int) (ActiveTrips []models.TripsRecords, err error) {

	insertDynStmt := `UPDATE llevapp.trips_passenger ` +
		`SET has_confirmation=true ` +
		`WHERE trip_id = $1 AND passenger_user_id = $2`

	_, err = db.Exec(insertDynStmt, tripId, passangerUserID)
	if err != nil {
		return
	}
	return
}

func DeclineTripsRequest(db *sql.DB, tripId, passangerUserID int) (ActiveTrips []models.TripsRecords, err error) {

	insertDynStmt := `UPDATE llevapp.trips_passenger ` +
		`SET has_confirmation=false, is_valid=false ` +
		`WHERE trip_id = $1 AND passenger_user_id = $2`
	_, err = db.Exec(insertDynStmt, tripId, passangerUserID)
	if err != nil {
		return
	}
	return
}

func GetRequestStatus(db *sql.DB, tripId, passangerUserID int) (state string, err error) {

	rows, err := db.Query(`SELECT is_valid,has_confirmation `+
		`FROM llevapp.trips_passenger as t `+
		`WHERE trip_id = $1 AND passenger_user_id = $2`, tripId, passangerUserID)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var is_valid, has_confirmation bool
		err = rows.Scan(&is_valid, &has_confirmation)
		if err != nil {
			panic(err)
		}

		if is_valid && has_confirmation {
			return "accepted", nil
		} else {
			if is_valid {
				return "waiting", nil
			} else {
				return "decline", nil
			}
		}
	}

	return
}

func InvalidTripsRequest(db *sql.DB, tripId int) (err error) {

	insertDynStmt := `UPDATE llevapp.trips_passenger ` +
		`SET  is_valid=false ` +
		`WHERE trip_id = $1 `
	_, err = db.Exec(insertDynStmt, tripId)
	if err != nil {
		return
	}
	return
}

func GetNewRequestTrips(db *sql.DB, tripId int) (ActiveRequest []models.TripRequest, err error) {

	rows, err := db.Query(`SELECT u.name,c.name,t.longitude, t.latitude, contribution `+
		`FROM llevapp.trips_passenger as t `+
		`INNER JOIN llevapp.users as u on u.id = t.passenger_user_id `+
		`INNER JOIN llevapp.career as c on c.id = u.career_id `+
		`WHERE t.is_valid = true AND t.trip_id= $1`, tripId)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var request models.TripRequest
		err = rows.Scan(&request.UserName, &request.UserCareer, &request.Longitude, &request.Latitude, &request.Contribution)
		if err != nil {
			panic(err)
		}
		ActiveRequest = append(ActiveRequest, request)
	}

	return
}
