package interfaces

type GetAller interface{
	GetAll(int) (interface{}, error)
}

type GetByIder interface{
	GetById(uint) (interface{}, error)
}

type Exister interface{
	Exists(string) (bool, error)
}
