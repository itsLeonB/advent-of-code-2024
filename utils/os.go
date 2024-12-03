package utils

import (
	"log"
	"os"
)

func OpenFile(input string) *os.File {
	file, err := os.Open(input)
	if err != nil {
		log.Fatalf("error opening file: %s", err.Error())
	}

	return file
}

func ReadFile(input string) []byte {
	bytes, err := os.ReadFile(input)
	if err != nil {
		log.Fatalf("error reading file: %s", err.Error())
	}

	return bytes
}
