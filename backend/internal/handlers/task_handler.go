package handlers

import (
	"net/http"
	"slices"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/naufalb95/trackr/internal/models"
)

var tasks = []models.Task{
	{ID: 1, Title: "Learn Go basics", Description: "Complete Go tour", Status: "in_progress"},
	{ID: 2, Title: "Build Trackr API", Description: "Create REST endpoints", Status: "todo"},
	{ID: 3, Title: "Learn React hooks", Description: "Refresh React knowledge", Status: "todo"},
}

func GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func CreateTask(c *gin.Context) {
	var newTask models.Task

	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}

	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	c.JSON(http.StatusCreated, newTask)
}

func UpdateTaskStatus(c *gin.Context) {
	var updatedTask models.TaskUpdateDTO

	if err := c.BindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}

	taskIdParam := c.Param("id")
	taskId, err := strconv.Atoi(taskIdParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID parameter"})
		return
	}

	for i := range tasks {
		if tasks[i].ID == taskId {
			tasks[i].Status = updatedTask.Status
			break
		}
	}

	c.JSON(http.StatusOK, updatedTask)
}

func DeleteTask(c *gin.Context) {
	taskIdParam := c.Param("id")
	taskId, err := strconv.Atoi(taskIdParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID parameter"})
		return
	}

	taskIndex := slices.IndexFunc(tasks, func(task models.Task) bool {
		return task.ID == taskId
	})

	tasks = slices.Delete(tasks, taskIndex, taskIndex+1)
}
