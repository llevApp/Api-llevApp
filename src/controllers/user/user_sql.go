package controllers

import (
	"database/sql"
	"llevapp/src/models"
)

func GetUserInfo(db *sql.DB, UserEmail string) (User models.User, err error) {

	rows, err := db.Query(`SELECT u.id,u.email,u.name,c.name,u.nick_name,u.uuid_fb `+
		`FROM llevapp.users as u `+
		`INNER JOIN llevapp.career as c on c.id = u.career_id `+
		`WHERE u.email = $1 `, UserEmail)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&User.UserID, &User.Email, &User.Name, &User.CareerName, &User.Nickname, &User.UUID)
		if err != nil {
			panic(err)
		}
	}

	return
}
func GetUserTripInfo(db *sql.DB, Id string) (Trip []models.TripRequest, err error) {
	rows, err := db.Query(`SELECT u2.name,u2.id,tp.location,tp.contribution,u2.uuid_fb `+
		`FROM llevapp.trips_passenger as tp `+
		`INNER JOIN llevapp.users as u2 on u2.id = tp.passenger_user_id `+
		`WHERE tp.trip_id = $1 `, Id)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var request models.TripRequest
		err = rows.Scan(&request.UserName, &request.Location, &request.Contribution, &request.UUID)
		if err != nil {
			panic(err)
		}
		Trip = append(Trip, request)
	}
	return
}

func GetUserInfoById(db *sql.DB, UserID int) (User models.User, err error) {

	rows, err := db.Query(`SELECT u.id,u.email,u.name,c.name,u.nick_name,uuid_fb `+
		`FROM llevapp.users as u `+
		`INNER JOIN llevapp.career as c on c.id = u.career_id `+
		`WHERE u.id = $1 `, UserID)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&User.UserID, &User.Email, &User.Name, &User.CareerName, &User.Nickname, &User.UUID)
		if err != nil {
			panic(err)
		}
	}

	return
}

func CreateNewUser(db *sql.DB, User models.User) (err error) {

	insertDynStmt := `INSERT INTO llevapp.users` +
		`(email, name, career_id, nick_name, uuid_fb)` +
		`VALUES($1, $2, $3, $4,$5)`

	_, err = db.Exec(insertDynStmt, User.Email, User.Name, User.CareerId, User.Nickname, User.UUID)
	if err != nil {
		return
	}
	return

}

func GetCarrer(db *sql.DB) (carrers []models.Carrer, err error) {
	rows, err := db.Query(`SELECT id, name ` +
		`FROM llevapp.career as tp `)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var carrer models.Carrer
		err = rows.Scan(&carrer.Id, &carrer.Name)
		if err != nil {
			panic(err)
		}
		carrers = append(carrers, carrer)
	}
	return
}

/* SELECT id, name
FROM llevapp.career;
*/
