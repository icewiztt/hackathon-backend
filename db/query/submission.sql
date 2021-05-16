-- name: CreateSubmission :one
INSERT INTO submissions (
    from_user_id,
    to_task_id,
    task_subtasks, 
    submission_answers,
    submission_score,
    submission_results
) VALUES (
    $1 , $2 , $3 , $4 , $5 , $6
) RETURNING *;

-- name: ListSubmissions :many
SELECT * FROM submissions
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: ListAllSubmissions :many
SELECT * FROM submissions
ORDER BY id;
