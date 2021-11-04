package migration

import (
	"github.com/MikelSot/melody/domain"
	"github.com/MikelSot/melody/interfaces"
	"log"
)

type insertData struct {
	insdt  interfaces.InsertDataer
	pl interfaces.Pooler
}

func NewInsertData(pl interfaces.Pooler) *insertData {
	return &insertData{pl:pl}
}

func (i *insertData) InsertData() error {
	typeMessage := domain.TypeMessages{
		{Type: "texto"},
		{Type: "documento"},
		{Type: "imagen"},
		{Type: "audio"},
		{Type: "video"},
	}

	db := i.pl.Pool()
	db.Create(&typeMessage)
	log.Println("Insercion de datos hecha")
	return nil
}