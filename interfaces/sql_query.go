package interfaces

import "github.com/MikelSot/melody/domain"

type QueryUser interface {
	GetAllNotAddedUsers(int, uint) (domain.Users, error)
	GetById(uint) (domain.User, error)
	EmailFieldExists(string) (domain.User,error)
	NicknameFieldExists(string) (domain.User, error)
}


type QueryChatUser interface {
	Match(myId uint, yourId uint) (domain.ChatUser,error)
	AllPeopleAddedToMyChat(max uint,myId uint) (domain.ChatUsers, error)
	AllGroupsAddedToMyChat(max uint,myId uint) (domain.ChatUsers, error)
	AllAddedToMyChat(max uint, myId uint) (domain.ChatUsers, error)
}

type QueryMessage interface {
	GetMessageFromChat(max, myId uint) (domain.Messages, error)
	GetAllStarred() (domain.Messages, error)
	GetLatestChatMessageFromPeopleAdded(chatID uint) (domain.LastMessageFromPeoples, error)
	GetLatestChatMessageFromGroupAdded(chatID uint) (domain.LastMessageFromGroups, error)
	GetLatestChatMessageFromAllAdded(chatID uint) (domain.LastMessageFromAlls, error)
}

