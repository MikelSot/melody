package interfaces

import (
	"database/sql"
	"gorm.io/gorm"
)

type Pooler interface{
	Pool() *sql.DB
}

type Connectioner interface{
	Connection() *gorm.DB
}