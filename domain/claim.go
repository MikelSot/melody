package domain

import "github.com/dgrijalva/jwt-go"

type Claim struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}
