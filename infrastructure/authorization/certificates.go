package authorization

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"sync"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
	once       sync.Once
)

type loaJWTdFile struct{}

func NewLoaJWTdFile() *loaJWTdFile {
	return &loaJWTdFile{}
}

func (l loaJWTdFile) loadFile(privateFile, publicFile string) error {
	privateBytes, err := ioutil.ReadFile(privateFile)
	if err != nil {
		return err
	}

	publicBytes, err := ioutil.ReadFile(publicFile)
	if err != nil {
		return err
	}

	return l.parseRSA(privateBytes, publicBytes)
}

func (l loaJWTdFile) parseRSA(privateBytes, publicBytes []byte) error {
	var err error
	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		return err
	}
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		return err
	}
	return nil
}

func (l loaJWTdFile) LoadJWTFile(priv, pub string) error {
	var err error
	once.Do(func() {
		err = l.loadFile(priv, pub)
	})
	return err
}
