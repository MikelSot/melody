package interfaces

import "github.com/MikelSot/melody/domain"

type PersistenceUser interface {
	Create(*domain.CreateUser) error
	Update(*domain.UpdateUser) error
	UpdatePictureField(*domain.FieldString) error
	UpdateNicknameField(*domain.FieldString) error
	UpdateOnlineField(*domain.State) error
}

type PersistenceChater interface {
	Create(*domain.Chat) error
	Update(*domain.Chat) error
	GetById(uint) (domain.Chat, error)
}

type PersistenceChatUser interface {
	Create(user *domain.ChatUser) error
}


type PersistenceMessage interface {
	Create(*domain.Message) error
	UpdateStarField(id uint) error
	DeleteMessage(id uint) error
}
