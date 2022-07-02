package controllers

import (
	"database/sql"
	"llevapp/src/models"

	"github.com/gin-gonic/gin"
)

func UserInfo(c *gin.Context, db *sql.DB) {

	UserEmail := c.Param("email")
	UserInfo, err := GetUserInfo(db, UserEmail)
	if err == nil {
		c.JSON(200, UserInfo)
	} else {
		c.JSON(204, err.Error())

	}

}
func UserTripInfo(c *gin.Context, db *sql.DB) {
	Id := c.Param("id")
	TripInfo, err := GetUserTripInfo(db, Id)
	if err == nil {
		c.JSON(200, TripInfo)
	} else {
		c.JSON(204, err.Error())
	}
}

func CreateUser(c *gin.Context, db *sql.DB) {
	var (
		user models.User
	)

	if err := c.ShouldBindJSON(&user); err == nil {
		err := CreateNewUser(db, user)
		if err == nil {
			response := "user: " + user.Name + " created successfully"
			c.JSON(200, response)
		} else {
			c.JSON(203, err.Error())
		}
	} else {
		c.JSON(203, err.Error())
	}
}

func Carrers(c *gin.Context, db *sql.DB) {
	carrers, err := GetCarrer(db)
	if err == nil {
		c.JSON(200, carrers)
	} else {
		c.JSON(204, err.Error())
	}
}
