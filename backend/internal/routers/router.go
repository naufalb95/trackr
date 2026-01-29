package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/naufalb95/trackr/internal/handlers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// CORS Middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// API Routers
	api := router.Group("/api")
	{
		api.GET("/tasks", handlers.GetTasks)
		api.POST("/tasks", handlers.CreateTask)
	}

	return router
}
