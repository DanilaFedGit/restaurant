package routes

import (
	"github.com/DanilaFedGit/restaurant/controllers"
	"github.com/gin-gonic/gin"
)

func FoodRoutes(router *gin.Engine) {
	router.GET("/food", controllers.GetFoods())
	router.GET("/food/:food_id", controllers.GetFood())
	router.POST("/food", controllers.CreateFood())
	router.PATCH("/food", controllers.UpdateFood())
}
