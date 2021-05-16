package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thanhqt2002/hackathon/db/util"
)

func CreateRandomTask(t *testing.T) Task {
	Subtasks_tmp := util.RandomSubtasks()
	arg := CreateTaskParams{
		Shortname:     util.RandomShortname(),
		Problemname:   util.RandomProblemname(),
		Content:       util.RandomContent(),
		Subtasks:      Subtasks_tmp,
		SubtasksScore: util.RandomSubScore(int(Subtasks_tmp)),
		Answers:       util.RandomAnswers(int(Subtasks_tmp)),
		Official:      true,
	}

	task, err := testQueries.CreateTask(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, task)

	require.Equal(t, arg.Shortname, task.Shortname)
	require.Equal(t, arg.Problemname, task.Problemname)
	require.Equal(t, arg.Content, task.Content)
	require.Equal(t, arg.Subtasks, task.Subtasks)
	require.Equal(t, arg.Answers, task.Answers)
	require.Len(t, task.Answers, int(task.Subtasks))
	require.NotZero(t, task.ID)
	require.NotZero(t, task.CreatedAt)
	return task
}

func TestCreateRandomTask(t *testing.T) {
	CreateRandomTask(t)
}

func TestGetTask(t *testing.T) {
	Task := CreateRandomTask(t)
	task_get, err := testQueries.GetTask(context.Background(), Task.ID)
	require.NoError(t, err)
	require.NotEmpty(t, task_get)
	require.Equal(t, Task, task_get)
}

func TestListTasks(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomTask(t)
	}
	tasks_list, err := testQueries.ListTasks(context.Background())
	require.NoError(t, err)
	for _, task := range tasks_list {
		require.NotEmpty(t, task)
		require.True(t, task.Official)
	}
}

func TestListTasksAdmin(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomTask(t)
	}
	tasks_list, err := testQueries.ListTasksAdmin(context.Background())
	require.NoError(t, err)
	for _, task := range tasks_list {
		require.NotEmpty(t, task)
	}
}

func TestDeleteTask(t *testing.T) {
	task := CreateRandomTask(t)
	err := testQueries.DeleteTask(context.Background(), task.ID)
	require.NoError(t, err)
	task_get, err := testQueries.GetTask(context.Background(), task.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, task_get)
}

func TestUpdateOfficial(t *testing.T) {
	task := CreateRandomTask(t)
	arg := UpdateOfficialParams{
		ID:       task.ID,
		Official: task.Official,
	}
	task_update, err := testQueries.UpdateOfficial(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, task_update.Official, arg.Official)
	task.Official = task_update.Official
	require.Equal(t, task, task_update)
}
