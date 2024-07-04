package app

import (
	"context"
	"errors"
	"testing"

	"github.com/Zhima-Mochi/easy-task-api/app/dto"
	"github.com/Zhima-Mochi/easy-task-api/domain/entity"
	"github.com/Zhima-Mochi/easy-task-api/domain/service/mocks"
	"github.com/golang/mock/gomock"
)

var (
	mockCtx = context.Background()

	mockTask = &entity.Task{
		ID:     "1",
		Name:   "task1",
		Status: 0,
	}
)

type Mocks struct {
	mockTaskService *mocks.MockTaskService
}

func NewMocks(ctrl *gomock.Controller) *Mocks {
	return &Mocks{
		mockTaskService: mocks.NewMockTaskService(ctrl),
	}
}

func TestCreateTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks := NewMocks(ctrl)

	taskAppService := NewTaskAppService(mocks.mockTaskService)

	t.Run("CreateTask", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			mocks.mockTaskService.EXPECT().CreateTask(mockCtx, gomock.Any()).Return(nil)

			req := &dto.TaskCreateRequest{
				Name: "task1",
			}

			_, err := taskAppService.CreateTask(mockCtx, req)
			if err != nil {
				t.Errorf("error should be nil, but got: %v", err)
			}
		})

		t.Run("Error", func(t *testing.T) {
			mocks.mockTaskService.EXPECT().CreateTask(mockCtx, gomock.Any()).Return(errors.New("error"))

			req := &dto.TaskCreateRequest{
				Name: "task1",
			}

			_, err := taskAppService.CreateTask(mockCtx, req)
			if err == nil {
				t.Error("error should not be nil")
			}
		})
	})
}

func TestGetAllTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks := NewMocks(ctrl)

	taskAppService := NewTaskAppService(mocks.mockTaskService)

	t.Run("GetAllTask", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			mocks.mockTaskService.EXPECT().GetAllTask(mockCtx).Return([]*entity.Task{mockTask}, nil)

			res, err := taskAppService.GetAllTask(mockCtx)
			if err != nil {
				t.Errorf("error should be nil, but got: %v", err)
			}

			if len(res) != 1 {
				t.Errorf("response length should be 1, but got: %v", len(res))
			}

			if res[0].ID != mockTask.ID {
				t.Errorf("response ID should be %v, but got: %v", mockTask.ID, res[0].ID)
			}
		})

		t.Run("Error", func(t *testing.T) {
			mocks.mockTaskService.EXPECT().GetAllTask(mockCtx).Return(nil, errors.New("error"))

			_, err := taskAppService.GetAllTask(mockCtx)
			if err == nil {
				t.Error("error should not be nil")
			}
		})
	})
}

func TestGetTaskByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks := NewMocks(ctrl)

	taskAppService := NewTaskAppService(mocks.mockTaskService)

	t.Run("GetTaskByID", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			mocks.mockTaskService.EXPECT().GetTaskByID(mockCtx, mockTask.ID).Return(mockTask, nil)

			res, err := taskAppService.GetTaskByID(mockCtx, mockTask.ID)
			if err != nil {
				t.Errorf("error should be nil, but got: %v", err)
			}

			if res.ID != mockTask.ID {
				t.Errorf("response ID should be %v, but got: %v", mockTask.ID, res.ID)
			}
		})

		t.Run("Error", func(t *testing.T) {
			mocks.mockTaskService.EXPECT().GetTaskByID(mockCtx, mockTask.ID).Return(nil, errors.New("error"))

			_, err := taskAppService.GetTaskByID(mockCtx, mockTask.ID)
			if err == nil {
				t.Error("error should not be nil")
			}
		})
	})
}

func TestUpdateTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks := NewMocks(ctrl)

	taskAppService := NewTaskAppService(mocks.mockTaskService)

	t.Run("UpdateTask", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			mocks.mockTaskService.EXPECT().GetTaskByID(mockCtx, mockTask.ID).Return(mockTask, nil)
			mocks.mockTaskService.EXPECT().UpdateTask(mockCtx, gomock.Any()).Return(nil)

			req := &dto.TaskUpdateRequest{
				ID: mockTask.ID,
			}
			targetName := "task2"
			req.Name = &targetName

			err := taskAppService.UpdateTask(mockCtx, req)
			if err != nil {
				t.Errorf("error should be nil, but got: %v", err)
			}
		})

		t.Run("Error", func(t *testing.T) {
			mocks.mockTaskService.EXPECT().GetTaskByID(mockCtx, mockTask.ID).Return(nil, errors.New("error"))

			req := &dto.TaskUpdateRequest{
				ID: mockTask.ID,
			}
			targetName := "task2"
			req.Name = &targetName

			err := taskAppService.UpdateTask(mockCtx, req)
			if err == nil {
				t.Error("error should not be nil")
			}
		})
	})
}

func TestDeleteTask(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mocks := NewMocks(ctrl)

	taskAppService := NewTaskAppService(mocks.mockTaskService)

	t.Run("DeleteTask", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			mocks.mockTaskService.EXPECT().DeleteTask(mockCtx, mockTask.ID).Return(nil)

			err := taskAppService.DeleteTask(mockCtx, mockTask.ID)
			if err != nil {
				t.Errorf("error should be nil, but got: %v", err)
			}
		})

		t.Run("Error", func(t *testing.T) {
			mocks.mockTaskService.EXPECT().DeleteTask(mockCtx, mockTask.ID).Return(errors.New("error"))

			err := taskAppService.DeleteTask(mockCtx, mockTask.ID)
			if err == nil {
				t.Error("error should not be nil")
			}
		})
	})
}
