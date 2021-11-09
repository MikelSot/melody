package database

import (
	"database/sql"
	"github.com/MikelSot/melody/database/connection"
	"github.com/MikelSot/melody/database/migration"
	"github.com/MikelSot/melody/infrastructure"
)

func Database() *sql.DB {
	infrastructure.NewEnv().LoadEnv()
	connection.NewGetEnvDB().GetEnvDataBase()
	c := connection.NewConnection(); c.Connection()
	migration.NewMigration(c.Pool()).Migration()
	return c.Pool()
}
