package main 

import (
	"os"
	"golang-restaurant-management/database"
	"golang-restaurant-management/routes"
	"golang-restaurant-management/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gin-gonic/gin"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	//fmt.Println("Hello")

	// Connect to the Database

	port:= os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router:= gin.New()

	router.Use(gin.Logger())

	//Initializing the routes

	routes.UserRoutes(router)

	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)


	router.Run(":" + port)
}