package main

import (
	"github.com/MikelSot/melody/database"
	"github.com/MikelSot/melody/infrastructure/authorization"
	"github.com/MikelSot/melody/infrastructure/routes"
	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
	"net/http"
)

//type tmGetById interfaces.GetAller

func main() {
	pool := database.Database()
	authorization.NewAuthorization(authorization.NewLoaJWTdFile()).LoadJwtFile(
		"certificates/app.rsa",
		"certificates/app.rsa.public",
	)
	r := mux.NewRouter().StrictSlash(true)
	router := routes.NewRouter(r, pool, authorization.NewMyJWT())
	router.User()

	log.Fatal(http.ListenAndServe(":3000", r))
}



