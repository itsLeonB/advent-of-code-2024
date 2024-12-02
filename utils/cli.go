package utils

import "log"

func CleanInput(input string) string {
	return input[:len(input)-1] // Trim the trailing newline
}

func InputError(err error) {
	log.Fatalf("Error reading input: %s", err.Error())
}
