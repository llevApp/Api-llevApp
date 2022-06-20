package routes

import (
	"database/sql"
	controllers_trips "llevapp/src/controllers/trips"
	controllers_user "llevapp/src/controllers/user"
	ws_chat "llevapp/src/websocket/chat"
	ws_location "llevapp/src/websocket/location"
	ws_request "llevapp/src/websocket/trip_request"

	"github.com/gin-gonic/gin"
)

func EndpointGroup(Engine *gin.Engine, db *sql.DB, hub_request *ws_request.Hub, hub_chat *ws_chat.Hub, hub_location *ws_location.Hub) error {

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

			/* passengers.POST("/trip-request", func(c *gin.Context) {
				controllers_trips.TripRequest(c, db)
			})

			passengers.GET("/request-state", func(c *gin.Context) {
				controllers_trips.RequestState(c, db)
			}) */
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

	ws := Engine.Group("/ws")
	{
		ws.GET("/trip-request/:driverId", func(c *gin.Context) {
			UserRoom := c.Param("driverId")
			ws_request.ServeWs(c.Writer, c.Request, UserRoom, db, hub_request)
		})
		ws.GET("/chat/:driverId/:passengerId", func(c *gin.Context) {
			DriverId := c.Param("driverId")
			PassengerId := c.Param("passengerId")

			UserRoom := DriverId + "_" + PassengerId
			ws_chat.ServeWs(c.Writer, c.Request, UserRoom, hub_chat)
		})
		ws.GET("/location/:driverId", func(c *gin.Context) {
			UserRoom := c.Param("driverId")
			ws_location.ServeWs(c.Writer, c.Request, UserRoom, db, hub_location)
		})
	}
	return nil
}
