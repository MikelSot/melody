package interfaces

// las interfaces deben recibir y enviar ppunteros (los que sonpara la base de datos y consultas)

type Creater interface {
	Create(*interface{}) error
}

type Updater interface {
	Update(uint, *interface{}) error
}

type DeleteSofter interface{
	DeleteSoft(uint) error
}

type DeletePermanenter interface {
	DeletePermanent(uint) error
}

type UpdateNicknameFielder interface {
	UpdateNicknameField(uint, string) error
}

type UpdatePictureFielder interface {
	UpdatePictureField(uint, string) error
}

type UpdateBoolFielder interface {
	UpdateBoolField(uint, bool) error
}
