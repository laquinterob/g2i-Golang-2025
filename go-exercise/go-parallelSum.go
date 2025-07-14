package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 30}
	parts := 3

	total := parallelSum(numbers, parts)
	fmt.Println("Total sum:", total)
}

// parallelSum splits the slice and sums in parallel
func parallelSum(nums []int, workers int) int {
	return 0
}

// Wait for all goroutines to finish and close the channel
