package routes

import (
	"github.com/DanilaFedGit/restaurant/controllers"
	"github.com/gin-gonic/gin"
)

func MenuRoutes(router *gin.Engine) {
	router.GET("/menus", controllers.GetMenus())
	router.GET("/menus/:menu_id", controllers.GetMenu())
	router.POST("/menus", controllers.CreateMenu())
	router.PATCH("/menus/:menu_id", controllers.UpdateMenu())
}
