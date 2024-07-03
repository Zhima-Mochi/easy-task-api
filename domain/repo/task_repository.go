package repo

import (
	"context"

	"github.com/Zhima-Mochi/easy-task-api/domain/entity"
)

type TaskRepository interface {
	FindAll(ctx context.Context) ([]*entity.Task, error)
	FindByID(ctx context.Context, id string) (*entity.Task, error)
	Create(ctx context.Context, task *entity.Task) error
	Update(ctx context.Context, task *entity.Task) error
	Delete(ctx context.Context, id string) error
}

type impl struct {
}

func NewTaskRepository() TaskRepository {
	return &impl{}
}

func (r *impl) FindAll(ctx context.Context) ([]*entity.Task, error) {
	return nil, nil
}

func (r *impl) FindByID(ctx context.Context, id string) (*entity.Task, error) {
	return nil, nil
}

func (r *impl) Create(ctx context.Context, task *entity.Task) error {
	return nil
}

func (r *impl) Update(ctx context.Context, task *entity.Task) error {
	return nil
}

func (r *impl) Delete(ctx context.Context, id string) error {
	return nil
}
