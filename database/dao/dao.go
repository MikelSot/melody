package dao

import (
	"github.com/MikelSot/melody/database/connection"
	"gorm.io/gorm"
)

var db *gorm.DB

func Dao() {
	db = connection.NewPool().Pool()
}