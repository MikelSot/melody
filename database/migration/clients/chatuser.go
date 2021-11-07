package clients

import "github.com/MikelSot/melody/interfaces"

type chatUser struct {
	migrate interfaces.Migrater
}

func NewChatUser(m interfaces.Migrater) *chatUser {
	return &chatUser{m}
}

func (c *chatUser) CreateTable() error{
	err := c.migrate.Migrate()
	if err != nil {
		return err
	}
	return nil
}
