package utils

import "log"

func ErrorInputParse(err error) {
	log.Fatalf("error reading input: %s", err.Error())
}
