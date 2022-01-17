package utility

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var tokenKey = []byte(os.Getenv("API_JWT_KEY"))

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(username string) (string, time.Time, error) {
	expiredTime := time.Now().Add(60 * time.Minute)
	expiredAt := expiredTime.Unix()

	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredAt,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(tokenKey)

	return tokenString, expiredTime, err
}

func DecodeToken(tokenStr string) (string, error) {
	claims := &Claims{}

	_, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return tokenKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", err
		}
	}

	if claims.Username == "" {
		return "", err
	}

	return claims.Username, nil
}
