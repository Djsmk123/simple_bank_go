package db

import (
	"context"
	"testing"
	"time"

	"github.com/djsmk123/simplebank/utils"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	createRandomUser(t)

}
func createRandomUser(t *testing.T) User {
	hashPassword, err := utils.HashPassword(utils.RandomString(6))
	require.NoError(t, err)

	args := CreateUserParams{
		Username:       utils.RandomOwner(),
		Email:          utils.RandomEmail(),
		HashedPassword: hashPassword,
		FullName:       utils.RandomOwner(),
	}
	user, err := testQueries.CreateUser(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, args.Email, user.Email)
	require.Equal(t, args.HashedPassword, user.HashedPassword)
	require.Equal(t, args.FullName, user.FullName)
	require.NotZero(t, user.CreatedAt)
	require.True(t, user.PasswordChangedAt.IsZero())
	return user
}

//func getAccount

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)

	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)

	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}
