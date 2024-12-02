package utils

import (
	"log"
	"strconv"
)

func Atoi(val string) int {
	num, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalf("error converting to int: %s", err.Error())
	}

	return num
}
