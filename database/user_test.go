package database_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/ochom/sdp-lib/models"
	"github.com/stretchr/testify/require"
)

func Test_impl_CreateUser(t *testing.T) {
	i := initDB(t)
	ctx := context.Background()
	data := &models.User{
		ID:        uuid.NewString(),
		FirstName: "test",
	}

	err := i.CreateUser(ctx, data)
	require.NoError(t, err)
}

func Test_impl_UpdateUser(t *testing.T) {
	i := initDB(t)
	ctx := context.Background()

	data := &models.User{
		ID:        uuid.NewString(),
		FirstName: "Test",
	}

	err := i.CreateUser(ctx, data)
	require.NoError(t, err)

	data.FirstName = "New FirstName"
	err = i.UpdateUser(ctx, data)
	require.NoError(t, err)

	got, err := i.GetUser(ctx, &models.User{ID: data.ID})
	require.NoError(t, err)
	require.Equal(t, data.FirstName, got.FirstName)
}

func Test_impl_DeleteUser(t *testing.T) {
	i := initDB(t)
	ctx := context.Background()

	data := &models.User{
		ID:        uuid.NewString(),
		FirstName: "test",
	}
	err := i.CreateUser(ctx, data)
	require.NoError(t, err)

	err = i.DeleteUser(ctx, &models.User{ID: data.ID})
	require.NoError(t, err)

	_, err = i.GetUser(ctx, &models.User{ID: data.ID})
	require.Error(t, err)
}

func Test_impl_GetUser(t *testing.T) {
	i := initDB(t)
	ctx := context.Background()

	data := &models.User{
		ID:        uuid.NewString(),
		FirstName: "test",
	}
	err := i.CreateUser(ctx, data)
	require.NoError(t, err)

	got, err := i.GetUser(ctx, &models.User{ID: data.ID})
	require.NoError(t, err)
	require.Equal(t, data.ID, got.ID)

	got, err = i.GetUser(ctx, &models.User{ID: "not-found"})
	require.Error(t, err)
	require.Nil(t, got)
}

func Test_impl_GetUsers(t *testing.T) {
	i := initDB(t)
	ctx := context.Background()

	existing, err := i.GetUsers(ctx, &models.User{})
	require.NoError(t, err)

	for _, e := range existing {
		err = i.DeleteUser(ctx, e)
		require.NoError(t, err)
	}

	data := []*models.User{
		{
			ID:        uuid.NewString(),
			FirstName: "test",
		},
		{
			ID:        uuid.NewString(),
			FirstName: "test",
		},
		{
			ID:        uuid.NewString(),
			FirstName: "test",
		},
	}

	for _, d := range data {
		err = i.CreateUser(ctx, d)
		require.NoError(t, err)
	}

	got, err := i.GetUsers(ctx, &models.User{})
	require.NoError(t, err)
	require.Len(t, got, len(data))
}
