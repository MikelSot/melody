package interfaces

import "github.com/MikelSot/melody/domain"

type User interface {
	PersistenceUser
	QueryUser
}


type PersistenceUser interface {
	Create(*domain.CreateUser) error
	Update(*domain.UpdateUser) error
	UpdateNicknameField(*domain.FieldString) error
	UpdatePictureField(*domain.FieldString) error
	UpdateOnlineField(*domain.State) error
}

type QueryUser interface {
	GetAllNotAddedUsers(int, uint) (*domain.Users, error)
	GetById(uint) (*domain.User, error)
	EmailFieldExists(string) (bool,*domain.User,error)
	NicknameFieldExists(string) (bool, *domain.User, error)
}
