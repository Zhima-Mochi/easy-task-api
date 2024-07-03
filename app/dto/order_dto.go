package dto

// TaskCreateRequest represents a request to create a task.
type TaskCreateRequest struct {
	Name string `json:"name"`
}

// TaskCreateResponse represents a response to create a task.
type TaskCreateResponse struct {
	ID string `json:"id"`
}
