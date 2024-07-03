// Code generated by MockGen. DO NOT EDIT.
// Source: task_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	entity "github.com/Zhima-Mochi/easy-task-api/domain/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockTaskService is a mock of TaskService interface.
type MockTaskService struct {
	ctrl     *gomock.Controller
	recorder *MockTaskServiceMockRecorder
}

// MockTaskServiceMockRecorder is the mock recorder for MockTaskService.
type MockTaskServiceMockRecorder struct {
	mock *MockTaskService
}

// NewMockTaskService creates a new mock instance.
func NewMockTaskService(ctrl *gomock.Controller) *MockTaskService {
	mock := &MockTaskService{ctrl: ctrl}
	mock.recorder = &MockTaskServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskService) EXPECT() *MockTaskServiceMockRecorder {
	return m.recorder
}

// CreateTask mocks base method.
func (m *MockTaskService) CreateTask(ctx context.Context, task *entity.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTask", ctx, task)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTask indicates an expected call of CreateTask.
func (mr *MockTaskServiceMockRecorder) CreateTask(ctx, task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTask", reflect.TypeOf((*MockTaskService)(nil).CreateTask), ctx, task)
}

// DeleteTask mocks base method.
func (m *MockTaskService) DeleteTask(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTask", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTask indicates an expected call of DeleteTask.
func (mr *MockTaskServiceMockRecorder) DeleteTask(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTask", reflect.TypeOf((*MockTaskService)(nil).DeleteTask), ctx, id)
}

// GetAllTask mocks base method.
func (m *MockTaskService) GetAllTask(ctx context.Context) ([]*entity.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTask", ctx)
	ret0, _ := ret[0].([]*entity.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTask indicates an expected call of GetAllTask.
func (mr *MockTaskServiceMockRecorder) GetAllTask(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTask", reflect.TypeOf((*MockTaskService)(nil).GetAllTask), ctx)
}

// GetTaskByID mocks base method.
func (m *MockTaskService) GetTaskByID(ctx context.Context, id string) (*entity.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskByID", ctx, id)
	ret0, _ := ret[0].(*entity.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTaskByID indicates an expected call of GetTaskByID.
func (mr *MockTaskServiceMockRecorder) GetTaskByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskByID", reflect.TypeOf((*MockTaskService)(nil).GetTaskByID), ctx, id)
}

// UpdateTask mocks base method.
func (m *MockTaskService) UpdateTask(ctx context.Context, task *entity.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTask", ctx, task)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTask indicates an expected call of UpdateTask.
func (mr *MockTaskServiceMockRecorder) UpdateTask(ctx, task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTask", reflect.TypeOf((*MockTaskService)(nil).UpdateTask), ctx, task)
}