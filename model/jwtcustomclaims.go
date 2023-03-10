package model

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

// modelo del JWT
type JWTCustomClaims struct {
	UserID  uuid.UUID `json:"user_id"`
	Email   string    `json:"email"`
	IsAdmin bool      `json:"is_admin"`
	jwt.StandardClaims
}
