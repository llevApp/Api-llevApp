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
			user.GET("/trip-info/:id", func(c *gin.Context) {
				controllers_user.UserTripInfo(c, db)

			})
			user.POST("/new", func(c *gin.Context) {
				controllers_user.CreateUser(c, db)

			})
			user.GET("/university-career", func(c *gin.Context) {
				controllers_user.Carrers(c, db)

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
				controllers_trips.TripsDriver(c, db)
			})
			driver.GET("/active-trip/:id", func(c *gin.Context) {
				controllers_trips.ActiveTripDriver(c, db)
			})

		}
		passengers := api.Group("/passengers")
		{
			passengers.GET("/trips", func(c *gin.Context) {
				controllers_trips.ActiveTrips(c, db)
			})
		}
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
