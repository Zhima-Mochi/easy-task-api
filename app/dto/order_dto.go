package dto

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
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"update_at"`
}

// TaskUpdateRequest represents a request to update a task.
type TaskUpdateRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
