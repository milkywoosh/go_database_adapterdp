package db

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	// Username, Email, Firstname, Lastname
	arg := CreateUserParams{
		Username:  "toni",
		Email:     "toni@gmail.com",
		Firstname: "toni1",
		Lastname:  "toni2",
	}

	users, err := testStore.CreateUser(context.Background(), arg)
	log.Printf("arg ==> %v	", arg)
	log.Printf("check log ==> %v	", err)
	require.NoError(t, err, "check error")
	require.NotEmpty(t, users)
	// require.Equal(t, arg, users) // must be treated deepEqual with map{}
	require.Equal(t, arg.Username, users.Username)
	require.Equal(t, arg.Email, users.Email)
	require.Equal(t, arg.Firstname, users.Firstname)
	require.Equal(t, arg.Lastname, users.Lastname)
}
