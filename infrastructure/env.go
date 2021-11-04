package infrastructure

import (
	"github.com/MikelSot/melody/interfaces"
	"github.com/joho/godotenv"
	"log"
)

type env struct{
	envData interfaces.LoadEnver
}

func (e env) LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("ERROR al cargar las variables de entorno ->", err)
	}
}

func NewEnv() *env {
	return &env{}
}