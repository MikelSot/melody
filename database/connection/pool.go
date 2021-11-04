package connection

import (
	"github.com/MikelSot/melody/interfaces"
	"gorm.io/gorm"
)

type pool struct {
	pool interfaces.Pooler
}

func NewPool() *pool {
	return &pool{}
}

func (c pool) Pool() *gorm.DB {
	return db
}

