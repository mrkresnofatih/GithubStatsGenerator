package utils

import (
	errors "errors"
	log "log"
)

func HandleIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func BuildError(message string) error {
	e := errors.New(message)
	return e
}
