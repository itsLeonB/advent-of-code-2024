package solvers

import (
	"bufio"
	"fmt"

	"github.com/itsLeonB/advent-of-code-2024/utils"
)

type patrolDirection struct {
	x int
	y int
}

func NewPatrolDirection() []patrolDirection {
	return []patrolDirection{
		{0, -1}, {1, 0}, {0, 1}, {-1, 0},
	}
}

func SetupGrid(input string) ([][]rune, int, int, int, int) {
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

	return grid, currentX, currentY, maxX, maxY
}

func Day6Part1(input string) int {
	grid, currentX, currentY, maxX, maxY := SetupGrid(input)
	directions := NewPatrolDirection()

	var uniqueCount int
	dirIdx := 0
	currentX, currentY = stepForward(currentX, currentY, directions[dirIdx])

	for currentY < maxY && currentX < maxX && currentX >= 0 && currentY >= 0 {
		currentCell := grid[currentY][currentX]

		if currentCell == '#' {
			currentX, currentY = stepBehind(currentX, currentY, directions[dirIdx])
			dirIdx = (dirIdx + 1) % len(directions)
		} else if currentCell == '.' {
			uniqueCount++
			grid[currentY][currentX] = 'X'
		}

		currentX, currentY = stepForward(currentX, currentY, directions[dirIdx])
	}

	return uniqueCount + 1
}

func Day6Part2(input string) int {
	grid, currentX, currentY, _, _ := SetupGrid(input)
	directions := NewPatrolDirection()

	loopCount := 0
	dirIdx := 0
	testedObstacles := make(map[string]bool)

	for currentY >= 0 && currentY < len(grid) && currentX >= 0 && currentX < len(grid[0]) {
		nextX, nextY := stepForward(currentX, currentY, directions[dirIdx])
		obstacleKey := fmt.Sprintf("%d,%d", nextX, nextY)

		if nextY >= 0 && nextY < len(grid) && nextX >= 0 && nextX < len(grid[0]) && !testedObstacles[obstacleKey] {
			originalCell := grid[nextY][nextX]
			grid[nextY][nextX] = '#'

			if causesLoop(grid, currentX, currentY, dirIdx, directions) {
				loopCount++
			}

			grid[nextY][nextX] = originalCell
			testedObstacles[obstacleKey] = true
		}

		currentX, currentY = stepForward(currentX, currentY, directions[dirIdx])

		if currentY < 0 || currentY >= len(grid) || currentX < 0 || currentX >= len(grid[0]) {
			break
		}

		currentCell := grid[currentY][currentX]

		if currentCell == '#' {
			currentX, currentY = stepBehind(currentX, currentY, directions[dirIdx])
			dirIdx = (dirIdx + 1) % len(directions)
			currentX, currentY = stepForward(currentX, currentY, directions[dirIdx])
		}
	}

	return loopCount
}

func causesLoop(grid [][]rune, startX, startY, dirIdx int, directions []patrolDirection) bool {
	visited := make(map[string]bool)
	x, y := startX, startY

	for y >= 0 && y < len(grid) && x >= 0 && x < len(grid[0]) {
		state := stateKey(x, y, dirIdx)

		if visited[state] {
			return true
		}

		visited[state] = true

		cell := grid[y][x]

		if cell == '#' {
			x, y = stepBehind(x, y, directions[dirIdx])
			dirIdx = (dirIdx + 1) % len(directions)
			x, y = stepForward(x, y, directions[dirIdx])
			continue
		}

		x, y = stepForward(x, y, directions[dirIdx])
	}

	return false
}

func stateKey(x, y, dirIdx int) string {
	return fmt.Sprintf("%d,%d,%d", x, y, dirIdx)
}

func stepForward(x, y int, dir patrolDirection) (int, int) {
	return x + dir.x, y + dir.y
}

func stepBehind(x, y int, dir patrolDirection) (int, int) {
	return x - dir.x, y - dir.y
}
