package dao

import "github.com/MikelSot/melody/domain"

type message struct{}

func NewMessage() *message {
	return &message{}
}

func (m message) Create(data *interface{}) error {
	db.Create(&data)
	return nil
}

func (m message) DeleteSoft(id uint) error {
	ms := domain.Message{}
	ms.ID = id
	db.Delete(&ms)
	return nil
}

func (m message) UpdateBoolField(id uint, state bool) error {
	db.Table("messages").Where("id=?", id).Update("star", state)
	return nil
}

// updateFile, not picture
func (m message) UpdatePictureField(id uint, archive string) error {
	db.Table("messages").Where("id=?", id).Update("archive", archive)
	return nil
}