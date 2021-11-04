package database

import (
	"github.com/MikelSot/melody/database/connection"
	"github.com/MikelSot/melody/database/migration"
	"github.com/MikelSot/melody/infrastructure"
)

func Database()  {
	infrastructure.NewEnv().LoadEnv()
	connection.NewGetEnvDB().GetEnvDataBase()
	connection.NewConnection().Connection()
	migration.Migration()
}
