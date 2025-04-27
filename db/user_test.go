package db

import (
	"context"
	"log"
	"testing"

	"github.com/luke_design_pattern/util"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	// Username, Email, Firstname, Lastname
	arg := CreateUserParams{
		Username:  util.RandomOwner(),
		Email:     util.RandomEmail(),
		Firstname: util.RandomOwner(),
		Lastname:  util.RandomOwner(),
		Password:  util.RandomString(10),
	}

	users, err := testStoreOra.CreateUser(context.Background(), arg)
	log.Printf("arg ==> %v	", arg)
	log.Printf("users ora ==> %v	", users)
	log.Printf("check log error ora ==> %v	", err)
	// require.NoError(t, err, "check error")
	require.NotEmpty(t, users)
	// require.Equal(t, arg, users) // must be treated deepEqual with map{}
	require.Equal(t, arg.Username, users.Username)
	require.Equal(t, arg.Email, users.Email)
	require.Equal(t, arg.Firstname, users.Firstname)
	require.Equal(t, arg.Lastname, users.Lastname)

	users, err = testStorePG.CreateUser(context.Background(), arg)
	log.Printf("arg ==> %v	", arg)
	log.Printf("users pg ==> %v	", users)

	log.Printf("check log error pg ==> %v	", err)
	require.NoError(t, err, "check error")
	// require.NotEmpty(t, users)
	// require.Equal(t, arg, users) // must be treated deepEqual with map{}
	require.Equal(t, arg.Username, users.Username)
	require.Equal(t, arg.Email, users.Email)
	require.Equal(t, arg.Firstname, users.Firstname)
	require.Equal(t, arg.Lastname, users.Lastname)
}
