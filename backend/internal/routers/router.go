package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/naufalb95/trackr/internal/handlers"
	"github.com/naufalb95/trackr/internal/middlewares"
	repository "github.com/naufalb95/trackr/internal/repositories"
)

func SetupRouter(pool *pgxpool.Pool) *gin.Engine {
	router := gin.Default()

	// CORS Middleware
	router.Use(middlewares.CORS())

	// Repository
	taskRepo := repository.NewPostgresTaskRepository(pool)

	// Task Handler
	taskHandler := handlers.NewTaskHandler(taskRepo)

	// API Routers
	api := router.Group("/api")
	{
		api.GET("/tasks", taskHandler.GetTasks)
		api.POST("/tasks", taskHandler.CreateTask)
		api.DELETE("/tasks/:id", taskHandler.DeleteTask)
		api.PATCH("/tasks/:id", taskHandler.UpdateTaskStatus)
	}

	return router
}
