package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/itsLeonB/advent-of-code-2024/solvers"
	"github.com/itsLeonB/advent-of-code-2024/utils"
)

func main() {
	solver := map[string]func(input string) int{
		"1-1": solvers.Day1Part1,
		"1-2": solvers.Day1Part2,
		"2-1": solvers.Day2Part1,
		"2-2": solvers.Day2Part2,
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter the day to solve: ")
		dayInput := utils.ReadNewLine(reader)

		filePath := fmt.Sprintf("inputs/day%s.txt", dayInput)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			fmt.Printf("File %s does not exist.\n", filePath)
			continue
		}

		fmt.Print("Enter the part to solve (e.g., 1 or 2): ")
		partInput := utils.ReadNewLine(reader)

		solverKey := fmt.Sprintf("%s-%s", dayInput, partInput)
		solverFunc, exists := solver[solverKey]
		if !exists {
			fmt.Printf("No solver found for day %s, part %s.\n", dayInput, partInput)
			continue
		}

		ans := solverFunc(filePath)
		fmt.Printf("Answer: %d\n", ans)

		fmt.Print("Solve another? (yes/no): ")
		again := utils.ReadNewLine(reader)

		if again != "yes" {
			fmt.Println("Exiting...")
			break
		}
	}
}
