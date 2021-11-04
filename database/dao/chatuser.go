package dao

type chatUser struct {}

func NewChatUser() *chatUser {
	return &chatUser{}
}

func (c chatUser) Create(data *interface{}) error {
	db.Create(&data)
	return nil
}
