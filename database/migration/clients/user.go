package clients

import "github.com/MikelSot/melody/interfaces"

type user struct {
	migrate interfaces.Migrater
}

func NewUser(m interfaces.Migrater) *user{
	return &user{m}
}

func (u *user) CreateTable() error{
	err := u.migrate.Migrate()
	if err != nil {
		return err
	}
	return nil
}