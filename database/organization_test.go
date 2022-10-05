package database_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/ochom/sdp-lib/models"
	"github.com/stretchr/testify/require"
)

func Test_impl_CreateOrganization(t *testing.T) {
	i := initDB(t)
	ctx := context.Background()
	data := &models.Organization{
		ID:   uuid.NewString(),
		Name: "test",
	}

	err := i.CreateOrganization(ctx, data)
	require.NoError(t, err)
}

func Test_impl_UpdateOrganization(t *testing.T) {
	i := initDB(t)
	ctx := context.Background()

	data := &models.Organization{
		ID:   uuid.NewString(),
		Name: "Test",
	}

	err := i.CreateOrganization(ctx, data)
	require.NoError(t, err)

	data.Name = "New Name"
	err = i.UpdateOrganization(ctx, data)
	require.NoError(t, err)

	got, err := i.GetOrganization(ctx, &models.Organization{ID: data.ID})
	require.NoError(t, err)
	require.Equal(t, data.Name, got.Name)
}

func Test_impl_DeleteOrganization(t *testing.T) {
	i := initDB(t)
	ctx := context.Background()

	data := &models.Organization{
		ID:   uuid.NewString(),
		Name: "test",
	}
	err := i.CreateOrganization(ctx, data)
	require.NoError(t, err)

	err = i.DeleteOrganization(ctx, &models.Organization{ID: data.ID})
	require.NoError(t, err)

	_, err = i.GetOrganization(ctx, &models.Organization{ID: data.ID})
	require.Error(t, err)
}

func Test_impl_GetOrganization(t *testing.T) {
	i := initDB(t)
	ctx := context.Background()

	data := &models.Organization{
		ID:   uuid.NewString(),
		Name: "test",
	}
	err := i.CreateOrganization(ctx, data)
	require.NoError(t, err)

	got, err := i.GetOrganization(ctx, &models.Organization{ID: data.ID})
	require.NoError(t, err)
	require.Equal(t, data.ID, got.ID)

	got, err = i.GetOrganization(ctx, &models.Organization{ID: "not-found"})
	require.Error(t, err)
	require.Nil(t, got)
}

func Test_impl_GetOrganizations(t *testing.T) {
	i := initDB(t)
	ctx := context.Background()

	existing, err := i.GetOrganizations(ctx, &models.Organization{})
	require.NoError(t, err)

	for _, e := range existing {
		err = i.DeleteOrganization(ctx, e)
		require.NoError(t, err)
	}

	data := []*models.Organization{
		{
			ID:   uuid.NewString(),
			Name: "test",
		},
		{
			ID:   uuid.NewString(),
			Name: "test",
		},
		{
			ID:   uuid.NewString(),
			Name: "test",
		},
	}

	for _, d := range data {
		err = i.CreateOrganization(ctx, d)
		require.NoError(t, err)
	}

	got, err := i.GetOrganizations(ctx, &models.Organization{})
	require.NoError(t, err)
	require.Len(t, got, len(data))
}

func Test_impl_CreateOffer(t *testing.T) {
	i := initDB(t)
	ctx := context.Background()
	data := &models.Offer{
		ID:   uuid.NewString(),
		Name: "test",
	}

	err := i.CreateOffer(ctx, data)
	require.NoError(t, err)

	got, err := i.GetOffer(ctx, &models.Offer{ID: data.ID})
	require.NoError(t, err)
	require.Equal(t, data.ID, got.ID)
}

func Test_impl_UpdateOffer(t *testing.T) {
	i := initDB(t)
	ctx := context.Background()

	data := &models.Offer{
		ID:   uuid.NewString(),
		Name: "Test",
	}

	err := i.CreateOffer(ctx, data)
	require.NoError(t, err)

	data.Name = "New Name"
	err = i.UpdateOffer(ctx, data)
	require.NoError(t, err)

	got, err := i.GetOffer(ctx, &models.Offer{ID: data.ID})
	require.NoError(t, err)
	require.Equal(t, data.Name, got.Name)
}

func Test_impl_DeleteOffer(t *testing.T) {
	i := initDB(t)
	ctx := context.Background()

	data := &models.Offer{
		ID:   uuid.NewString(),
		Name: "test",
	}
	err := i.CreateOffer(ctx, data)
	require.NoError(t, err)

	err = i.DeleteOffer(ctx, &models.Offer{ID: data.ID})
	require.NoError(t, err)

	_, err = i.GetOffer(ctx, &models.Offer{ID: data.ID})
	require.Error(t, err)
}

func Test_impl_GetOffer(t *testing.T) {
	i := initDB(t)
	ctx := context.Background()

	data := &models.Offer{
		ID:   uuid.NewString(),
		Name: "test",
	}
	err := i.CreateOffer(ctx, data)
	require.NoError(t, err)

	got, err := i.GetOffer(ctx, &models.Offer{ID: data.ID})
	require.NoError(t, err)
	require.Equal(t, data.ID, got.ID)

	got, err = i.GetOffer(ctx, &models.Offer{ID: "not-found"})
	require.Error(t, err)
	require.Nil(t, got)
}

func Test_impl_GetOffers(t *testing.T) {
	i := initDB(t)
	ctx := context.Background()

	existing, err := i.GetOffers(ctx, &models.Offer{})
	require.NoError(t, err)

	for _, e := range existing {
		err = i.DeleteOffer(ctx, e)
		require.NoError(t, err)
	}

	data := []*models.Offer{
		{
			ID:   uuid.NewString(),
			Name: "test",
		},
		{
			ID:   uuid.NewString(),
			Name: "test",
		},
		{
			ID:   uuid.NewString(),
			Name: "test",
		},
	}

	for _, d := range data {
		err = i.CreateOffer(ctx, d)
		require.NoError(t, err)
	}

	got, err := i.GetOffers(ctx, &models.Offer{})
	require.NoError(t, err)
	require.Len(t, got, len(data))
}
