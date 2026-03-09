package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/naufalb95/trackr/internal/model"
	"github.com/naufalb95/trackr/internal/repository"
)

type TaskService interface {
	GetAllTasks(ctx context.Context) ([]model.Task, error)
	GetSingleTask(ctx context.Context, taskId string) (*model.Task, error)
	CreateTask(ctx context.Context, task *model.Task) (*model.Task, error)
	UpdateTask(ctx context.Context, taskId string, taskStatus string) error
	DeleteTask(ctx context.Context, taskId string) error
}

type taskService struct {
	taskRepo repository.TaskRepository
}

func NewTaskService(taskRepo repository.TaskRepository) TaskService {
	return &taskService{taskRepo: taskRepo}
}

func (s *taskService) GetAllTasks(ctx context.Context) ([]model.Task, error) {
	tasks, err := s.taskRepo.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *taskService) GetSingleTask(ctx context.Context, taskId string) (*model.Task, error) {
	task, err := s.taskRepo.FindById(ctx, taskId)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *taskService) CreateTask(ctx context.Context, task *model.Task) (*model.Task, error) {
	//! Validate name
	//! Validate description

	task.Title = strings.TrimSpace(task.Title)
	task.Description = strings.TrimSpace(task.Description)

	if task.Status != "todo" &&
		task.Status != "in_progress" &&
		task.Status != "done" {
		return nil, fmt.Errorf("Error invalid task status")
	}

	err := s.taskRepo.Create(ctx, task)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *taskService) UpdateTask(ctx context.Context, taskId string, taskStatus string) error {
	err := s.taskRepo.UpdateStatus(ctx, taskId, taskStatus)

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
