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
