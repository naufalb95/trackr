package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	model "github.com/naufalb95/trackr/internal/models"
	repository "github.com/naufalb95/trackr/internal/repositories"
)

type TaskHandler struct {
	repo repository.TaskRepository
}

func NewTaskHandler(repo repository.TaskRepository) *TaskHandler {
	return &TaskHandler{repo: repo}
}

func (h *TaskHandler) GetTasks(c *gin.Context) {
	tasks, err := h.repo.FindAll(context.Background())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var newTask model.Task

	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}

	// newTask.ID = len(tasks) + 1
	// tasks = append(tasks, newTask)

	c.JSON(http.StatusCreated, newTask)
}

func (h *TaskHandler) UpdateTaskStatus(c *gin.Context) {
	var updatedTask model.TaskUpdateDTO

	if err := c.BindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}

	// taskIdParam := c.Param("id")
	// taskId, err := strconv.Atoi(taskIdParam)

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID parameter"})
	// 	return
	// }

	// for i := range tasks {
	// 	if tasks[i].ID == taskId {
	// 		tasks[i].Status = updatedTask.Status
	// 		break
	// 	}
	// }

	c.JSON(http.StatusOK, updatedTask)
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	// taskIdParam := c.Param("id")
	// taskId, err := strconv.Atoi(taskIdParam)

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID parameter"})
	// 	return
	// }

	// taskIndex := slices.IndexFunc(tasks, func(task model.Task) bool {
	// 	return task.ID == taskId
	// })

	// tasks = slices.Delete(tasks, taskIndex, taskIndex+1)
}
