package solvers

import (
	"strings"

	"github.com/itsLeonB/advent-of-code-2024/utils"
)

func Day5Part1(input string) int {
	file := utils.ReadFile(input)
	content := string(file)
	groups := strings.Split(content, "\n\n")

	var rules []string
	var lines [][]string

	ruleTexts := strings.Split(groups[0], "\n")
	for _, rule := range ruleTexts {
		ruleVals := strings.Split(rule, "|")
		rules = append(rules, ruleVals...)
	}

	lineTexts := strings.Split(groups[1], "\n")
	for _, lineText := range lineTexts {
		lines = append(lines, strings.Split(lineText, ","))
	}

	var count int

	for _, line := range lines {
		if validateLine(rules, line) {
			mid := line[len(line)/2]
			count += utils.Atoi(mid)
		}
	}

	return count
}

func validateLine(rules []string, line []string) bool {
	for i := 0; i < len(line); i++ {
		first := line[i]
		for j := i + 1; j < len(line); j++ {
			second := line[j]
			if !validatePair(rules, first, second) {
				return false
			}
		}
	}

	return true
}

func validatePair(rules []string, first string, second string) bool {
	for i := 0; i < len(rules)-1; i += 2 {
		if rules[i] == first && rules[i+1] == second {
			return true
		}
	}

	return false
}
