package clients

import "github.com/MikelSot/melody/interfaces"

type typeMessage struct {
	migrate interfaces.Migrater
}

func NewTypeMessage(m interfaces.Migrater) *typeMessage {
	return &typeMessage{m}
}

func (u *typeMessage) CreateTable() error {
	err := u.migrate.Migrate()
	if err != nil {
		return err
	}
	return nil
}