package main

const (
	StatusTodo       = 0
	StatusInProgress = 1
	StatusDone       = 2
)

type TaskTracker struct {
	ID          int  `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	Status      int    `json:"status,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}


// var	demoTask = TaskTracker{
	// 	ID:          1,
	// 	Description: "Complete the Go project",
	// 	Status:      StatusTodo,
	// 	CreatedAt:   "2023-10-01T10:00:00Z",
	// 	UpdatedAt:   "2023-10-01T10:00:00Z",
	// }
	// var	demoTaskUpdated = TaskTracker{
	// 	ID:          1,
	// 	Description: "Complete the Go project New Task Updating",
	// 	Status:      StatusTodo,
	// 	CreatedAt:   "2023-10-01T10:00:00Z",
	// 	UpdatedAt:   "2023-10-01T10:00:00Z",
	// }
