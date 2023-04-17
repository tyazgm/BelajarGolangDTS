package helper

import (
	"challenge2/model"

	"github.com/google/uuid"
	// "os/exec"

	// "github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func GenerateID() string {
	return uuid.New().String()
}

func GenerateToken(userID string) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
	})

	var tokenString string

	tokenString, err := jwtToken.SignedString([]byte("4l0h4m0r4"))

	return tokenString, err
}

func Hash(plain string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(plain), 8)
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func IsHashValid(hash, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))

	return err == nil
}

func VerifyToken(token string) (*jwt.Token, error) {
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, model.ErrorInvalidToken
		}

		return []byte("4l0h4m0r4"), nil
	})

	return jwtToken, err
}
