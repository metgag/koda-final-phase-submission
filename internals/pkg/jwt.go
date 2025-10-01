package pkg

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID int `json:"id"`
	jwt.RegisteredClaims
}

func NewJWTClaims(userId int) *Claims {
	return &Claims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			Issuer:    os.Getenv("JWT_ISSUER"),
		},
	}
}

func (c *Claims) GenToken() (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return "", errors.New("jwt secret's not found")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(jwtSecret))
}

func (c *Claims) VerifyToken(token string) error {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return errors.New("jwt secret's not found")
	}

	parsedToken, err := jwt.ParseWithClaims(
		token,
		c,
		func(t *jwt.Token) (any, error) {
			return []byte(jwtSecret), nil
		},
	)
	if err != nil {
		return err
	}
	if !parsedToken.Valid {
		return jwt.ErrTokenInvalidClaims
	}

	iss, err := parsedToken.Claims.GetIssuer()
	if err != nil {
		return err
	}
	if iss != os.Getenv("JWT_ISSUER") {
		return jwt.ErrTokenInvalidIssuer
	}

	return nil
}
