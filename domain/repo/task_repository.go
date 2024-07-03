//go:generate mockgen -source=task_repository.go -destination=mocks/mock_task_repository.go -package=mocks
package repo

import (
	"context"
	"errors"

	"github.com/Zhima-Mochi/easy-task-api/domain/entity"
	"github.com/Zhima-Mochi/easy-task-api/infra/persistence"
)

var (
	ErrorTaskNotFound      = errors.New("task not found")
	ErrorTaskAlreadyExists = errors.New("task already exists")
)

type TaskRepository interface {
	FindAll(ctx context.Context) ([]*entity.Task, error)
	Find(ctx context.Context, id string) (*entity.Task, error)
	Create(ctx context.Context, task *entity.Task) error
	Update(ctx context.Context, task *entity.Task) error
	Delete(ctx context.Context, id string) error
}

type impl struct {
	persistence *persistence.Persistence
}

func NewTaskRepository(persistence *persistence.Persistence) TaskRepository {
	return &impl{persistence: persistence}
}

func (r *impl) FindAll(ctx context.Context) ([]*entity.Task, error) {
	taskMap := r.persistence.GetAll()
	tasks := make([]*entity.Task, 0, len(taskMap))
	for _, val := range taskMap {
		task := val.(entity.Task)
		tasks = append(tasks, &task)
	}
	return tasks, nil
}

func (r *impl) Find(ctx context.Context, id string) (*entity.Task, error) {
	val, ok := r.persistence.Get(id)
	if !ok {
		return nil, ErrorTaskNotFound
	}
	task := val.(entity.Task)
	return &task, nil
}

func (r *impl) Create(ctx context.Context, task *entity.Task) error {
	if _, ok := r.persistence.Get(task.ID); ok {
		return ErrorTaskAlreadyExists
	}
	r.persistence.Set(task.ID, *task)
	return nil
}

func (r *impl) Update(ctx context.Context, task *entity.Task) error {
	if _, ok := r.persistence.Get(task.ID); !ok {
		return ErrorTaskNotFound
	}
	r.persistence.Set(task.ID, *task)
	return nil
}

func (r *impl) Delete(ctx context.Context, id string) error {
	r.persistence.Delete(id)
	return nil
}
