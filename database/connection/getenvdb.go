package connection

import (
	"os"
)

var (
	host     string
	user     string
	password string
	dbname   string
	port     string
)

type getEnvDB struct {}

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
