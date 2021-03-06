// Code generated by sqlc. DO NOT EDIT.
// source: task.sql

package db

import (
	"context"

	"github.com/lib/pq"
)

const createTask = `-- name: CreateTask :one
INSERT INTO tasks (
    shortname,
    problemname,
    content,
    subtasks,
    answers,
    subtasks_score,
    official
) VALUES (
    $1 , $2 , $3 , $4 , $5 , $6 , $7
) RETURNING id, shortname, problemname, content, subtasks, answers, subtasks_score, official, created_at
`

type CreateTaskParams struct {
	Shortname     string    `json:"shortname"`
	Problemname   string    `json:"problemname"`
	Content       string    `json:"content"`
	Subtasks      int32     `json:"subtasks"`
	Answers       []string  `json:"answers"`
	SubtasksScore []float64 `json:"subtasks_score"`
	Official      bool      `json:"official"`
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, createTask,
		arg.Shortname,
		arg.Problemname,
		arg.Content,
		arg.Subtasks,
		pq.Array(arg.Answers),
		pq.Array(arg.SubtasksScore),
		arg.Official,
	)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Shortname,
		&i.Problemname,
		&i.Content,
		&i.Subtasks,
		pq.Array(&i.Answers),
		pq.Array(&i.SubtasksScore),
		&i.Official,
		&i.CreatedAt,
	)
	return i, err
}

const deleteTask = `-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = $1
`

func (q *Queries) DeleteTask(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteTask, id)
	return err
}

const getTask = `-- name: GetTask :one
SELECT id, shortname, problemname, content, subtasks, answers, subtasks_score, official, created_at FROM tasks 
WHERE id = $1
`

func (q *Queries) GetTask(ctx context.Context, id int32) (Task, error) {
	row := q.db.QueryRowContext(ctx, getTask, id)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Shortname,
		&i.Problemname,
		&i.Content,
		&i.Subtasks,
		pq.Array(&i.Answers),
		pq.Array(&i.SubtasksScore),
		&i.Official,
		&i.CreatedAt,
	)
	return i, err
}

const listTasks = `-- name: ListTasks :many
SELECT id, shortname, problemname, content, subtasks, answers, subtasks_score, official, created_at FROM tasks
WHERE official = true
ORDER BY shortname
`

func (q *Queries) ListTasks(ctx context.Context) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, listTasks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Task{}
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.Shortname,
			&i.Problemname,
			&i.Content,
			&i.Subtasks,
			pq.Array(&i.Answers),
			pq.Array(&i.SubtasksScore),
			&i.Official,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTasksAdmin = `-- name: ListTasksAdmin :many
SELECT id, shortname, problemname, content, subtasks, answers, subtasks_score, official, created_at FROM tasks
ORDER BY (official , shortname)
`

func (q *Queries) ListTasksAdmin(ctx context.Context) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, listTasksAdmin)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Task{}
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.Shortname,
			&i.Problemname,
			&i.Content,
			&i.Subtasks,
			pq.Array(&i.Answers),
			pq.Array(&i.SubtasksScore),
			&i.Official,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateOfficial = `-- name: UpdateOfficial :one
UPDATE tasks
SET official = $2
WHERE id = $1
RETURNING id, shortname, problemname, content, subtasks, answers, subtasks_score, official, created_at
`

type UpdateOfficialParams struct {
	ID       int32 `json:"id"`
	Official bool  `json:"official"`
}

func (q *Queries) UpdateOfficial(ctx context.Context, arg UpdateOfficialParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, updateOfficial, arg.ID, arg.Official)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Shortname,
		&i.Problemname,
		&i.Content,
		&i.Subtasks,
		pq.Array(&i.Answers),
		pq.Array(&i.SubtasksScore),
		&i.Official,
		&i.CreatedAt,
	)
	return i, err
}
