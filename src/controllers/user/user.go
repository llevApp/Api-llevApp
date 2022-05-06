package controllers

import (
	"database/sql"

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
