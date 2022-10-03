package utils

import (
	"context"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/ochom/sdp-lib/models"
	"github.com/stretchr/testify/require"
)

func getUser(t *testing.T) *models.User {
	return &models.User{
		ID:             uuid.NewString(),
		FirstName:      "John",
		LastName:       "Doe",
		Email:          "john.doe@mail.com",
		Mobile:         "1234567890",
		OrganizationID: uuid.NewString(),
	}
}

func TestHashPassword(t *testing.T) {
	password := "helloWorld#123"
	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)
	require.NotEqual(t, password, hashedPassword)
}

func TestComparePassword(t *testing.T) {
	password := "helloWorld#123"
	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)
	ok := ComparePassword(hashedPassword, password)
	require.True(t, ok)
}

func TestGenerateOTP(t *testing.T) {
	otp := GenerateOTP(6)
	require.Len(t, otp, 6)
}

func TestGenerateAuthToken(t *testing.T) {
	user := getUser(t)

	token, err := GenerateAuthToken(user)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	authUser, err := ValidateToken(token.AccessToken)
	require.NoError(t, err)
	require.NotEmpty(t, authUser)

	authUser, err = ValidateToken(token.RefreshToken)
	require.NoError(t, err)
	require.NotEmpty(t, authUser)
}

func TestCreateAuthContext(t *testing.T) {
	user := getUser(t)
	ctx := CreateAuthContext(context.Background(), user)
	authUser := GetAuthUser(ctx)
	require.Equal(t, user, authUser)
}

func TestValidateToken(t *testing.T) {
	user := getUser(t)
	token, err := GenerateAuthToken(user)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	type args struct {
		token string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.User
		wantErr bool
	}{
		{
			name: "valid token",
			args: args{
				token: token.AccessToken,
			},
			want:    user,
			wantErr: false,
		},
		{
			name: "invalid token",
			args: args{
				token: token.RefreshToken[:len(token.RefreshToken)-1],
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateToken(tt.args.token)
			require.Equal(t, tt.wantErr, err != nil)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestGetAuthUser(t *testing.T) {
	user := getUser(t)
	ctx := context.Background()

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
		want *models.User
	}{
		{
			name: "invalid context",
			args: args{
				ctx: context.Background(),
			},
			want: nil,
		},
		{
			name: "valid context",
			args: args{
				ctx: CreateAuthContext(ctx, user),
			},
			want: user,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAuthUser(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAuthUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
