-- name: CreateTask :one
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
) RETURNING *;

-- name: GetTask :one
SELECT * FROM tasks 
WHERE id = $1;

-- name: ListTasks :many
SELECT * FROM tasks
WHERE official = true
ORDER BY shortname;

-- name: ListTasksAdmin :many
SELECT * FROM tasks
ORDER BY (official , shortname);

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = $1;

-- name: UpdateOfficial :one
UPDATE tasks
SET official = $2
WHERE id = $1
RETURNING *;


