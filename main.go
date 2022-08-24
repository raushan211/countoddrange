package main

import "fmt"

func main() {
	low := 8
	high := 10

	fmt.Println(countOdds(low, high))
}
func countOdds(low int, high int) int {
	count := 0
	for i := low; i <= high; i++ {
		if i%2 != 0 {
			count++
		}
	}
	return count
}
