package clients

import (
	"github.com/MikelSot/melody/interfaces"
)

type message struct {
	migrate interfaces.Migrater
}

func NewMessage(m interfaces.Migrater) *message {
	return &message{m}
}

func (u *message) CreateTable() error{
	err := u.migrate.Migrate()
	if err != nil {
		return err
	}
	return nil
}
