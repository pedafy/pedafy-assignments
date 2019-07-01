package database

import "time"

// Assignments refers to the assignments table in the database
type Assignments struct {
	ID             int        `db:"id" json:"id"`
	CreatorID      string     `db:"creator_id" json:"creator_id"`
	AssignedID     string     `db:"assigned_id" json:"assigned_id"`
	StatusID       int        `db:"status_id" json:"status_id"`
	TaskID         int        `db:"task_id" json:"task_id"`
	CreatedAt      *time.Time `db:"created_at" json:"created_at"`
	LastEdit       *time.Time `db:"last_edit" json:"last_edit"`
	DueDate        *time.Time `db:"due_date" json:"due_date"`
	CompletionDate *time.Time `db:"completion_date" json:"completion_date"`
	Title          *string    `db:"title" json:"title"`
	Description    *string    `db:"description" json:"description"`
}

// Status refers to the status table
type Status struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}
