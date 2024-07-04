package assembler

import (
	"github.com/Zhima-Mochi/easy-task-api/app/dto"
	"github.com/Zhima-Mochi/easy-task-api/domain/entity"
)

func ToDomainTask(req *dto.TaskCreateRequest) *entity.Task {
	return entity.NewTask(req.Name)
}

func ToCreateTaskResponse(task *entity.Task) *dto.TaskCreateResponse {
	return &dto.TaskCreateResponse{
		ID: task.ID,
	}
}

func ToTaskResponse(task *entity.Task) *dto.TaskResponse {
	return &dto.TaskResponse{
		ID:        task.ID,
		Name:      task.Name,
		Status:    task.Status,
		CreatedAt: task.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdateAt:  task.UpdateAt.Format("2006-01-02 15:04:05"),
	}
}
