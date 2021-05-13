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
		Shortname:   util.RandomShortname(),
		Problemname: util.RandomProblemname(),
		Content:     util.RandomContent(),
		Subtasks:    Subtasks_tmp,
		Answers:     util.RandomAnswers(int(Subtasks_tmp)),
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
	arg := ListTasksParams{
		Limit:  5,
		Offset: 5,
	}
	tasks_list, err := testQueries.ListTasks(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, tasks_list, 5)

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

func TestUpdateTaskShortname(t *testing.T) {
	task := CreateRandomTask(t)
	arg := UpdateShortnameParams{
		ID:        task.ID,
		Shortname: util.RandomShortname(),
	}
	task_update, err := testQueries.UpdateShortname(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, task_update.Shortname, arg.Shortname)
	task.Shortname = task_update.Shortname
	require.Equal(t, task, task_update)
}

func TestUpdateTaskProblemname(t *testing.T) {
	task := CreateRandomTask(t)
	arg := UpdateProblemnameParams{
		ID:          task.ID,
		Problemname: util.RandomProblemname(),
	}
	task_update, err := testQueries.UpdateProblemname(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, task_update.Problemname, arg.Problemname)
	task.Problemname = task_update.Problemname
	require.Equal(t, task, task_update)
}

func TestUpdateTaskContent(t *testing.T) {
	task := CreateRandomTask(t)
	arg := UpdateContentParams{
		ID:      task.ID,
		Content: util.RandomContent(),
	}
	task_update, err := testQueries.UpdateContent(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, task_update.Content, arg.Content)
	task.Content = task_update.Content
	require.Equal(t, task, task_update)
}

func TestUpdateTaskSubtasks(t *testing.T) {
	task := CreateRandomTask(t)
	arg := UpdateSubtasksParams{
		ID:       task.ID,
		Subtasks: util.RandomSubtasks(),
	}
	task_update, err := testQueries.UpdateSubtasks(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, task_update.Subtasks, arg.Subtasks)
	task.Subtasks = task_update.Subtasks
	require.Equal(t, task, task_update)
}

func TestUpdateTaskAnswers(t *testing.T) {
	task := CreateRandomTask(t)
	arg := UpdateAnswersParams{
		ID:      task.ID,
		Answers: util.RandomAnswers(int(task.Subtasks)),
	}
	task_update, err := testQueries.UpdateAnswers(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, task_update.Answers, arg.Answers)
	task.Answers = task_update.Answers
	require.Equal(t, task, task_update)
}
