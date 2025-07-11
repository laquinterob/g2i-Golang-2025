package strings

import "fmt"

func MostFrequent(words []string) (string, int) {
	// TODO: implement
	return "", 0
}

func main() {
	words := []string{"go", "rust", "go", "python", "go", "rust"}
	word, count := MostFrequent(words)
	fmt.Printf("Most frequent: %s (%d times)\n", word, count)
}
