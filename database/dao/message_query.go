package dao

import "github.com/MikelSot/melody/domain"

func (m message) GetById(id uint) (interface{}, error) {
	ms := domain.Message{}
	db.First(&ms, id)
	return ms, nil
}


