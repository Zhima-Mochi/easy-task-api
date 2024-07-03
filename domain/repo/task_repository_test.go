package repo

import (
	"context"
	"sort"
	"testing"
	"time"

	"github.com/Zhima-Mochi/easy-task-api/domain/entity"
	"github.com/Zhima-Mochi/easy-task-api/infra/persistence"
)

var (
	mockCtx = context.Background()

	mockTask1 = entity.Task{
		ID:        "1",
		Name:      "task1",
		CreatedAt: time.Unix(0, 0),
	}

	mockTask2 = entity.Task{
		ID:        "2",
		Name:      "task2",
		CreatedAt: time.Unix(0, 1),
	}
)

type Mocks struct {
}

func TestFindAll(t *testing.T) {
	tests := []struct {
		name  string
		setUp func(persistence *persistence.Persistence)
		check func(t *testing.T, res []*entity.Task, err error)
	}{
		{
			name: "success",
			setUp: func(persistence *persistence.Persistence) {
				persistence.Set(mockTask1.ID, mockTask1)
				persistence.Set(mockTask2.ID, mockTask2)
			},
			check: func(t *testing.T, res []*entity.Task, err error) {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}

				if len(res) != 2 {
					t.Errorf("unexpected result: %v", res)
				}

				sort.Slice(res, func(i, j int) bool {
					return res[i].CreatedAt.Before(res[j].CreatedAt)
				})

				if *res[0] != mockTask1 {
					t.Errorf("unexpected result: %v", res)
				}

				if *res[1] != mockTask2 {
					t.Errorf("unexpected result: %v", res)
				}
			},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			persistence := persistence.NewPersistence()
			taskRepo := NewTaskRepository(persistence)

			// Set up
			if tt.setUp != nil {
				tt.setUp(persistence)
			}

			// Call the code we are testing.
			res, err := taskRepo.FindAll(mockCtx)

			// Check the results.
			if tt.check != nil {
				tt.check(t, res, err)
			}
		})
	}
}

func TestFind(t *testing.T) {
	tests := []struct {
		name  string
		setUp func(persistence *persistence.Persistence)
		check func(t *testing.T, res *entity.Task, err error)
	}{
		{
			name: "success",
			setUp: func(persistence *persistence.Persistence) {
				persistence.Set(mockTask1.ID, mockTask1)
			},
			check: func(t *testing.T, res *entity.Task, err error) {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}

				if *res != mockTask1 {
					t.Errorf("unexpected result: %v", res)
				}
			},
		},
		{
			name: "not found",
			setUp: func(persistence *persistence.Persistence) {
			},
			check: func(t *testing.T, res *entity.Task, err error) {
				if err != ErrorTaskNotFound {
					t.Errorf("unexpected error: %v", err)
				}

				if res != nil {
					t.Errorf("unexpected result: %v", res)
				}
			},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			persistence := persistence.NewPersistence()
			taskRepo := NewTaskRepository(persistence)

			// Set up
			if tt.setUp != nil {
				tt.setUp(persistence)
			}

			// Call the code we are testing.
			res, err := taskRepo.Find(mockCtx, mockTask1.ID)

			// Check the results.
			if tt.check != nil {
				tt.check(t, res, err)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	mockTask := mockTask1
	tests := []struct {
		name  string
		setUp func(persistence *persistence.Persistence)
		check func(t *testing.T, err error)
	}{
		{
			name: "success",
			setUp: func(persistence *persistence.Persistence) {
			},
			check: func(t *testing.T, err error) {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			},
		},
		{
			name: "already exists",
			setUp: func(persistence *persistence.Persistence) {
				persistence.Set(mockTask1.ID, mockTask1)
			},
			check: func(t *testing.T, err error) {
				if err != ErrorTaskAlreadyExists {
					t.Errorf("unexpected error: %v", err)
				}
			},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			persistence := persistence.NewPersistence()
			taskRepo := NewTaskRepository(persistence)

			// Set up
			if tt.setUp != nil {
				tt.setUp(persistence)
			}

			// Call the code we are testing.
			err := taskRepo.Create(mockCtx, &mockTask)

			// Check the results.
			if tt.check != nil {
				tt.check(t, err)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	mockUpdatedTask := mockTask1
	mockUpdatedTask.Name = "updated"

	tests := []struct {
		name  string
		setUp func(persistence *persistence.Persistence)
		check func(t *testing.T, task *entity.Task, err error)
	}{
		{
			name: "success",
			setUp: func(persistence *persistence.Persistence) {
				persistence.Set(mockTask1.ID, mockTask1)
			},
			check: func(t *testing.T, task *entity.Task, err error) {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}

				if *task != mockUpdatedTask {
					t.Errorf("unexpected result: %v", task)
				}
			},
		},
		{
			name: "not found",
			setUp: func(persistence *persistence.Persistence) {
			},
			check: func(t *testing.T, task *entity.Task, err error) {
				if err != ErrorTaskNotFound {
					t.Errorf("unexpected error: %v", err)
				}

				if task != nil {
					t.Errorf("unexpected result: %v", task)
				}
			},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			persistence := persistence.NewPersistence()
			taskRepo := NewTaskRepository(persistence)

			// Set up
			if tt.setUp != nil {
				tt.setUp(persistence)
			}

			// Call the code we are testing.
			err := taskRepo.Update(mockCtx, &mockUpdatedTask)
			var task *entity.Task
			if val, ok := persistence.Get(mockUpdatedTask.ID); ok {
				t := val.(entity.Task)
				task = &t
			}

			// Check the results.
			if tt.check != nil {
				tt.check(t, task, err)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	mockDeletedID := mockTask1.ID

	tests := []struct {
		name  string
		setUp func(persistence *persistence.Persistence)
		check func(t *testing.T, err error)
	}{
		{
			name: "success",
			setUp: func(persistence *persistence.Persistence) {
				persistence.Set(mockTask1.ID, mockTask1)
			},
			check: func(t *testing.T, err error) {
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
			},
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			persistence := persistence.NewPersistence()
			taskRepo := NewTaskRepository(persistence)

			// Set up
			if tt.setUp != nil {
				tt.setUp(persistence)
			}

			// Call the code we are testing.
			err := taskRepo.Delete(mockCtx, mockDeletedID)

			// Check the results.
			if tt.check != nil {
				tt.check(t, err)
			}

			if _, ok := persistence.Get(mockDeletedID); ok {
				t.Errorf("task not deleted")
			}
		})
	}
}
