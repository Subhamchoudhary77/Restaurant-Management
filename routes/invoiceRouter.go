// main->router->controller->models->database     - This is the architecture layer

package routes

import (
	controller "golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoices", controller.GetInvoices())
	incomingRoutes.GET("/invoices/:invoice_id", controller.GetInvoice())
	incomingRoutes.GET("/invoices", controller.CreateInvoice())
	incomingRoutes.GET("/invoices/:invoice_id", controller.UpdateInvoice())
}
