package pkg

import (
	"log"

	"github.com/alexedwards/argon2id"
)

func CreateHash(password string) string {

	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		log.Fatal(err)
	}

	return hash
}

func ComparePassword(password, hash string) bool {

	match, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		log.Fatal(err)
	}
	return match
}
