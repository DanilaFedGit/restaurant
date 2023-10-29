package routes

import (
	"github.com/DanilaFedGit/restaurant/controllers"
	"github.com/gin-gonic/gin"
)

func TableRoutes(router *gin.Engine) {
	router.GET("/tables", controllers.GetTables())
	router.GET("/tables/:table_id", controllers.GetTable())
	router.POST("/tables", controllers.CreateTable())
	router.PATCH("/tables/:table_id", controllers.UpdateTable())
}
