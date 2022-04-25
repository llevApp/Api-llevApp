package controllers

import (
	"llevapp/src/models"

	"github.com/gin-gonic/gin"
)

func GetTrips(c *gin.Context) {
	var (
		user models.User
	)

	id := c.Param("id")
	if id == "1" {
		user.Name = "Diosito"
	}
	if user.Name != "" {
		c.JSON(200, user)

	} else {
		c.JSON(204, "No data")

	}

}

func NewUser(c *gin.Context) {
	var (
		user models.User
	)

	/* name := c.PostForm("name")
	email := c.PostForm("email") */

	if err := c.ShouldBindJSON(&user); err == nil {
		c.JSON(200, user)
	} else {
		c.JSON(204, "No data")
	}

}
