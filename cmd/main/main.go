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
		"3-1": solvers.Day3Part1,
		"3-2": solvers.Day3Part2,
		"4-1": solvers.Day4Part1,
		"4-2": solvers.Day4Part2,
		"5-1": solvers.Day5Part1,
	}

	answers := map[string]int{
		"1-1": 2742123,
		"1-2": 21328497,
		"2-1": 585,
		"2-2": 626,
		"3-1": 173529487,
		"3-2": 99532691,
		"4-1": 2378,
		"4-2": 1796,
		"5-1": 5651,
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

		problemKey := fmt.Sprintf("%s-%s", dayInput, partInput)
		solverFunc, exists := solver[problemKey]
		if !exists {
			fmt.Printf("No solver found for day %s, part %s.\n", dayInput, partInput)
			continue
		}

		ans := solverFunc(filePath)
		fmt.Printf("Answer: %d\n", ans)

		if trueAns, exists := answers[problemKey]; exists {
			if trueAns == ans {
				fmt.Println("Your solution is correct!")
			} else {
				fmt.Println("Your solution is incorrect. Please look at it again.")
			}
		} else {
			fmt.Println("True answer is not found.")
		}

		fmt.Print("Solve another? (y/n): ")
		again := utils.ReadNewLine(reader)

		if again != "y" {
			fmt.Println("Exiting...")
			break
		}
	}
}
