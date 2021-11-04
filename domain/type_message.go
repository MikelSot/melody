package domain

type TypeMessage struct {
	ID       uint8  `gorm:"primaryKey" json:"id,omitempty"`
	Message  string `gorm:"uniqueIndex; type:varchar(60); not null" json:"message,omitempty"`
	Messages []Message
}
