package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/naufalb95/trackr/internal/handlers"
	"github.com/naufalb95/trackr/internal/middlewares"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// CORS Middleware
	router.Use(middlewares.CORS())

	// API Routers
	api := router.Group("/api")
	{
		api.GET("/tasks", handlers.GetTasks)
		api.POST("/tasks", handlers.CreateTask)
		api.DELETE("/tasks/:id", handlers.DeleteTask)
	}

	return router
}
