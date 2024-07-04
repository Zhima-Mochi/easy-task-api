package service

import (
	"context"
	"errors"
	"testing"

	"github.com/Zhima-Mochi/easy-task-api/domain/entity"
	"github.com/Zhima-Mochi/easy-task-api/domain/repo/mocks"
	"github.com/golang/mock/gomock"
)

var (
	mockCtx = context.Background()
)

func TestGetAllTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTaskRepo := mocks.NewMockTaskRepository(ctrl)
	taskService := NewTaskService(mockTaskRepo)

	t.Run("Success", func(t *testing.T) {
		expectedTasks := []*entity.Task{
			{ID: "1", Name: "task1", Status: 0},
			{ID: "2", Name: "task2", Status: 1},
		}

		mockTaskRepo.EXPECT().FindAll(mockCtx).Return(expectedTasks, nil)

		tasks, err := taskService.GetAllTask(mockCtx)
		if err != nil {
			t.Errorf("error should be nil, but got: %v", err)
		}

		if len(tasks) != len(expectedTasks) {
			t.Errorf("expected %d tasks, but got %d", len(expectedTasks), len(tasks))
		}
	})

	t.Run("Error", func(t *testing.T) {
		mockTaskRepo.EXPECT().FindAll(mockCtx).Return(nil, errors.New("error"))

		_, err := taskService.GetAllTask(mockCtx)
		if err == nil {
			t.Error("error should not be nil")
		}
	})
}

func TestGetTaskByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTaskRepo := mocks.NewMockTaskRepository(ctrl)
	taskService := NewTaskService(mockTaskRepo)

	t.Run("Success", func(t *testing.T) {
		expectedTask := &entity.Task{ID: "1", Name: "task1", Status: 0}

		mockTaskRepo.EXPECT().Find(mockCtx, "1").Return(expectedTask, nil)

		task, err := taskService.GetTaskByID(mockCtx, "1")
		if err != nil {
			t.Errorf("error should be nil, but got: %v", err)
		}

		if task.ID != expectedTask.ID {
			t.Errorf("expected task id %s, but got %s", expectedTask.ID, task.ID)
		}
	})

	t.Run("Error", func(t *testing.T) {
		mockTaskRepo.EXPECT().Find(mockCtx, "1").Return(nil, errors.New("error"))

		_, err := taskService.GetTaskByID(mockCtx, "1")
		if err == nil {
			t.Error("error should not be nil")
		}
	})
}

func TestCreateTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTaskRepo := mocks.NewMockTaskRepository(ctrl)
	taskService := NewTaskService(mockTaskRepo)

	t.Run("Success", func(t *testing.T) {
		task := &entity.Task{Name: "task1", Status: 0}

		mockTaskRepo.EXPECT().Create(mockCtx, task).Return(nil)

		err := taskService.CreateTask(mockCtx, task)
		if err != nil {
			t.Errorf("error should be nil, but got: %v", err)
		}
	})

	t.Run("Error", func(t *testing.T) {
		task := &entity.Task{Name: "task1", Status: 0}

		mockTaskRepo.EXPECT().Create(mockCtx, task).Return(errors.New("error"))

		err := taskService.CreateTask(mockCtx, task)
		if err == nil {
			t.Error("error should not be nil")
		}
	})
}

func TestUpdateTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTaskRepo := mocks.NewMockTaskRepository(ctrl)
	taskService := NewTaskService(mockTaskRepo)

	t.Run("Success", func(t *testing.T) {
		task := &entity.Task{ID: "1", Name: "task1", Status: 0}

		mockTaskRepo.EXPECT().Update(mockCtx, task).Return(nil)

		err := taskService.UpdateTask(mockCtx, task)
		if err != nil {
			t.Errorf("error should be nil, but got: %v", err)
		}
	})

	t.Run("Error", func(t *testing.T) {
		task := &entity.Task{ID: "1", Name: "task1", Status: 0}

		mockTaskRepo.EXPECT().Update(mockCtx, task).Return(errors.New("error"))

		err := taskService.UpdateTask(mockCtx, task)
		if err == nil {
			t.Error("error should not be nil")
		}
	})
}

func TestDeleteTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTaskRepo := mocks.NewMockTaskRepository(ctrl)
	taskService := NewTaskService(mockTaskRepo)

	t.Run("Success", func(t *testing.T) {
		mockTaskRepo.EXPECT().Delete(mockCtx, "1").Return(nil)

		err := taskService.DeleteTask(mockCtx, "1")
		if err != nil {
			t.Errorf("error should be nil, but got: %v", err)
		}
	})

	t.Run("Error", func(t *testing.T) {
		mockTaskRepo.EXPECT().Delete(mockCtx, "1").Return(errors.New("error"))

		err := taskService.DeleteTask(mockCtx, "1")
		if err == nil {
			t.Error("error should not be nil")
		}
	})
}
