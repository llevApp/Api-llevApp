package controllers

import (
	"database/sql"
	"llevapp/src/models"
)

func GetUserInfo(db *sql.DB, UserEmail string) (User models.User, err error) {

	rows, err := db.Query(`SELECT u.id,u.email,u.name,c.name,u.nick_name `+
		`FROM llevapp.users as u `+
		`INNER JOIN llevapp.career as c on c.id = u.career_id `+
		`WHERE u.email = $1 `, UserEmail)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&User.UserID, &User.Email, &User.Name, &User.CareerName, &User.Nickname)
		if err != nil {
			panic(err)
		}
	}

	return
}
