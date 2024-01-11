package entity

import "time"

type TaskID int64
type TaskStatus string

const (
	TaskStatusTodo  TaskStatus = "todo"
	TaskStatusDoing TaskStatus = "doing"
	TaskStatusDone  TaskStatus = "done"
)

type Task struct {
	ID      TaskID     `db:"id"`
	Title   string     `db:"title"`
	Status  TaskStatus `db:"status"`
	Created time.Time  `db:"created_at"`
	Updated time.Time  `db:"updated_at"`
}

type Tasks []*Task
