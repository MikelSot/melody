package connection

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

var once sync.Once
type connection struct {
	db   *sql.DB
}

func NewConnection() *connection {
	return &connection{}
}

func (c connection) connection() {
	once.Do(func() {
		var err error
		dns := c.getDnsConnection()
		c.db, err = sql.Open("postgres", dns)
		if err != nil {
			log.Fatalf("ERROR CONNECCION -> %v", err)
		}

		if err = c.db.Ping(); err != nil {
			log.Fatalf("ERROR PING-> %v", err)
		}
		log.Println("CONNECTADO!!")
	})
}

func (c connection) getDnsConnection() string {
	var dns string = "postgres://"+user+":"+password+"@"+host+":"+port+"/"+dbname+"?sslmode=disable"
	return dns
}

func (c connection) Connection(){
	c.connection()
}


func (c connection) Pool() *sql.DB {
	log.Println("ctmr")
	return c.db
}

