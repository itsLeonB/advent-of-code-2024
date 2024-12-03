package solvers

import (
	"bufio"
	"regexp"

	"github.com/itsLeonB/advent-of-code-2024/utils"
)

func Day3Part1(input string) int {
	file := utils.OpenFile(input)
	defer file.Close()

	reAll := regexp.MustCompile(`mul\(\d+,\d+\)`)
	reMul := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	var total int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		matches := reAll.FindAllString(scanner.Text(), -1)

		for _, match := range matches {
			matches := reMul.FindStringSubmatch(match)
			x := utils.Atoi(matches[1])
			y := utils.Atoi(matches[2])
			total += x * y
		}
	}

	return total
}
