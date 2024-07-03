package entity

import (
	"time"

	vo "github.com/Zhima-Mochi/easy-task-api/domain/valueobject"
	"github.com/google/uuid"
)

type Task struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Status    vo.Status `json:"status"`
	CreatedAt time.Time `json:"-"`
	UpdateAt  time.Time `json:"-"`
}

func NewTask(name string) *Task {
	return &Task{
		ID:        uuid.New().String(),
		Name:      name,
		Status:    vo.Incomplete,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}
}

func (t *Task) Complete() {
	t.Status = vo.Completed
	t.UpdateAt = time.Now()
}

func (t *Task) UpdateName(name string) {
	t.Name = name
	t.UpdateAt = time.Now()
}
