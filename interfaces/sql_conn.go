package interfaces

import (
	"database/sql"
)

type Pooler interface{
	Pool() *sql.DB
}

type Connectioner interface{
	Connection()
}

type Migrater interface {
	Migrate() error
}