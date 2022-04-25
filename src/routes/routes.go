package routes

import (
	"database/sql"
	"llevapp/src/controllers"

	"github.com/gin-gonic/gin"
)

func EndpointGroup(Engine *gin.Engine, db *sql.DB) error {

	api := Engine.Group("/v1")
	{
		trips := api.Group("/trips")
		{
			trips.GET("/indicators/:id", func(c *gin.Context) {
				controllers.GetTrips(c)
			})

		}

		user := api.Group("/user")
		{
			user.POST("/info", func(c *gin.Context) {
				controllers.NewUser(c)
			})

		}

		users := api.Group("/users")
		{
			users.GET("/", func(c *gin.Context) {
				controllers.GetUsersDetail(db, c)
			})

		}
	}
	return nil
}
