package controllers

import (
	"llevapp/src/models"

	"github.com/gin-gonic/gin"
)

func GetTrips(c *gin.Context) {
	var (
		user models.User
	)
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
	if err := c.ShouldBindJSON(&user); err == nil {
		c.JSON(200, user)
	} else {
		c.JSON(204, "No data")
	}

}
