package database_test

import (
	"context"
	"os"
	"testing"

	"github.com/ochom/sdp-lib/database"
	"github.com/ochom/sdp-lib/models"
)

func Test_impl_CreateOrganization(t *testing.T) {
	os.Setenv("DATABASE_DNS", "db.sqlite")
	i, err := database.New(database.SQLite)
	if err != nil {
		t.Fatal(err)
	}

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
