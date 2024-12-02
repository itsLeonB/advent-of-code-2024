package utils

import "log"

func ErrorInputParse(err error) {
	log.Fatalf("error reading input: %s", err.Error())
}

func ErrorAtoi(err error) {
	log.Fatalf("error converting to int: %s", err.Error())
}
