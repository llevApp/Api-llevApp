package controllers

/* func GetTripsByDriver(db *sql.DB, id int) (trips []models.TripsRecords) {
	rows, err :=
		db.Query(
			`SELECT passeger.first_name,tp.longitude, tp.latitude, tp.contribution`+
				`FROM trips as t`+
				`INNER JOIN user as u on u.id = t.driver_user_id `+
				`INNER JOIN trips_passeger as tp on t.id = tp.trip_id `+
				`INNER JOIN user as passeger on passeger.id = tp.passeger_user_id `+
				`WHERE u.id = $1`, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var trip models.TripsRecords
		err = rows.Scan(&trip.Passager, &trip.Longitude, &trip.Latitude, &trip.Contribution)
		if err != nil {
			panic(err)
		}
		trips = append(trips, trip)

		err = rows.Err()
		if err != nil {
			panic(err)
		}
	}
	return
}
*/
