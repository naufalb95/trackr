package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/naufalb95/trackr/internal/dto"
	"github.com/naufalb95/trackr/internal/model"
	"github.com/naufalb95/trackr/internal/service"
)

type TaskHandler struct {
	service service.TaskService
}

func NewTaskHandler(service service.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) GetTasks(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	tasks, err := h.service.GetAllTasks(ctx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func (h *TaskHandler) GetTaskById(c *gin.Context) {
	taskId := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	task, err := h.service.GetSingleTask(ctx, taskId)

	if err != nil {
		fmt.Printf("Error when trying to retrieve task by ID: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var newTask model.Task

	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	_, err := h.service.CreateTask(ctx, &newTask)

	if err != nil {
		fmt.Printf("Error when creating task: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error occurred when creating task."})
		return
	}

	c.JSON(http.StatusCreated, newTask)
}

func (h *TaskHandler) UpdateTaskStatus(c *gin.Context) {
	var updatedTask dto.UpdateTaskDTO

	if err := c.BindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}

	taskId := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)

	defer cancel()

	err := h.service.UpdateTask(ctx, taskId, updatedTask)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedTask)
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	taskId := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err := h.service.DeleteTask(ctx, taskId)

	if err != nil {
		fmt.Printf("Error when creating task: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error occurred when creating task."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
