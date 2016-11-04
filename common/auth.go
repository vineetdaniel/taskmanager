package common

import (
	"io/ioutil"
)

const (
	privKeyPath = "keys/app.rsa"
	pubKeyPath  = "keys/app.rsa.pub"
)

var (
	verifyKey, signKey []byte
)

//Read the key files before starting http handlers

func initKeys() {
	var err error

	signKey, err = ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatalf("[initkeys]: %s\n", err)
	}
	verifyKey, err = iotuil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatalf("[initkeys]: %s\n", err)
	}
}
