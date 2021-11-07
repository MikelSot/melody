package clients

import "github.com/MikelSot/melody/interfaces"

type chat struct {
	migrate interfaces.Migrater
}

func NewChat(m interfaces.Migrater) *chat {
	return &chat{m}
}

func (u *chat) CreateTable() error{
	err := u.migrate.Migrate()
	if err != nil {
		return err
	}
	return nil
}
