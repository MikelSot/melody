package dao

import (
	"github.com/MikelSot/melody/domain"
)

type user struct {}

func NewUser() *user{
	return &user{}
}

func (u user) Create(data *interface{}) error {
	db.Create(&data)
	return nil
}

func (u user) Update(id uint, data *interface{}) error {
	db.Model(id).Updates(&data)
	return nil
}

func (u user) DeleteSoft(id uint) error {
	us := domain.User{}
	us.ID = id
	db.Delete(&us)
	return nil
}

func (u user) DeletePermanent(id uint) error {
	us := domain.User{}
	us.ID = id
	db.Unscoped().Delete(&us)
	return nil
}






