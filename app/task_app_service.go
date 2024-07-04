//go:generate mockgen -source=task_app_service.go -destination=mocks/mock_task_app_service.go -package=mocks
package app

import (
	"context"

	"github.com/Zhima-Mochi/easy-task-api/app/assembler"
	"github.com/Zhima-Mochi/easy-task-api/app/dto"
	"github.com/Zhima-Mochi/easy-task-api/domain/service"
	vo "github.com/Zhima-Mochi/easy-task-api/domain/valueobject"
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

func NewTaskAppService(taskService service.TaskService) TaskAppService {
	return &impl{taskService: taskService}
}

func (s *impl) CreateTask(ctx context.Context, req *dto.TaskCreateRequest) (*dto.TaskCreateResponse, error) {
	task := assembler.ToDomainTask(req)
	err := s.taskService.CreateTask(ctx, task)
	if err != nil {
		return nil, err
	}

	resp := assembler.ToCreateTaskResponse(task)
	return resp, nil
}

func (s *impl) GetAllTask(ctx context.Context) ([]*dto.TaskResponse, error) {
	tasks, err := s.taskService.GetAllTask(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]*dto.TaskResponse, 0, len(tasks))
	for _, task := range tasks {
		responses = append(responses, assembler.ToTaskResponse(task))
	}
	return responses, nil
}

func (s *impl) GetTaskByID(ctx context.Context, id string) (*dto.TaskResponse, error) {
	task, err := s.taskService.GetTaskByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return assembler.ToTaskResponse(task), nil
}

func (s *impl) UpdateTask(ctx context.Context, req *dto.TaskUpdateRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}
	task, err := s.taskService.GetTaskByID(ctx, req.ID)
	if err != nil {
		return err
	}

	if req.Name != nil {
		task.UpdateName(*req.Name)
	}

	if req.Status != nil && *req.Status == vo.Completed {
		task.Complete()
	}

	return s.taskService.UpdateTask(ctx, task)
}

func (s *impl) DeleteTask(ctx context.Context, id string) error {
	return s.taskService.DeleteTask(ctx, id)
}
