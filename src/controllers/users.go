package controllers

import (
	"database/sql"
	"llevapp/src/models"

	"github.com/gin-gonic/gin"
)

/*

SELECT id, email, first_name, surname, career_id, nick_name
FROM llevapp.users;


*/

func GetUsersDetail(b *sql.DB, c *gin.Context) {

	users := bdConsultAllUsers(b, c)

	if users != nil {
		c.JSON(200, users)
	} else {
		c.JSON(200, nil)
	}

}

func bdConsultAllUsers(db *sql.DB, c *gin.Context) (users []*models.User) {
	rows, err := db.Query("SELECT id, email, first_name, surname, career_id, nick_name FROM llevapp.users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var user *models.User
		err = rows.Scan(&user.UserID, &user.Email, &user.Name, &user.Surname, &user.CareerName, &user.Nickname)
		if err != nil {
			panic(err)
		} else {
			users = append(users, user)
		}
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return

}
