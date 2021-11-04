package dao

import (
	"fmt"
	"github.com/MikelSot/melody/domain"
)

type typeMessage struct {}

func NewTypeMessage() *typeMessage {
	return &typeMessage{}
}

func (t typeMessage) Create(data *interface{}) error {
	db.Create(&data)
	return nil
}

func (t typeMessage) GetById(id uint) (interface{}, error) {
	tm := domain.TypeMessage{}
	db.First(&tm, id)
	fmt.Println("SE EJECUTA -> GET BYID")
	return tm, nil
}
