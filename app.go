package main

import (
	"fmt"
)

func main() {
	resultInt := add(5, 10)
	fmt.Println(resultInt)

	resultFloat := add(5.5, 10.5)
	fmt.Println(resultFloat)

	resultString := add("Hello, ", "World!")
	fmt.Println(resultString)
}

func add[T int | float64 | string](a, b T) T {
	return a + b
}
