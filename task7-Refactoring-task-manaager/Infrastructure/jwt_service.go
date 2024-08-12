package infrastructure

import (
	"fmt"
	domain "task7-Refactoring-task-manaager/Domain"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService struct {
	jwt_secret []byte
}

func (j *JWTService) GenerateToken(user domain.User) (string, error) {
	j.jwt_secret = []byte("It has to be a secret")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	})
	token.Claims = jwt.MapClaims{
		"_id":      user.ID,
		"username": user.Username,
		"role":     user.Role,
	}

	jwt_token, err := token.SignedString(j.jwt_secret)
	if err != nil {
		return "", err
	}
	return jwt_token, nil
}

func (j *JWTService) CheckToken(authPart string) (*jwt.Token, error) {

	token, err := jwt.Parse(authPart, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.jwt_secret, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (j *JWTService) FindClaim(token *jwt.Token) (jwt.MapClaims, bool) {
	claims, ok := token.Claims.(jwt.MapClaims)
	return claims, ok
}
