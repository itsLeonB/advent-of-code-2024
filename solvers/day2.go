package solvers

import (
	"bufio"
	"strings"

	"github.com/itsLeonB/advent-of-code-2024/utils"
)

func Day2Part1(input string) int {
	file := utils.OpenFile(input)
	defer file.Close()

	var safeCount int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		vals := strings.Fields(line)

		if isSafeSequence(vals, false) {
			safeCount++
		}
	}

	return safeCount
}

func Day2Part2(input string) int {
	file := utils.OpenFile(input)
	defer file.Close()

	var safeCount int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		vals := strings.Fields(line)

		if isSafeSequence(vals, true) {
			safeCount++
		}
	}

	return safeCount
}

func isSafeSequence(vals []string, tolerateOnce bool) bool {
	num1 := utils.Atoi(vals[1])
	num2 := utils.Atoi(vals[0])
	isIncreasing := num1 > num2

	for i := 1; i < len(vals); i++ {
		num1 := utils.Atoi(vals[i])
		num2 := utils.Atoi(vals[i-1])
		diff := num1 - num2

		if (isIncreasing && diff <= 0) || (!isIncreasing && diff >= 0) || utils.Abs(diff) < 1 || utils.Abs(diff) > 3 {
			if !tolerateOnce {
				return false
			}

			tolerateOnce = false
		}
	}

	return true
}
