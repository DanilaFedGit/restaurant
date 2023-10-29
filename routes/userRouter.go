package routes

import (
	"github.com/DanilaFedGit/restaurant/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.GET("/user", controllers.GetUsers())
	router.GET("/user/:id", controllers.GetUser())
	router.POST("/user/signup", controllers.Signup())
	router.POST("/user/login", controllers.Login())
}
