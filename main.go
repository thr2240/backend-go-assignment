package main

import (
	"github.com/gin-gonic/gin"

	"github.com/thr2240.com/backend-go-assignment/initializers"
	"github.com/thr2240.com/backend-go-assignment/controllers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	app := gin.Default()
	
	// Endpoints for leads
	app.GET("/leads", controllers.GetAllLeads)
	app.GET("/leads/:id", controllers.GetLead)
	app.POST("/leads", controllers.CreateLead)
	app.DELETE("/leads/:id", controllers.DeleteLead)
	app.PATCH("/leads/:id", controllers.UpdateLead)

	app.Run()
}
