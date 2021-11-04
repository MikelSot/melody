package dao

import "github.com/MikelSot/melody/domain"

func (u user) GetById(id uint) (interface{}, error) {
	us := domain.User{}
	db.First(&us, id)
	return us, nil
}

func (u user) GetAll(max int) (interface{}, error) {
	us := domain.Users{}
	db.Limit(max).Find(&us)
	return us, nil
}
