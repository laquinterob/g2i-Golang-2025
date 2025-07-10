package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	parts := 3

	total := parallelSum(numbers, parts)
	fmt.Println("Total sum:", total)
}

// parallelSum splits the slice and sums in parallel
func parallelSum(nums []int, workers int) int {
	length := len(nums)
	chunkSize := (length + workers - 1) / workers

	var wg sync.WaitGroup
	resultChan := make(chan int, workers)

	for i := 0; i < workers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > length {
			end = length
		}

		chunk := nums[start:end]
		wg.Add(1)
		go func(data []int) {
			defer wg.Done()
			sum := 0
			for _, v := range data {
				sum += v
			}
			resultChan <- sum
		}(chunk)
	}

	// Wait for all goroutines to finish and close the channel
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	total := 0
	for partial := range resultChan {
		total += partial
	}

	return total
}
