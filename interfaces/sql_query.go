package interfaces

type GetAller interface{
	GetAll(int) (interface{}, error)
}

type GetByIder interface{
	GetById(uint) (interface{}, error)
}

type EmailFieldExistser interface{
	EmailFieldExists(string) (bool, error)
}

type NicknameFieldExistser interface{
	NicknameFieldExists(string) (bool, uint, error)
}
