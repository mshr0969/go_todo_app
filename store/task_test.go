package store

import (
	"context"
	"github/mshr0969/go_todo_app/clock"
	"github/mshr0969/go_todo_app/entity"
	"github/mshr0969/go_todo_app/testutil"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
	"github.com/jmoiron/sqlx"
)

func prepareTasks(ctx context.Context, t *testing.T, con Execer) entity.Tasks {
	t.Helper()
	// 一度きれいにしておく
	if _, err := con.ExecContext(ctx, "DELETE FROM task;"); err != nil {
		t.Logf("failed to initialize task: %v", err)
	}
	c := clock.FixedClocker{}
	wants := entity.Tasks{
		{
			Title: "want task 1", Status: "todo",
			Created: c.Now(), Updated: c.Now(),
		},
		{
			Title: "want task 2", Status: "todo",
			Created: c.Now(), Updated: c.Now(),
		},
		{
			Title: "want task 3", Status: "done",
			Created: c.Now(), Updated: c.Now(),
		},
	}
	result, err := con.ExecContext(ctx,
		`INSERT INTO task (title, status, created_at, updated_at)
			VALUES
				(?, ?, ?, ?),
				(?, ?, ?, ?),
				(?, ?, ?, ?);`,
		wants[0].Title, wants[0].Status, wants[0].Created, wants[0].Updated,
		wants[1].Title, wants[1].Status, wants[1].Created, wants[1].Updated,
		wants[2].Title, wants[2].Status, wants[2].Created, wants[2].Updated,
	)
	if err != nil {
		t.Fatal(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		t.Fatal(err)
	}
	wants[0].ID = entity.TaskID(id)
	wants[1].ID = entity.TaskID(id + 1)
	wants[2].ID = entity.TaskID(id + 2)
	return wants
}

func TestRepository_AddTask(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	c := clock.FixedClocker{}
	var wantID int64 = 2
	task := &entity.Task{
		Title:  "ok task",
		Status: "todo",
	}

	// 空のモックを作成
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { db.Close() })

	// go-sqlmockを使うことで、実際のデータベース接続を必要とせずに、テスト内でSQLドライバの動作をシミュレート
	mock.ExpectExec(
		`INSERT INTO tasks \(title, status\) VALUES \(\?, \?\)`,
	).WithArgs(task.Title, task.Status).
		// WillReturnRsultで、SQLの実行後に返すべき結果を定義
		// wantID: 適当なレコードID, 1: 1行が挿入される
		WillReturnResult(sqlmock.NewResult(wantID, 1))

	xdb := sqlx.NewDb(db, "mysql")
	r := &Repository{Clocker: c}
	if err := r.AddTask(ctx, xdb, task); err != nil {
		t.Errorf("want no error, but got %v", err)
	}
}

func TestRepository_ListTasks(t *testing.T) {
	ctx := context.Background()
	// トランザクションをはることで、テストケースだけのテーブルにする
	tx, err := testutil.OpenDBForTest(t).BeginTxx(ctx, nil)
	// 終わったらもどす
	t.Cleanup(func() { _ = tx.Rollback() })
	if err != nil {
		t.Fatal(err)
	}
	wants := prepareTasks(ctx, t, tx)

	sut := &Repository{}
	gots, err := sut.ListTasks(ctx, tx)
	if err != nil {
		t.Fatalf("unexected error: %v", err)
	}
	if d := cmp.Diff(gots, wants); len(d) != 0 {
		t.Errorf("differs: (-got +want)\n%s", d)
	}
}
