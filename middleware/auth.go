package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/ochom/sdp-lib/models"
	"github.com/ochom/sdp-lib/utils"
)

type ctxKey string

const (
	ctxUser ctxKey = "user"
)

// secretKey ...
var secretKey = utils.GetEnvOrDefault("JWT_SECRET", "secret")

// SignedDetails ...
type SignedDetails struct {
	ID           string              `json:"id,omitempty"`
	Email        string              `json:"email,omitempty"`
	Mobile       string              `json:"Mobile,omitempty"`
	Organization models.Organization `json:"organization,omitempty"`
	jwt.StandardClaims
}

// Token ...
type Token struct {
	AccessToken  string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

// GenerateAuthToken ...
func GenerateAuthToken(user *models.User, org *models.Organization) (*Token, error) {
	claims := &SignedDetails{
		ID:           user.ID,
		Email:        user.Email,
		Mobile:       user.Mobile,
		Organization: *org,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
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

// Authenticated ...
func Authenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "no auth header"})
			return
		}

		token := authHeader[7:]
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "no token"})
			return
		}

		claim, err := validateToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		user := models.User{
			ID:             claim.ID,
			Mobile:         claim.Mobile,
			Email:          claim.Email,
			OrganizationID: claim.Organization.ID,
		}

		authContext := AuthenticatedContext(c, &user)
		c.Request = c.Request.WithContext(authContext)

		c.Next()
	}
}

func validateToken(token string) (*SignedDetails, error) {
	claims := &SignedDetails{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return claims, nil
}

// GetContextUser gets user  from context returns nil if no user in context
func GetContextUser(ctx context.Context) *models.User {
	if ctx.Value(ctxUser) == nil {
		return nil
	}

	user, ok := ctx.Value(ctxUser).(*models.User)
	if !ok {
		return nil
	}

	return user
}

// AuthenticatedContext ....
func AuthenticatedContext(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, ctxUser, user)
}
