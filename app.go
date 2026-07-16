package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	total := sumup(numbers...)
	fmt.Println("The sum is:", total)
}

func sumup(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
