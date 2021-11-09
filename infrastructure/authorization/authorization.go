package authorization

import (
	"github.com/MikelSot/melody/interfaces"
	"log"
)

type authorization struct {
	LoadJWTFile interfaces.LoadJWTFiler
}

func NewAuthorization(load interfaces.LoadJWTFiler) *authorization {
	return &authorization{load}
}

func (a authorization) LoadJwtFile(private, public string)  {
	err := a.LoadJWTFile.LoadJWTFile(private, public)
	if err != nil {
		log.Fatalf("ERROR LOAF JWT FILE -> %v", err)
	}
}
