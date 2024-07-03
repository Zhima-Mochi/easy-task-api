package app

import (
	"context"

	"github.com/Zhima-Mochi/easy-task-api/app/dto"
	"github.com/Zhima-Mochi/easy-task-api/domain/service"
)

type TaskAppService interface {
	CreateTask(ctx context.Context, req *dto.TaskCreateRequest) (*dto.TaskCreateResponse, error)
}

type impl struct {
	taskService service.TaskService
}
