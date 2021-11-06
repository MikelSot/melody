package dao

type chatUser struct {}

func NewChatUser() *chatUser {
	return &chatUser{}
}

func (c chatUser) Create(data *interface{}) error {
	db.Create(&data)
	return nil
}

// TODO: consultar si la union que va crear ya existe (es decir si el el id user y el ide chat ya son una clave primeria en la tabla) [x]
//TODO: abtener todos los uaurios a los que nmp tengo agregado [x]
