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

	rows, err := db.Query(`SELECT u.name,c.name,t.init_longitude, t.init_latitude. t.init_time_utc ` +
		`FROM llevapp.trips as t ` +
		`INNER JOIN llevapp.users as u on u.id = r.driver_user_id ` +
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
