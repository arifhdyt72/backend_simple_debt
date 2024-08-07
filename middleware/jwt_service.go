package middleware

import (
	"backend_test_debt/models"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(data models.User) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

var SECRETKEY = []byte(os.Getenv("SECRET_KEY"))

func NewJWTService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(data models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": data.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(SECRETKEY))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(SECRETKEY), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, fmt.Errorf("invalid signature: %v", err)
		}
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, fmt.Errorf("token is malformed: %v", err)
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, fmt.Errorf("token is either expired or not active yet: %v", err)
			} else {
				return nil, fmt.Errorf("couldn't handle this token: %v", err)
			}
		}
		return nil, fmt.Errorf("couldn't parse token: %v", err)
	}

	return token, nil
}
