package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thanhqt2002/hackathon/db/util"
)

func CreateRandomUser(t *testing.T) User {
	PasswordEncoded, err := util.HassPassword(util.RandomPass())
	require.NoError(t, err)
	arg := CreateUserParams{
		Username:        util.RandomUsr(),
		Fullname:        util.RandomFullName(),
		PasswordEncoded: PasswordEncoded,
		Usertype:        util.RandomUserType(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.PasswordEncoded, user.PasswordEncoded)
	require.Equal(t, arg.Usertype, user.Usertype)
	require.Equal(t, arg.Fullname, user.Fullname)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
	return user
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := CreateRandomUser(t)
	user_get, err := testQueries.GetUser(context.Background(), user.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user_get)
	require.Equal(t, user, user_get)
}

func TestListUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomUser(t)
	}
	arg := ListUsersParams{
		Limit:  5,
		Offset: 5,
	}
	users_list, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users_list, 5)

	for _, user := range users_list {
		require.NotEmpty(t, user)
	}
}
func TestUpdateUserFullname(t *testing.T) {
	user := CreateRandomUser(t)
	arg := UpdateUserFullnameParams{
		ID:       user.ID,
		Fullname: util.RandomFullName(),
	}
	user_update, err := testQueries.UpdateUserFullname(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, user_update.Fullname, arg.Fullname)
	user.Fullname = user_update.Fullname
	require.Equal(t, user, user_update)
}

func TestDeleteUser(t *testing.T) {
	user := CreateRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user.ID)
	require.NoError(t, err)
	user_get, err := testQueries.GetUser(context.Background(), user.Username)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user_get)
}
