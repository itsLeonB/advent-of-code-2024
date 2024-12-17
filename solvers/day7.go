package solvers

import (
	"bufio"
	"strings"

	"github.com/itsLeonB/advent-of-code-2024/utils"
)

func Day7Part1(input string) int {
	file := utils.OpenFile(input)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var total int

	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Split(line, ": ")

		target := utils.Atoi(splits[0])
		values := strings.Fields(splits[1])

		vals := make([]int, len(values))

		for i := range len(values) {
			value := utils.Atoi(values[i])
			vals[i] = value
		}

		if canReachTarget(vals, target) {
			total += target
		}
	}

	return total
}

func canReachTarget(nums []int, target int) bool {
	var backtrack func([]int, int) bool

	backtrack = func(remaining []int, current int) bool {
		if len(remaining) == 0 {
			return current == target
		}

		for i, num := range remaining {
			newRemaining := make([]int, 0, len(remaining)-1)
			newRemaining = append(newRemaining, remaining[:i]...)
			newRemaining = append(newRemaining, remaining[i+1:]...)

			if backtrack(newRemaining, current+num) {
				return true
			}

			if backtrack(newRemaining, current*num) {
				return true
			}
		}

		return false
	}

	for i, num := range nums {
		remaining := make([]int, 0, len(nums)-1)
		remaining = append(remaining, nums[:i]...)
		remaining = append(remaining, nums[i+1:]...)

		if backtrack(remaining, num) {
			return true
		}
	}

	return false
}
