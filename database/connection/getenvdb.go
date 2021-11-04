package connection

import (
	"github.com/MikelSot/melody/interfaces"
	"os"
)

var (
	host     string
	user     string
	password string
	dbname   string
	port     string
)

type getEnvDB struct {
	env interfaces.GetEnvDataBaser
}

func (l *getEnvDB) GetEnvDataBase() {
	host = os.Getenv("DB_HOST")
	user = os.Getenv("DB_USERNAME")
	password = os.Getenv("DB_PASSWORD")
	dbname = os.Getenv("DB_NAME")
	port = os.Getenv("DB_PORT")
}

func NewGetEnvDB() *getEnvDB {
	return &getEnvDB{}
}
