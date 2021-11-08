package interfaces

import "github.com/MikelSot/melody/domain"

type PersistenceChatUser interface {
	Create(user *domain.ChatUser) error
}

type QueryChatUser interface {
	Match(myId uint, yourId uint) (domain.ChatUser,error)
	AllPeopleAddedToMyChat(max uint,myId uint) (domain.ChatUsers, error)
	AllGroupsAddedToMyChat(max uint,myId uint) (domain.ChatUsers, error)
	AllAddedToMyChat(max uint, myId uint) (domain.ChatUsers, error)
}

