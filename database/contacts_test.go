package database_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/ochom/sdp-lib/models"
	"github.com/stretchr/testify/require"
)

func Test_impl_CreateContactGroup(t *testing.T) {
	i := initDB(t)
	ctx := context.Background()
	data := &models.ContactGroup{
		ID:   uuid.NewString(),
		Name: "test",
	}

	err := i.CreateContactGroup(ctx, data)
	require.NoError(t, err)
}

func Test_impl_UpdateContactGroup(t *testing.T) {
	i := initDB(t)
	ctx := context.Background()

	data := &models.ContactGroup{
		ID:   uuid.NewString(),
		Name: "Test",
	}

	err := i.CreateContactGroup(ctx, data)
	require.NoError(t, err)

	data.Name = "New Name"
	err = i.UpdateContactGroup(ctx, data)
	require.NoError(t, err)

	got, err := i.GetContactGroup(ctx, &models.ContactGroup{ID: data.ID})
	require.NoError(t, err)
	require.Equal(t, data.Name, got.Name)
}

func Test_impl_DeleteContactGroup(t *testing.T) {
	i := initDB(t)
	ctx := context.Background()

	data := &models.ContactGroup{
		ID:   uuid.NewString(),
		Name: "test",
	}
	err := i.CreateContactGroup(ctx, data)
	require.NoError(t, err)

	err = i.DeleteContactGroup(ctx, data)
	require.NoError(t, err)

	got, err := i.GetContactGroup(ctx, &models.ContactGroup{ID: data.ID})
	require.Error(t, err)
	require.Nil(t, got)
}

func Test_impl_GetContactGroup(t *testing.T) {
	i := initDB(t)
	ctx := context.Background()

	data := &models.ContactGroup{
		ID:   uuid.NewString(),
		Name: "test",
	}
	err := i.CreateContactGroup(ctx, data)
	require.NoError(t, err)

	got, err := i.GetContactGroup(ctx, &models.ContactGroup{ID: data.ID})
	require.NoError(t, err)
	require.Equal(t, data.ID, got.ID)

	got, err = i.GetContactGroup(ctx, &models.ContactGroup{ID: "not-found"})
	require.Error(t, err)
	require.Nil(t, got)
}

func Test_impl_GetContactGroups(t *testing.T) {
	i := initDB(t)
	ctx := context.Background()

	existing, err := i.GetContactGroups(ctx, &models.ContactGroup{})
	require.NoError(t, err)

	for _, e := range existing {
		err = i.DeleteContactGroup(ctx, e)
		require.NoError(t, err)
	}

	data := []models.ContactGroup{
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
		err := i.CreateContactGroup(ctx, &d)
		require.NoError(t, err)
	}

	got, err := i.GetContactGroups(ctx, &models.ContactGroup{})
	require.NoError(t, err)
	require.Len(t, got, len(data))
}

func Test_impl_CreateContact(t *testing.T) {
	i := initDB(t)
	ctx := context.Background()
	data := &models.Contact{
		ID:     uuid.NewString(),
		Mobile: "1234567890",
	}
	err := i.CreateContact(ctx, data)
	require.NoError(t, err)
}

func Test_impl_UpdateContact(t *testing.T) {
	i := initDB(t)
	ctx := context.Background()

	data := &models.Contact{
		ID:     uuid.NewString(),
		Mobile: "1234567890",
	}
	err := i.CreateContact(ctx, data)
	require.NoError(t, err)

	data.Mobile = "0987654321"
	err = i.UpdateContact(ctx, data)
	require.NoError(t, err)

	got, err := i.GetContact(ctx, &models.Contact{ID: data.ID})
	require.NoError(t, err)
	require.Equal(t, data.Mobile, got.Mobile)
}

func Test_impl_DeleteContact(t *testing.T) {
	i := initDB(t)
	ctx := context.Background()

	data := &models.Contact{
		ID:     uuid.NewString(),
		Mobile: "1234567890",
	}
	err := i.CreateContact(ctx, data)
	require.NoError(t, err)

	err = i.DeleteContact(ctx, data)
	require.NoError(t, err)

	got, err := i.GetContact(ctx, &models.Contact{ID: data.ID})
	require.Error(t, err)
	require.Nil(t, got)
}

func Test_impl_GetContact(t *testing.T) {
	i := initDB(t)
	ctx := context.Background()

	data := &models.Contact{
		ID:     uuid.NewString(),
		Mobile: "1234567890",
	}
	err := i.CreateContact(ctx, data)
	require.NoError(t, err)

	got, err := i.GetContact(ctx, &models.Contact{ID: data.ID})
	require.NoError(t, err)
	require.Equal(t, data.ID, got.ID)

	got, err = i.GetContact(ctx, &models.Contact{ID: "not-found"})
	require.Error(t, err)
	require.Nil(t, got)
}

func Test_impl_GetContacts(t *testing.T) {
	i := initDB(t)
	ctx := context.Background()

	existing, err := i.GetContacts(ctx, &models.Contact{})
	require.NoError(t, err)

	for _, e := range existing {
		err = i.DeleteContact(ctx, e)
		require.NoError(t, err)
	}

	data := []models.Contact{
		{
			ID:     uuid.NewString(),
			Mobile: "1234567890",
		},
		{
			ID:     uuid.NewString(),
			Mobile: "1234567890",
		},
	}

	for _, d := range data {
		err := i.CreateContact(ctx, &d)
		require.NoError(t, err)
	}

	got, err := i.GetContacts(ctx, &models.Contact{})
	require.NoError(t, err)
	require.Len(t, got, len(data))
}
