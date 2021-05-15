package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomStr(20)
	passwordEncoded, err := HassPassword(password)

	require.NoError(t, err)
	err = CheckPassword(password, passwordEncoded)
	require.NoError(t, err)

	invalidpassword := RandomStr(20)
	err = CheckPassword(invalidpassword, passwordEncoded)
	require.Error(t, err)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

}
