package main

import (
	"github.com/gbrlsnchs/jwt"
	"log"
	"time"
)

func main() {
	now := time.Now()
	hs256 := jwt.NewHS256("secret")
	jot := &jwt.JWT{
		Issuer:         "gbrlsnchs",
		Subject:        "someone",
		Audience:       "gophers",
		ExpirationTime: now.Add(24 * 30 * 12 * time.Hour).Unix(),
		NotBefore:      now.Add(30 * time.Minute).Unix(),
		IssuedAt:       now.Unix(),
		ID:             "foobar",
	}
	jot.SetAlgorithm(hs256)
	jot.SetKeyID("kid")
	payload, err := jwt.Marshal(jot)
	if err != nil {
		// handle error
	}
	token, err := hs256.Sign(payload)
	if err != nil {
		// handle error
	}
	log.Printf("token = %s", token)
}
