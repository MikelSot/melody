package dao

import (
	"github.com/MikelSot/melody/database/connection"
)

func Dao() {
	connection.NewConnection().Pool()
}