package interfaces

import "github.com/MikelSot/melody/infrastructure/authorization"

type LoadJWTFiler interface {
	LoadJWTFile(pwiv, pub string) error
}

type JWT interface {
	GenerateToken(id uint) (string, error)
	ValidateToken(t string) (authorization.Claim, error)
}