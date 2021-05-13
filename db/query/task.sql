-- name: CreateTask :one
INSERT INTO tasks (
    shortname,
    problemname,
    content,
    subtasks,
    answers
) VALUES (
    $1 , $2 , $3 , $4 , $5 
) RETURNING *;

-- name: GetTask :one
SELECT * FROM tasks 
WHERE id = $1;

-- name: ListTasks :many
SELECT * FROM tasks
ORDER BY shortname
LIMIT $1
OFFSET $2;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = $1;

-- name: UpdateAnswers :one
UPDATE tasks
SET answers = $2
WHERE id = $1
RETURNING *;

-- name: UpdateSubtasks :one
UPDATE tasks
SET subtasks = $2
WHERE id = $1
RETURNING *;

-- name: UpdateShortname :one
UPDATE tasks
SET shortname = $2
WHERE id = $1
RETURNING *;

-- name: UpdateProblemname :one
UPDATE tasks
SET problemname = $2
WHERE id = $1
RETURNING *;

-- name: UpdateContent :one
UPDATE tasks
SET content = $2
WHERE id = $1
RETURNING *;