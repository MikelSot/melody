package main

import (
	"github.com/MikelSot/melody/database/connection"
	"github.com/MikelSot/melody/database/dao/clients"
	"github.com/MikelSot/melody/database/dao/persistence"
	"github.com/MikelSot/melody/domain"
	"github.com/MikelSot/melody/infrastructure"
	"log"
)

//type tmGetById interfaces.GetAller

func main() {
	//database.Database()

	infrastructure.NewEnv().LoadEnv()
	connection.NewGetEnvDB().GetEnvDataBase()
	c := connection.NewConnection(); c.Connection()
	u := persistence.NewUser(c.Pool())
	user := clients.NewUser(u)
	us := &domain.CreateUser{
		Name:     "mikel sanana",
		LastName: "soto",
		Email:    "quienmaspi@pi",
		Password: "123456",
	}

	up := &domain.UpdateUser{
		ID:          15,
		Name:        "soto",
		LastName:    "sanatana",
		Email:       "apolaya@djbnfkdj",
		Nickname:    "contreras",
		Description: "jajaja",
	}

	nick := &domain.FieldString{
		ID:    18,
		Field: "miandohuevada@hn",
	}

	pick := &domain.FieldString{
		ID:    19,
		Field: "/////////////////////comer",
	}

	on := &domain.State{
		ID:    10,
		Field: false,
	}

	err := user.Create(us)
	err = user.Update(up)
	err = user.UpdateNicknameField(nick)
	err = user.UpdatePictureField(pick)
	err = user.UpdateOnlineField(on)
	if err != nil {
		log.Println("ERROR", err)
	}
}



