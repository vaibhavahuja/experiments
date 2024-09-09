package main

import (
	"crypto"
	"crypto/sha256"

	_ "github.com/ldclabs/cose/key/hmac"
)

func init() {
	crypto.RegisterHash(crypto.SHA256, sha256.New)
}

func main() {
	ExampleMac0Message1()
	//FormPayload()
}
