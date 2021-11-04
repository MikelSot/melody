package migration

import (
	"github.com/MikelSot/melody/database/connection"
)

func Migration() {
	pl := connection.NewPool()
	NewCreateTables(pl).CreateTable()
	NewInsertData(pl).InsertData()
}
