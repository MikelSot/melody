package main

import (
	"github.com/MikelSot/melody/database"
	"github.com/MikelSot/melody/infrastructure/authorization"
	"github.com/MikelSot/melody/infrastructure/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

func main() {
	pool := database.Database()
	authorization.NewAuthorization(authorization.NewLoaJWTdFile()).LoadJwtFile(
		"certificates/app.rsa",
		"certificates/app.rsa.public",
	)
	r := mux.NewRouter().StrictSlash(true)
	router := routes.NewRouter(r, pool, authorization.NewMyJWT())
	router.User()

	port := os.Getenv("PORT_SERVER")
	handler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut},
		AllowCredentials: true,
	}).Handler(r)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}



