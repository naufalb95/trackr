package service

import (
	"context"
	"strings"

	"github.com/naufalb95/trackr/internal/dto"
	"github.com/naufalb95/trackr/internal/model"
	"github.com/naufalb95/trackr/internal/repository"
)

type TaskService interface {
	GetAllTasks(ctx context.Context) ([]dto.GetAllTasksResponse, error)
	GetSingleTask(ctx context.Context, taskId string) (*dto.GetTaskByIdResponse, error)
	CreateTask(ctx context.Context, task *dto.CreateTaskRequest) (*dto.CreateTaskResponse, error)
	UpdateTask(ctx context.Context, taskId string, updatedFields dto.UpdateTaskRequest) error
	DeleteTask(ctx context.Context, taskId string) error
}

type taskService struct {
	taskRepo repository.TaskRepository
}

func NewTaskService(taskRepo repository.TaskRepository) TaskService {
	return &taskService{taskRepo: taskRepo}
}

func (s *taskService) GetAllTasks(ctx context.Context) ([]dto.GetAllTasksResponse, error) {
	tasks, err := s.taskRepo.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	tasksResponse := make([]dto.GetAllTasksResponse, 0, len(tasks))

	for _, task := range tasks {
		task := dto.GetAllTasksResponse{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
			CreatedAt:   task.CreatedAt,
			UpdatedAt:   task.UpdatedAt,
		}

		tasksResponse = append(tasksResponse, task)
	}

	return tasksResponse, nil
}

func (s *taskService) GetSingleTask(ctx context.Context, taskId string) (*dto.GetTaskByIdResponse, error) {
	task, err := s.taskRepo.FindById(ctx, taskId)

	if err != nil {
		return nil, err
	}

	taskResponse := &dto.GetTaskByIdResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}

	return taskResponse, nil
}

func (s *taskService) CreateTask(ctx context.Context, request *dto.CreateTaskRequest) (*dto.CreateTaskResponse, error) {
	task := &model.Task{
		Title:       strings.TrimSpace(request.Title),
		Description: strings.TrimSpace(request.Description),
		Status:      request.Status,
	}

	err := s.taskRepo.Create(ctx, task)

	if err != nil {
		return nil, err
	}

	response := &dto.CreateTaskResponse{
		ID: task.ID,
	}

	return response, nil
}

func (s *taskService) UpdateTask(ctx context.Context, taskId string, updatedFields dto.UpdateTaskRequest) error {
	updates := make(map[string]any)

	if updatedFields.Title != nil {
		title := strings.TrimSpace(*updatedFields.Title)
		updates["title"] = title
	}

	if updatedFields.Description != nil {
		description := strings.TrimSpace(*updatedFields.Description)
		updates["description"] = description
	}

	if updatedFields.Status != nil {
		updates["status"] = updatedFields.Status
	}

	err := s.taskRepo.UpdateStatus(ctx, taskId, updates)

	if err != nil {
		return err
	}

	return nil
}

func (s *taskService) DeleteTask(ctx context.Context, taskId string) error {
	err := s.taskRepo.Delete(ctx, taskId)

	if err != nil {
		return err
	}

	return nil
}
