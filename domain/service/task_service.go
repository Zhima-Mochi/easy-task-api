//go:generate mockgen -source=task_service.go -destination=mocks/mock_task_service.go -package=mocks
package service

import (
	"context"
	"errors"

	"github.com/Zhima-Mochi/easy-task-api/domain/entity"
	"github.com/Zhima-Mochi/easy-task-api/domain/repo"
)

var (
	ErrorTaskNotFound = errors.New("task not found")

	ErrorTaskAlreadyExists = errors.New("task already exists")
)

type TaskService interface {
	GetAllTask(ctx context.Context) ([]*entity.Task, error)
	GetTaskByID(ctx context.Context, id string) (*entity.Task, error)
	CreateTask(ctx context.Context, task *entity.Task) error
	UpdateTask(ctx context.Context, task *entity.Task) error
	DeleteTask(ctx context.Context, id string) error
}

type impl struct {
	taskRepo repo.TaskRepository
}

func NewTaskService(taskRepo repo.TaskRepository) TaskService {
	return &impl{taskRepo: taskRepo}
}

func (s *impl) GetAllTask(ctx context.Context) ([]*entity.Task, error) {
	return s.taskRepo.FindAll(ctx)
}

func (s *impl) GetTaskByID(ctx context.Context, id string) (*entity.Task, error) {
	task, err := s.taskRepo.Find(ctx, id)
	if err != nil {
		if errors.Is(err, repo.ErrorTaskNotFound) {
			return nil, ErrorTaskNotFound
		}
		return nil, err
	}
	return task, nil
}

func (s *impl) CreateTask(ctx context.Context, task *entity.Task) error {
	if err := s.taskRepo.Create(ctx, task); err != nil {
		if errors.Is(err, repo.ErrorTaskAlreadyExists) {
			return ErrorTaskAlreadyExists
		}
		return err
	}
	return nil
}

func (s *impl) UpdateTask(ctx context.Context, task *entity.Task) error {
	if err := s.taskRepo.Update(ctx, task); err != nil {
		if errors.Is(err, repo.ErrorTaskNotFound) {
			return ErrorTaskNotFound
		}
		return err
	}
	return nil
}

func (s *impl) DeleteTask(ctx context.Context, id string) error {
	return s.taskRepo.Delete(ctx, id)
}
