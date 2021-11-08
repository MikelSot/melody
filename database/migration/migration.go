package migration

import (
	"database/sql"
	"github.com/MikelSot/melody/database/migration/clients"
	"github.com/MikelSot/melody/database/migration/tables"
	"log"
	"sync"
)

var once sync.Once

type migration struct {
	db *sql.DB
}

func NewMigration(db *sql.DB) *migration {
	return &migration{  db}
}

func (m migration) Migration(){
	once.Do(func(){
		var err error
		u := clients.NewUser(tables.NewUser(m.db)); err = u.CreateTable()
		c := clients.NewChat(tables.NewChat(m.db)); err = c.CreateTable()
		cu := clients.NewChatUser(tables.NewChatUsers(m.db)); err = cu.CreateTable()
		m := clients.NewMessage(tables.NewMessage(m.db)); err = m.CreateTable()
		if err != nil {
			log.Fatalf("ERROR MIGRATE -> %v", err)
		}
		log.Println("MIGRADO!!")
	})
}