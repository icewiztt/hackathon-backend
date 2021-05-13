-- name: CreateSubmission :one
INSERT INTO submissions (
    from_user_id,
    to_task_id,
    task_subtasks, 
    submission_answers
) VALUES (
    $1 , $2 , $3 , $4
) RETURNING *;

-- name: ListSubmissions :many
SELECT * FROM submissions
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: ListAllSubmissions :many
SELECT * FROM submissions
ORDER BY id;