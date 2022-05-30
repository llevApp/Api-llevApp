package routes

import (
	"database/sql"
	controllers_trips "llevapp/src/controllers/trips"
	controllers_user "llevapp/src/controllers/user"
	"llevapp/src/websocket"

	"github.com/gin-gonic/gin"
)

func EndpointGroup(Engine *gin.Engine, db *sql.DB, hub *websocket.Hub) error {

	api := Engine.Group("/api-llevapp")
	{
		user := api.Group("/user")
		{
			user.GET("/:email", func(c *gin.Context) {
				controllers_user.UserInfo(c, db)
			})
		}
		driver := api.Group("/driver")
		{
			driver.GET("/trips-request", func(c *gin.Context) {
				controllers_trips.GetRequestDriver(c, db)
			})
			driver.POST("/new-trip", func(c *gin.Context) {
				controllers_trips.InsertNewTrip(c, db)
			})
			driver.PUT("/end-trip/:id", func(c *gin.Context) {
				controllers_trips.EndTrip(c, db)
			})
			driver.PUT("/trip-request/:response", func(c *gin.Context) {
				controllers_trips.TripRequestDriver(c, db)
			})
			driver.GET("/trips/:id", func(c *gin.Context) {
				controllers_trips.ActiveTripsDriver(c, db)
			})

		}
		passengers := api.Group("/passengers")
		{
			//get all active trips
			passengers.GET("/trips", func(c *gin.Context) {
				controllers_trips.ActiveTrips(c, db)
			})

			passengers.POST("/trip-request", func(c *gin.Context) {
				controllers_trips.TripRequest(c, db)
			})

			passengers.GET("/request-state", func(c *gin.Context) {
				controllers_trips.RequestState(c, db)
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

	ws := Engine.Group("/websocket")
	{
		ws.GET("/request", func(c *gin.Context) {
			websocket.ServeWs(hub, c.Writer, c.Request, db)
		})
	}
	return nil
}
