package pkg

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// claims sama dengan payload
type Payload struct {
	Email string
	jwt.RegisteredClaims
}

func NewPayload(email string) *Payload {
	return &Payload{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    os.Getenv("JWT_ISSUER"),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		}}
}

func (item *Payload) CreateToken() (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, item.RegisteredClaims)
	result, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return result, nil
}

func (item *Payload) VerifyToken(token string) (*Payload, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	parsedToken, err := jwt.ParseWithClaims(token, &Payload{}, func(t *jwt.Token) (interface{}, error) {
		// interface{} sama dengan any
		return []byte(jwtSecret), nil
	})
	if err != nil{
		return nil, err
	}
	payload := parsedToken.Claims.(*Payload)
	return payload, nil
}
