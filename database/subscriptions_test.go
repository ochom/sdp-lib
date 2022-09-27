package database_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/ochom/sdp-lib/models"
	"github.com/stretchr/testify/require"
)

func Test_impl_CreateSubscriber(t *testing.T) {
	i := initDB(t)
	ctx := context.Background()
	data := &models.Subscriber{
		ID:        uuid.NewString(),
		FirstName: "test",
	}

	err := i.CreateSubscriber(ctx, data)
	require.NoError(t, err)
}

func Test_impl_DeleteSubscriber(t *testing.T) {
	i := initDB(t)
	ctx := context.Background()

	data := &models.Subscriber{
		ID:        uuid.NewString(),
		FirstName: "test",
	}
	err := i.CreateSubscriber(ctx, data)
	require.NoError(t, err)

	err = i.DeleteSubscriber(ctx, &models.Subscriber{ID: data.ID})
	require.NoError(t, err)

	got, err := i.GetSubscriber(ctx, &models.Subscriber{ID: data.ID})
	require.Error(t, err)
	require.Nil(t, got)
}

func Test_impl_GetSubscriber(t *testing.T) {
	i := initDB(t)
	ctx := context.Background()

	data := &models.Subscriber{
		ID:        uuid.NewString(),
		FirstName: "test",
	}
	err := i.CreateSubscriber(ctx, data)
	require.NoError(t, err)

	got, err := i.GetSubscriber(ctx, &models.Subscriber{ID: data.ID})
	require.NoError(t, err)
	require.Equal(t, data.ID, got.ID)

	got, err = i.GetSubscriber(ctx, &models.Subscriber{ID: "not-found"})
	require.Error(t, err)
	require.Nil(t, got)
}

func Test_impl_GetSubscribers(t *testing.T) {
	i := initDB(t)
	ctx := context.Background()

	existing, err := i.GetSubscribers(ctx, &models.Subscriber{})
	require.NoError(t, err)

	for _, e := range existing {
		err = i.DeleteSubscriber(ctx, &models.Subscriber{ID: e.ID})
		require.NoError(t, err)
	}

	data := []*models.Subscriber{
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
		err = i.CreateSubscriber(ctx, d)
		require.NoError(t, err)
	}

	got, err := i.GetSubscribers(ctx, &models.Subscriber{})
	require.NoError(t, err)
	require.Len(t, got, len(data))
}
