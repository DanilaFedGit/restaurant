package routes

import (
	"github.com/DanilaFedGit/restaurant/controllers"
	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(router *gin.Engine) {
	router.GET("/orderItems", controllers.GetOrderItems())
	router.GET("/orderItems/:orderItem_id", controllers.GetOrderItem())
	router.GET("/orderItems-order/:order_id", controllers.GetOrderItemByOrder())
	router.POST("/orderItems", controllers.CreateOrderItem())
	router.PATCH("/orderItems/:orderItem_id", controllers.UpdateOrderIteme())
}
