package utils

import (
	"bufio"
	"log"
)

func ReadNewLine(reader *bufio.Reader) string {
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("error reading input: %s", err.Error())
	}

	return input[:len(input)-1] // Trim the trailing newline
}
