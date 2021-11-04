package main

import (
	"github.com/joho/godotenv"
	"log"
)

func main() {



}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}