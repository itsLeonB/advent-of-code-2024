package solvers

import (
	"bufio"
	"log"
	"strings"

	"github.com/itsLeonB/advent-of-code-2024/utils"
)

func Day1Part1(input string) int {
	file := utils.OpenFile(input)
	defer file.Close()

	col1 := make([]int, 0)
	col2 := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		cols := strings.Fields(line)

		val1 := utils.Atoi(cols[0])
		val2 := utils.Atoi(cols[1])

		col1 = utils.InsertAndSort(col1, val1)
		col2 = utils.InsertAndSort(col2, val2)
	}

	len1 := len(col1)
	len2 := len(col2)
	if len1 != len2 {
		log.Fatal("two cols are not the same size")
	}

	var total int
	for i := 0; i < len1; i++ {
		total += utils.Abs(col1[i] - col2[i])
	}

	return total
}

func Day1Part2(input string) int {
	file := utils.OpenFile(input)
	defer file.Close()

	col1 := make([]string, 0)
	col2Map := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		cols := strings.Fields(line)

		col1 = append(col1, cols[0])
		col2Map[cols[1]]++
	}

	var ans int
	for i := 0; i < len(col1); i++ {
		score := col2Map[col1[i]]
		val := utils.Atoi(col1[i])
		ans += score * val
	}

	return ans
}
