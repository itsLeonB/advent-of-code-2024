package solvers

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/itsLeonB/advent-of-code-2024/utils"
)

func Day1Part1(input string) int {
	file, err := os.Open(input)
	if err != nil {
		log.Panicf("error opening file: %s", err.Error())
	}
	defer file.Close()

	col1 := make([]int, 0)
	col2 := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		cols := strings.Fields(line)

		val1, err := strconv.Atoi(cols[0])
		if err != nil {
			log.Panicf("error converting to int: %s", err.Error())
		}

		val2, err := strconv.Atoi(cols[1])
		if err != nil {
			log.Panicf("error converting to int: %s", err.Error())
		}

		col1 = utils.InsertAndSort(col1, val1)
		col2 = utils.InsertAndSort(col2, val2)
	}

	len1 := len(col1)
	len2 := len(col2)
	if len1 != len2 {
		log.Panic("two cols are not the same size")
	}

	var total int
	for i := 0; i < len(col1); i++ {
		diff := col1[i] - col2[i]
		if diff < 0 {
			diff *= -1
		}

		total += diff
	}

	return total
}
