package routers

import (
	"apiProject/go-gin/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "This is my homepage!",
		})
	})
	r.POST("/gps", getGPS)
	r.POST("/update", updateGPS)
	return r
}

func getGPS(c *gin.Context) {
	orderVehicleId := c.PostForm("order_vehicle_id")
	diriverId := c.PostForm("driver_id")
	order_vehicle_id, _ := strconv.Atoi(orderVehicleId)
	driver_id, _ := strconv.Atoi(diriverId)
	listGPS := models.GetGPS(order_vehicle_id, driver_id)
	b, err := json.Marshal(listGPS)
	if err != nil {
		panic(err.Error())
	}
	c.Data(200, "application/json", b)
}
func updateGPS(c *gin.Context) {
	var JSON models.UpdateJSON
	c.Bind(&JSON)
	listGPS := models.UpdateGPS(JSON.OrderVehicleID, JSON.DriverID, JSON.BodyGps)
	b, _ := json.Marshal(listGPS)
	c.Data(200, "application/json", b)
}
