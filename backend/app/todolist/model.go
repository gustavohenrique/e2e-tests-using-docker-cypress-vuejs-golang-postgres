package todolist

import "time"

type Task struct {
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
}
