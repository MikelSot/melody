package interfaces

import (
	"github.com/MikelSot/melody/domain"
)

type LoadJWTFiler interface {
	LoadJWTFile(priv, pub string) error
}

type JWT interface {
	GenerateToken(id uint) (string, error)
	ValidateToken(t string) (domain.Claim, error)
}