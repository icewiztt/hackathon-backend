-- name: CreateSubmission :one
INSERT INTO submissions (
    username,
    fullname,
    taskid,
    taskname,
    task_subtasks, 
    submission_answers,
    submission_score,
    submission_results
) VALUES (
    $1 , $2 , $3 , $4 , $5 , $6 , $7 , $8
) RETURNING *;

-- name: ListSubmissions :many
SELECT * FROM submissions
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: ListAllSubmissions :many
SELECT * FROM submissions
ORDER BY id;

-- name: ListScores :many
SELECT MAX(submission_score) as score , fullname as user , taskname as task
FROM submissions
GROUP BY (fullname,taskname)
ORDER BY (fullname,taskname);