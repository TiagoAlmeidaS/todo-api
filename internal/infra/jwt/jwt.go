package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"todo_project.com/internal/app/security"
)

type Payload struct {
	ID        string
	Name      string
	ExpiresAt time.Time
}

func (payload Payload) Valid() error {
	if time.Now().After(payload.ExpiresAt) {
		return security.ErrUnauthorized
	}
	return nil
}

type Authenticator struct {
	secret   string
	duration time.Duration
}

func NewAuthenticator(secret string, duration time.Duration) *Authenticator {
	return &Authenticator{
		secret:   secret,
		duration: duration,
	}
}

func (a *Authenticator) Generate(user security.User) (string, error) {
	payload := Payload{
		ID:        user.ID,
		Name:      user.Name,
		ExpiresAt: time.Now().Add(a.duration),
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := at.SignedString([]byte(a.secret))
	if err != nil {
		return "", security.ErrUnauthorized
	}

	return token, nil
}

func (a *Authenticator) Validate(token string) (*security.User, error) {
	payload := Payload{}
	_, err := jwt.ParseWithClaims(token, &payload, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, security.ErrUnauthorized
		}
		return []byte(a.secret), nil
	})
	if err != nil {
		return nil, security.ErrUnauthorized
	}

	return &security.User{
		ID:   payload.ID,
		Name: payload.Name,
	}, nil
}

func (a *Authenticator) RefreshToken(token string) (string, error) {
	payload := Payload{}
	_, err := jwt.ParseWithClaims(token, &payload, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			if payload.ID == "" || payload.Name == "" {
				return "", security.ErrUnauthorized
			}

			newToken, err := a.Generate(security.User{ID: payload.ID, Name: payload.Name})
			if err != nil {
				return "", security.ErrUnauthorized
			}

			return newToken, nil
		}
		return []byte(a.secret), nil
	})
	if err != nil {
		return "", security.ErrUnauthorized
	}
	return token, nil
}
