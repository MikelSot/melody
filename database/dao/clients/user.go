package clients

import (
	"github.com/MikelSot/melody/domain"
	"github.com/MikelSot/melody/interfaces"
)

type user struct {
	persistence interfaces.PersistenceUser
}

func NewUser(p interfaces.PersistenceUser) *user{
	return &user{p}
}

func (u *user) Create(m *domain.CreateUser) error{
	return u.persistence.Create(m)
}

func (u user) Update(m *domain.UpdateUser) error{
	return u.persistence.Update(m)
}

func (u user) UpdateNicknameField(m *domain.FieldString) error{
	return u.persistence.UpdateNicknameField(m)
}

func (u user) UpdatePictureField(m *domain.FieldString) error{
	return u.persistence.UpdatePictureField(m)
}

func (u user) UpdateOnlineField(m *domain.State) error{
	return u.persistence.UpdateOnlineField(m)
}
