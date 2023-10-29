package routes

import (
	"github.com/DanilaFedGit/restaurant/controllers"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.Engine) {
	router.GET("/orders", controllers.GetOrders())
	router.GET("/orders/:order_id", controllers.GetOrder())
	router.POST("/order", controllers.CreateOrder())
	router.PATCH("/order/:order_id", controllers.UpdateOrder())
}
