package authorization

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type myJWT struct {}

func NewMyJWT() *myJWT {
	return &myJWT{}
}

func (m myJWT) GenerateToken(id uint) (string, error){
	claim := Claim{
		ID:             id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour*2).Unix(),
			Issuer: "Melody",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (m myJWT) ValidateToken(t string) (Claim, error){
	token, err := jwt.ParseWithClaims(t, &Claim{},m.getPublicKey)
	if err != nil {
		return Claim{},err
	}
	if !token.Valid{
		return Claim{}, errors.New("token no valido")
	}
	claim, ok := token.Claims.(*Claim)
	if !ok{
		return Claim{}, errors.New("no se puedo obtener los claim")
	}
	return *claim, nil
}

func (m myJWT) getPublicKey (t *jwt.Token) (interface{}, error){
	return publicKey, nil
}