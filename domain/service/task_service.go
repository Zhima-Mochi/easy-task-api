package service

import (
	"context"

	"github.com/Zhima-Mochi/easy-task-api/domain/entity"
	"github.com/Zhima-Mochi/easy-task-api/domain/repo"
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
	return s.taskRepo.FindByID(ctx, id)
}

func (s *impl) CreateTask(ctx context.Context, task *entity.Task) error {
	return s.taskRepo.Create(ctx, task)
}

func (s *impl) UpdateTask(ctx context.Context, task *entity.Task) error {
	return s.taskRepo.Update(ctx, task)
}

func (s *impl) DeleteTask(ctx context.Context, id string) error {
	return s.taskRepo.Delete(ctx, id)
}
