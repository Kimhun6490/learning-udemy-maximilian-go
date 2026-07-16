package main

import "fmt"

func main() {
	fact := factorial(5)
	fmt.Printf("Factorial of 5 is: %d\n", fact)
}

// Recursion
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// func factorial(n int) int {
// 	result := 1

// 	for i := 1; i <= n; i++ {
// 		result *= i
// 	}

// 	return result
// }
