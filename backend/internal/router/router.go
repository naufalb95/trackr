package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/naufalb95/trackr/internal/handler"
	"github.com/naufalb95/trackr/internal/middleware"
	"github.com/naufalb95/trackr/internal/repository"
	"github.com/naufalb95/trackr/internal/service"
)

func SetupRouter(pool *pgxpool.Pool) *gin.Engine {
	router := gin.Default()

	// CORS Middleware
	router.Use(middleware.CORS())

	// Repository
	taskRepo := repository.NewPostgresTaskRepository(pool)

	// Task Service
	taskService := service.NewTaskService(taskRepo)

	// Task Handler
	taskHandler := handler.NewTaskHandler(taskService)

	// API Routers
	api := router.Group("/api")
	{
		api.GET("/tasks", taskHandler.GetTasks)
		api.GET("/tasks/:id", taskHandler.GetTaskById)
		api.POST("/tasks", taskHandler.CreateTask)
		api.DELETE("/tasks/:id", taskHandler.DeleteTask)
		api.PATCH("/tasks/:id", taskHandler.UpdateTaskStatus)
	}

	return router
}
