package domain

type TypeMessage struct {
	ID       uint8  `gorm:"primaryKey" json:"id,omitempty"`
	Type     string `gorm:"uniqueIndex; type:varchar(60); not null" json:"type,omitempty"`
	Messages []Message
}

type TypeMessages []TypeMessage
