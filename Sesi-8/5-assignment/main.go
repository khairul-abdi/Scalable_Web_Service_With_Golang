package main

import (
	"assignment_2/config"
	"assignment_2/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.GET("/orders", inDB.GetOrders)
	router.GET("/order/:orderID", inDB.GetOrder)
	router.POST("/orders", inDB.CreateOrder)
	router.PUT("/orders/:orderID", inDB.UpdateOrder)
	router.DELETE("/orders/:orderID", inDB.DeleteOrder)

	router.Run(":8080")

}
