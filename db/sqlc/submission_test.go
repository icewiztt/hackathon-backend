package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thanhqt2002/hackathon/db/util"
)

func CreateRandomSubmission(t *testing.T, user User, task Task) Submission {
	arg := CreateSubmissionParams{
		Username:          user.Username,
		Fullname:          user.Fullname,
		Taskid:            task.ID,
		Taskname:          task.Shortname,
		TaskSubtasks:      task.Subtasks,
		SubmissionAnswers: util.RandomAnswers(int(task.Subtasks)),
		SubmissionResults: util.RandomSubmissionResult(int(task.Subtasks)),
		SubmissionScore:   float64(util.RandomInt(1, 100)),
	}

	submission, err := testQueries.CreateSubmission(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, submission)

	require.Equal(t, arg.Taskid, submission.Taskid)
	require.Equal(t, arg.TaskSubtasks, submission.TaskSubtasks)
	require.Equal(t, arg.SubmissionAnswers, submission.SubmissionAnswers)
	require.Equal(t, arg.SubmissionAnswers, submission.SubmissionAnswers)
	require.Len(t, submission.SubmissionAnswers, int(submission.TaskSubtasks))
	require.Len(t, submission.SubmissionResults, int(submission.TaskSubtasks))
	require.Equal(t, arg.SubmissionResults, submission.SubmissionResults)

	require.NotZero(t, submission.ID)
	require.NotZero(t, submission.CreatedAt)
	return submission
}

func TestCreateSubmission(t *testing.T) {
	user := CreateRandomUser(t)
	task := CreateRandomTask(t)
	CreateRandomSubmission(t, user, task)
}

func TestListSubmissions(t *testing.T) {
	for i := 0; i < 10; i++ {
		user := CreateRandomUser(t)
		task := CreateRandomTask(t)
		CreateRandomSubmission(t, user, task)
	}
	arg := ListSubmissionsParams{
		Limit:  5,
		Offset: 5,
	}
	submissions_list, err := testQueries.ListSubmissions(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, submissions_list, int(arg.Limit))

	for _, submission := range submissions_list {
		require.NotEmpty(t, submission)
	}
}

func TestListAllSubmissions(t *testing.T) {
	for i := 0; i < 10; i++ {
		user := CreateRandomUser(t)
		task := CreateRandomTask(t)
		CreateRandomSubmission(t, user, task)
	}
	submissions_list, err := testQueries.ListAllSubmissions(context.Background())
	require.NoError(t, err)

	for _, submission := range submissions_list {
		require.NotEmpty(t, submission)
	}
}

func TestListScores(t *testing.T) {
	for i := 0; i < 10; i++ {
		user := CreateRandomUser(t)
		task := CreateRandomTask(t)
		CreateRandomSubmission(t, user, task)
	}
	scores_list, err := testQueries.ListScores(context.Background())
	require.NoError(t, err)

	for _, score := range scores_list {
		require.NotEmpty(t, score)
	}
}
