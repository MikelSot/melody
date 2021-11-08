package interfaces

import "github.com/MikelSot/melody/domain"

type PersistenceChater interface {
	Create(*domain.Chat) error
	Update(*domain.Chat) error
	GetById(uint) (domain.Chat, error)
}
