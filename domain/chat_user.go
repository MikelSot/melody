package domain

type ChatUser struct {
	ChatID uint `gorm:"primaryKey"`
	UserID uint `gorm:"primaryKey"`
}
