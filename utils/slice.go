package utils

import "sort"

// inserts a value into a sorted slice and maintains order
func InsertAndSort(slice []int, value int) []int {
	index := sort.Search(len(slice), func(i int) bool {
		return slice[i] >= value
	})

	return append(slice[:index], append([]int{value}, slice[index:]...)...)
}
