package utils

import (
	"context"
	"crypto/rand"
	"math/big"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/ochom/sdp-lib/models"
	"golang.org/x/crypto/bcrypt"
)

type ctxKey string

const (
	ctxUser ctxKey = "user"
)

// secretKey ...
var secretKey = GetEnvOrDefault("JWT_SECRET", "secret")

type signedDetails struct {
	User models.User
	jwt.StandardClaims
}

// Token ...
type Token struct {
	AccessToken  string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

// HashPassword ...
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// ComparePassword ...
func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

//GenerateOTP generates a 6 figure OTP
func GenerateOTP(size int) string {
	otp := make([]byte, size)

	letters := "0123456789"
	max := big.NewInt(int64(len(letters)))
	for i := 0; i < size; i++ {
		num, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "123456"
		}
		otp[i] = letters[num.Int64()]
	}

	return string(otp)
}

// GenerateAuthToken ...
func GenerateAuthToken(user *models.User) (*Token, error) {
	claims := &signedDetails{
		User: *user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	refreshClaims := &signedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	return &Token{
		AccessToken:  token,
		RefreshToken: refreshToken,
	}, nil
}

// CreateAuthContext ....
func CreateAuthContext(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, ctxUser, user)
}

// ValidateToken ...
func ValidateToken(token string) (*models.User, error) {
	claims := &signedDetails{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return &claims.User, nil
}

// AuthenticatedUser gets user  from context returns nil if no user in context
func AuthenticatedUser(ctx context.Context) *models.User {
	if ctx.Value(ctxUser) == nil {
		return nil
	}

	user, ok := ctx.Value(ctxUser).(*models.User)
	if !ok {
		return nil
	}

	return user
}
