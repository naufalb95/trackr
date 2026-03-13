package dto

import "time"

// Get All Tasks
type GetAllTasksResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Get Task by ID
type GetTaskByIdResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Create Task
type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required,min=3,max=100"`
	Description string `json:"description" binding:"max=500"`
	Status      string `json:"status" binding:"required,oneof=todo in_progress done"`
}

type CreateTaskResponse struct {
	ID string `json:"id"`
}

// Update Task
type UpdateTaskRequest struct {
	Title       *string `json:"title,omitempty" binding:"omitempty,min=3,max=100"`
	Description *string `json:"description,omitempty" binding:"omitempty,max=500"`
	Status      *string `json:"status,omitempty" binding:"omitempty,oneof=todo in_progress done"`
}
