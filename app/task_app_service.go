//go:generate mockgen -source=task_app_service.go -destination=mocks/mock_task_app_service.go -package=mocks
package app

import (
	"context"

	"github.com/Zhima-Mochi/easy-task-api/app/dto"
	"github.com/Zhima-Mochi/easy-task-api/domain/service"
)

type TaskAppService interface {
	CreateTask(ctx context.Context, req *dto.TaskCreateRequest) (*dto.TaskCreateResponse, error)
	GetAllTask(ctx context.Context) ([]*dto.TaskResponse, error)
	GetTaskByID(ctx context.Context, id string) (*dto.TaskResponse, error)
	UpdateTask(ctx context.Context, req *dto.TaskUpdateRequest) error
	DeleteTask(ctx context.Context, id string) error
}

type impl struct {
	taskService service.TaskService
}
