package dao

import "github.com/MikelSot/melody/domain"

const (
	maxLimit = 1
)

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

func (u user) NicknameFieldExists(nickname string) (bool, uint, error) {
	us := domain.User{}
	data := db.Limit(maxLimit).Select("ID").Find(&us, "nickname= ?", nickname)
	if data.RowsAffected == maxLimit  {
		return true, us.ID, nil
	}
	return false, 0, nil
}

func (u user) EmailFieldExists(email string) (bool, error) {
	us := domain.User{}
	data := db.Limit(maxLimit).Select("email").Find(&us, "email= ?", email)
	if data.RowsAffected == maxLimit  {
		return true, nil
	}
	return false, nil
}

// TODO: filtar todos los grupos, chat con otros contactos, y todo en donde yo este [consulta hecha, solo faltamigrara]
