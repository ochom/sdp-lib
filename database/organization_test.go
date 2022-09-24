package database_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/ochom/sdp-lib/models"
)

var orgID = uuid.NewString()

func Test_impl_CreateOrganization(t *testing.T) {
	i := initDB(t)
	type args struct {
		ctx  context.Context
		data *models.Organization
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "create organization",
			args: args{
				ctx: context.Background(),
				data: &models.Organization{
					ID:   orgID,
					Name: "test",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := i.CreateOrganization(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("impl.CreateOrganization() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_impl_UpdateOrganization(t *testing.T) {
	i := initDB(t)
	type args struct {
		ctx  context.Context
		data *models.Organization
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "update organization",
			args: args{
				ctx: context.Background(),
				data: &models.Organization{
					Name: "new name",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			org, err := i.GetOrganization(tt.args.ctx, &models.Organization{ID: orgID})
			if err != nil {
				t.Errorf("impl.UpdateOrganization() error = %v, wantErr %v", err, tt.wantErr)
			}

			org.Name = tt.args.data.Name
			if err := i.UpdateOrganization(tt.args.ctx, org); (err != nil) != tt.wantErr {
				t.Errorf("impl.UpdateOrganization() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_impl_GetOrganization(t *testing.T) {
	i := initDB(t)
	type args struct {
		ctx   context.Context
		query *models.Organization
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Organization
		wantErr bool
	}{
		{
			name: "get organization",
			args: args{
				ctx:   context.Background(),
				query: &models.Organization{ID: orgID},
			},
			want: &models.Organization{
				ID:   orgID,
				Name: "new name",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := i.GetOrganization(tt.args.ctx, tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("impl.GetOrganization() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.ID != tt.want.ID {
				t.Errorf("impl.GetOrganization() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_impl_GetOrganizations(t *testing.T) {
	i := initDB(t)
	type args struct {
		ctx   context.Context
		query *models.Organization
	}
	tests := []struct {
		name    string
		args    args
		want    []*models.Organization
		wantErr bool
	}{
		{
			name: "get organizations",
			args: args{
				ctx:   context.Background(),
				query: &models.Organization{},
			},
			want: []*models.Organization{
				{
					ID:   orgID,
					Name: "new name",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := i.GetOrganizations(tt.args.ctx, tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("impl.GetOrganizations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != len(tt.want) {
				t.Errorf("impl.GetOrganizations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_impl_DeleteOrganization(t *testing.T) {
	i := initDB(t)
	type args struct {
		ctx   context.Context
		query *models.Organization
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "delete organization",
			args: args{
				ctx:   context.Background(),
				query: &models.Organization{ID: orgID},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := i.DeleteOrganization(tt.args.ctx, tt.args.query); (err != nil) != tt.wantErr {
				t.Errorf("impl.DeleteOrganization() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := i.DeleteOrganization(tt.args.ctx, tt.args.query); (err != nil) != tt.wantErr {
				t.Errorf("impl.DeleteOrganization() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
