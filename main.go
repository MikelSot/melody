package main

import (
	"fmt"
	"github.com/MikelSot/melody/database"
	"github.com/MikelSot/melody/database/dao"
	"github.com/MikelSot/melody/interfaces"
)

type tmGetById interfaces.GetAller

func main() {
	database.Database()
	tmGetById := dao.NewTypeMessage()
	//var data domain.TypeMessage
	data, _:= tmGetById.GetById(2)
	fmt.Printf("%+v",data)
}
