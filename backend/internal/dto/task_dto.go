package dto

type CreateTaskDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type UpdateTaskDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
