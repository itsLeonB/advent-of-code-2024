package solvers

import (
	"bufio"

	"github.com/itsLeonB/advent-of-code-2024/utils"
)

type patrolDirection struct {
	x int
	y int
}

func Day6Part1(input string) int {
	file := utils.OpenFile(input)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var grid [][]rune
	var currentX, currentY int

	for scanner.Scan() {
		line := []rune(scanner.Text())
		for i, ch := range line {
			if ch == '^' {
				currentX = i
				currentY = len(grid)
			}
		}
		grid = append(grid, line)
	}

	maxX, maxY := len(grid[0]), len(grid)

	directions := []patrolDirection{
		{0, -1}, {1, 0}, {0, 1}, {-1, 0},
	}

	var uniqueCount int
	dirIdx := 0
	currentX, currentY = stepForward(currentX, currentY, directions[dirIdx])

	for currentY < maxY && currentX < maxX && currentX >= 0 && currentY >= 0 {
		currentCell := grid[currentY][currentX]

		if currentCell == '#' {
			currentX, currentY = stepBehind(currentX, currentY, directions[dirIdx])
			dirIdx = (dirIdx + 1) % len(directions)
			currentX, currentY = stepForward(currentX, currentY, directions[dirIdx])
			continue
		}

		if currentCell == '.' {
			uniqueCount++
			grid[currentY][currentX] = 'X'
		}

		currentX, currentY = stepForward(currentX, currentY, directions[dirIdx])
	}

	return uniqueCount + 1
}

func stepForward(x, y int, dir patrolDirection) (int, int) {
	return x + dir.x, y + dir.y
}

func stepBehind(x, y int, dir patrolDirection) (int, int) {
	return x - dir.x, y - dir.y
}
