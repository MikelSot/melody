package migration

import (
	"github.com/MikelSot/melody/domain"
	"github.com/MikelSot/melody/interfaces"
	"log"
)

type createTables struct {
	ct interfaces.CreateTabler
	pl interfaces.Pooler
}

func NewCreateTables(pl interfaces.Pooler) *createTables {
	return &createTables{pl:pl}
}

func (t createTables) CreateTable() error {
	db := t.pl.Pool()
	err := db.AutoMigrate(
			&domain.User{},
			&domain.TypeMessage{},
			&domain.Chat{},
			&domain.Message{},
			&domain.ChatUser{},
		)

	if err != nil {
		return err
	}

	log.Println("Tablas creadas")
	return nil
}

