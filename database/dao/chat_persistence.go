package dao

import "github.com/MikelSot/melody/domain"

type chat struct {}

func NewChat() *chat{
	return &chat{}
}

func (c chat) Create(data *interface{}) error {
	db.Create(&data)
	return nil
}

func (c chat) DeleteSoft(id uint) error {
	ch := domain.Chat{}
	ch.ID = id
	db.Delete(&ch)
	return nil
}

func (c chat) Update(id uint, data *interface{}) error {
	db.Model(id).Updates(&data)
	return nil
}