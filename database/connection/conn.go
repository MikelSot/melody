package connection

import (
	"github.com/MikelSot/melody/interfaces"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sync"
)

var (
	db   *gorm.DB
	once sync.Once
)

type connection struct {
	conn interfaces.Connectioner
}

func NewConnection() *connection {
	return &connection{}
}

func (c connection) Connection() {
	once.Do(func() {
		var err error
		dns := c.getDnsConnection()
		db, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
		if err != nil {
			log.Fatalf("no se pudo conectar a la base de datos --> %v", err)
		}
		log.Println("CONECTADO!!")
	})
}

func (c connection) getDnsConnection() string {
	hs := "host=" + host
	ur := " user=" + user
	pss := " password=" + password
	dbn := " dbname=" + dbname
	p := " port=" + port
	tz := " sslmode=disable TimeZone=America/Lima"

	var dns string = hs + ur + pss + dbn + p + tz
	return dns
}

