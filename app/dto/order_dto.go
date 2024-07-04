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
	Name string `json:"name"`
}

// TaskCreateResponse represents a response to create a task.
type TaskCreateResponse struct {
	ID string `json:"id"`
}

// TaskResponse represents a response of a task.
type TaskResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Status    vo.Status `json:"status"`
	CreatedAt string    `json:"created_at"`
	UpdateAt  string    `json:"update_at"`
}

// TaskUpdateRequest represents a request to update a task.
type TaskUpdateRequest struct {
	ID     string    `json:"id"`
	Name   string    `json:"name"`
	Status vo.Status `json:"status"`
}

func (r *TaskUpdateRequest) Validate() error {
	if r.ID == "" {
		return ErrInvalidID
	}
	if r.Name == "" {
		return ErrInvalidName
	}
	if r.Status != vo.Completed {
		return ErrInvalidStatus
	}
	return nil
}
