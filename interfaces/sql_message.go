package interfaces

import "github.com/MikelSot/melody/domain"

type PersistenceMessage interface {
	Create(*domain.Message) error
	UpdateStarField(id uint) error
	DeleteMessage(id uint) error
}

type QueryMessage interface {
	GetMessageFromChat(max, myId uint) (domain.Messages, error)
	GetAllStarred() (domain.Messages, error)
	GetLatestChatMessageFromPeopleAdded(chatID uint) (domain.LastMessageFromPeoples, error)
	GetLatestChatMessageFromGroupAdded(chatID uint) (domain.LastMessageFromGroups, error)
	GetLatestChatMessageFromAllAdded(chatID uint) (domain.LastMessageFromAlls, error)
}

