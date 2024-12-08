package solvers

import (
	"sort"
	"strings"

	"github.com/itsLeonB/advent-of-code-2024/utils"
)

func Day5Part1(input string) int {
	rules, lines := parseInput(input)
	var count int

	for _, line := range lines {
		if validateLine(rules, line) {
			mid := line[len(line)/2]
			count += utils.Atoi(mid)
		}
	}

	return count
}

func Day5Part2(input string) int {
	rules, lines := parseInput(input)
	var count int

	for _, line := range lines {
		if !validateLine(rules, line) {
			reordered := reorderLine(rules, line)
			mid := reordered[len(reordered)/2]
			count += utils.Atoi(mid)
		}
	}

	return count
}

func reorderLine(rules []string, line []string) []string {
	reorderedLine := make([]string, len(line))
	copy(reorderedLine, line)

	sort.Slice(reorderedLine, func(i, j int) bool {
		for k := 0; k < len(rules)-1; k += 2 {
			if reorderedLine[i] == rules[k] && reorderedLine[j] == rules[k+1] {
				return true
			}
		}
		return false
	})

	return reorderedLine
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

func parseInput(input string) ([]string, [][]string) {
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

	return rules, lines
}
