package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thanhqt2002/hackathon/db/util"
)

func CreateRandomSubmission(t *testing.T, user User, task Task) Submission {
	arg := CreateSubmissionParams{
		FromUserID:        user.ID,
		ToTaskID:          task.ID,
		TaskSubtasks:      task.Subtasks,
		SubmissionAnswers: util.RandomAnswers(int(task.Subtasks)),
	}

	submission, err := testQueries.CreateSubmission(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, submission)

	require.Equal(t, arg.FromUserID, submission.FromUserID)
	require.Equal(t, arg.ToTaskID, submission.ToTaskID)
	require.Equal(t, arg.TaskSubtasks, submission.TaskSubtasks)
	require.Equal(t, arg.SubmissionAnswers, submission.SubmissionAnswers)
	require.Len(t, submission.SubmissionAnswers, int(submission.TaskSubtasks))

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
