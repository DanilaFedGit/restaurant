package routes

import (
	"github.com/DanilaFedGit/restaurant/controllers"
	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(router *gin.Engine) {
	router.GET("/invoices", controllers.GetInvoices())
	router.GET("/invoice/:invoice_id", controllers.GetInvoice())
	router.POST("/invoice", controllers.CreateInvoice())
	router.PATCH("/invoice/:invoice_id", controllers.UpdateInvoic())
}
