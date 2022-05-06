package routes

import (
	"database/sql"
	controllers "llevapp/src/controllers/trips"

	"github.com/gin-gonic/gin"
)

func EndpointGroup(Engine *gin.Engine, db *sql.DB) error {

	api := Engine.Group("/trips")
	{
		driver := api.Group("/driver")
		{
			driver.GET("/:id", func(c *gin.Context) {
				//controllers.GetTrips(c)
			})
			driver.POST("/new-trip", func(c *gin.Context) {
				controllers.InsertNewTrip(c, db)
			})
			driver.PUT("/end-trip/:id", func(c *gin.Context) {
				controllers.EndTrip(c, db)
			})

		}
		passengers := api.Group("/passengers")
		{
			//get all active trips
			passengers.GET("/trips", func(c *gin.Context) {
				controllers.ActiveTrips(c, db)
			})
		}

		/* user := api.Group("/user")
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

		} */
	}
	return nil
}
