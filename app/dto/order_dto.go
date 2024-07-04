package dto

import (
	"errors"

	vo "github.com/Zhima-Mochi/easy-task-api/domain/valueobject"
)

var (
	ErrInvalidID     = errors.New("invalid id")
	ErrInvalidName   = errors.New("invalid name")
	ErrInvalidStatus = errors.New("invalid status")
)

// TaskCreateRequest represents a request to create a task.
type TaskCreateRequest struct {
	Name string `json:"name" binding:"required" example:"task name"`
}

func (r *TaskCreateRequest) Validate() error {
	if r.Name == "" {
		return ErrInvalidName
	}
	return nil
}

// TaskCreateResponse represents a response to create a task.
type TaskCreateResponse struct {
	ID string `json:"id" example:"AEF4D3E1-4D3E-4D3E-4D3E-4D3E4D3E4D3E"`
}

// TaskResponse represents a response of a task.
type TaskResponse struct {
	ID        string    `json:"id" example:"AEF4D3E1-4D3E-4D3E-4D3E-4D3E4D3E4D3E"`
	Name      string    `json:"name" example:"task name"`
	Status    vo.Status `json:"status" example:"0"`
	CreatedAt string    `json:"created_at" example:"2024-07-04 12:00:00"`
	UpdateAt  string    `json:"update_at" example:"2024-07-04 15:00:00"`
}

// TaskUpdateRequest represents a request to update a task.
type TaskUpdateRequest struct {
	ID     string     `json:"id" example:"AEF4D3E1-4D3E-4D3E-4D3E-4D3E4D3E4D3E"`
	Name   *string    `json:"name" example:"task name"`
	Status *vo.Status `json:"status" example:"1"`
}

func (r *TaskUpdateRequest) Validate() error {
	if r.ID == "" {
		return ErrInvalidID
	}
	if r.Name != nil && *r.Name == "" {
		return ErrInvalidName
	}
	if r.Status != nil && *r.Status != vo.Completed {
		return ErrInvalidStatus
	}
	if r.Name == nil && r.Status == nil {
		return errors.New("no update field")
	}

	return nil
}
