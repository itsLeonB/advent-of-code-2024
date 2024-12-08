package solvers

import (
	"strings"

	"github.com/itsLeonB/advent-of-code-2024/utils"
)

type direction struct {
	x int
	y int
}

var directions = []direction{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}
var runes = []rune{'X', 'M', 'A', 'S'}

func Day4Part1(input string) int {
	filteredLines, maxX, maxY := readInput(input)
	var count int

	for i := 0; i < maxY; i++ {
		for j := 0; j < maxX; j++ {
			for _, dir := range directions {
				if searchPart1(filteredLines, j, i, 0, dir) {
					count++
				}
			}
		}
	}

	return count
}

func Day4Part2(input string) int {
	filteredLines, maxX, maxY := readInput(input)
	var count int

	for i := 1; i < maxY-1; i++ {
		for j := 1; j < maxX-1; j++ {
			if filteredLines[i][j] == 'A' && searchPart2(filteredLines, j, i) {
				count++
			}
		}
	}

	return count
}

func searchPart1(haystack [][]rune, x, y, runeIdx int, dir direction) bool {
	if runeIdx == 4 {
		return true
	}

	maxX := len(haystack[0])
	maxY := len(haystack)

	if x < 0 || x >= maxX || y < 0 || y >= maxY {
		return false
	}

	if haystack[y][x] == runes[runeIdx] {
		isFound := searchPart1(haystack, x+dir.x, y+dir.y, runeIdx+1, dir)
		return isFound
	}

	return false
}

func searchPart2(haystack [][]rune, x, y int) bool {
	if haystack[y-1][x-1] == 'S' && haystack[y-1][x+1] == 'S' && haystack[y+1][x-1] == 'M' && haystack[y+1][x+1] == 'M' {
		return true
	}

	if haystack[y-1][x-1] == 'M' && haystack[y-1][x+1] == 'S' && haystack[y+1][x-1] == 'M' && haystack[y+1][x+1] == 'S' {
		return true
	}

	if haystack[y-1][x-1] == 'M' && haystack[y-1][x+1] == 'M' && haystack[y+1][x-1] == 'S' && haystack[y+1][x+1] == 'S' {
		return true
	}

	if haystack[y-1][x-1] == 'S' && haystack[y-1][x+1] == 'M' && haystack[y+1][x-1] == 'S' && haystack[y+1][x+1] == 'M' {
		return true
	}

	return false
}

func readInput(input string) ([][]rune, int, int) {
	file := utils.ReadFile(input)
	lines := strings.Split(string(file), "\n")

	var filteredLines [][]rune
	for _, line := range lines {
		if len(strings.TrimSpace(line)) > 0 {
			filteredLines = append(filteredLines, []rune(line))
		}
	}

	maxX := len(filteredLines[0])
	maxY := len(filteredLines)

	return filteredLines, maxX, maxY
}
