package store

import (
	"context"
	"github/mshr0969/go_todo_app/entity"
)

func (r *Repository) AddTask(
	ctx context.Context, db Execer, t *entity.Task,
) error {
	sql := `INSERT INTO tasks (title, status)
	VALUES (?, ?)`
	result, err := db.ExecContext(ctx, sql, t.Title, t.Status)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	t.ID = entity.TaskID(id)
	return nil
}

func (r *Repository) ListTasks(
	ctx context.Context, db Queryer,
) (entity.Tasks, error) {
	tasks := entity.Tasks{}
	sql := `SELECT id, title, status, created_at, updated_at FROM task;`
	if err := db.SelectContext(ctx, &tasks, sql); err != nil {
		return nil, err
	}
	return tasks, nil
}
