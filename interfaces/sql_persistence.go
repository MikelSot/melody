package interfaces

type Creater interface {
	Create(interface{}) error
}

type Updater interface {
	Update(uint, interface{}) error
}

type DeleteSofter interface{
	DeleteSoft(uint) error
}

type DeletePermanenter interface {
	DeletePermanent(uint) error
}

