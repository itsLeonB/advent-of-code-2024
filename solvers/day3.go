package solvers

import (
	"bufio"
	"regexp"

	"github.com/itsLeonB/advent-of-code-2024/utils"
)

func searchRegex(withEnabler bool) (*regexp.Regexp, *regexp.Regexp) {
	if withEnabler {
		return regexp.MustCompile(`mul\((\d+),(\d+)\)|do\(\)|don't\(\)`), regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	}

	return regexp.MustCompile(`mul\(\d+,\d+\)`), regexp.MustCompile(`mul\((\d+),(\d+)\)`)
}

func Day3Part1(input string) int {
	file := utils.OpenFile(input)
	defer file.Close()

	reAll, reMul := searchRegex(false)

	var total int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		matches := reAll.FindAllString(scanner.Text(), -1)

		for _, match := range matches {
			mul := calculateMul(reMul, match)
			total += mul
		}
	}

	return total
}

func Day3Part2(input string) int {
	file := utils.OpenFile(input)
	defer file.Close()

	reAll, reMul := searchRegex(true)

	var total int
	scanner := bufio.NewScanner(file)
	var isEnabled = true

	for scanner.Scan() {
		matches := reAll.FindAllString(scanner.Text(), -1)

		for _, match := range matches {
			switch match {
			case "do()":
				isEnabled = true
			case "don't()":
				isEnabled = false
			default:
				if isEnabled {
					mul := calculateMul(reMul, match)
					total += mul
				}
			}
		}
	}

	return total
}

func calculateMul(reMul *regexp.Regexp, match string) int {
	matches := reMul.FindStringSubmatch(match)
	x := utils.Atoi(matches[1])
	y := utils.Atoi(matches[2])

	return x * y
}
