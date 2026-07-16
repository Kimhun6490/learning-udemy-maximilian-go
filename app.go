package main

import "fmt"

type transformFunc func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	moreNumbers := []int{6, 7, 8, 9, 10}

	doubled := transformNumbers(numbers, double)
	triple := transformNumbers(numbers, triple)

	fmt.Println("Doubled numbers:", doubled)
	fmt.Println("Tripled numbers:", triple)

	transformerFn1 := getTransformFunc(&numbers)
	transformerFn2 := getTransformFunc(&moreNumbers)

	transformed1 := transformNumbers(numbers, transformerFn1)
	transformed2 := transformNumbers(moreNumbers, transformerFn2)

	fmt.Println("Transformed numbers 1:", transformed1)
	fmt.Println("Transformed numbers 2:", transformed2)

	//Anonymous function
	transformed3 := transformNumbers(numbers, func(num int) int {
		return num * 4
	})
	fmt.Println("Transformed numbers 3 (quadrupled):", transformed3)

	//Closure
	multiplier := 5
	transformerFn3 := createTransformFunc(multiplier)
	transformed4 := transformNumbers(numbers, transformerFn3)
	fmt.Println("Transformed numbers 4 (quintupled):", transformed4)
}

func transformNumbers(nums []int, transformFunc transformFunc) []int {
	numbers := []int{}

	for _, num := range nums {
		numbers = append(numbers, transformFunc(num))
	}
	return numbers
}

func getTransformFunc(numbers *[]int) transformFunc {
	if len(*numbers) == 0 {
		return double
	}
	if (*numbers)[0]%2 == 0 {
		return double
	}
	return triple
}

func double(num int) int {
	return num * 2
}

func triple(num int) int {
	return num * 3
}

func createTransformFunc(multiplier int) transformFunc {
	return func(num int) int {
		return num * multiplier
	}
}
