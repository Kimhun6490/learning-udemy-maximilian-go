package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	// userNames := []string{}
	// userNames = append(userNames, "Alice", "Bob", "Charlie")
	// fmt.Println(userNames)

	userNames := make([]string, 2, 5)
	userNames[0] = "Alice"
	userNames[1] = "Bob"
	fmt.Println(userNames)

	// courseRatings := map[string]float64{}
	// courseRatings["Go"] = 4.5
	// courseRatings["Python"] = 4.7
	// courseRatings["JavaScript"] = 4.3
	// fmt.Println(courseRatings)

	courseRatings := make(floatMap, 3)
	courseRatings["Go"] = 4.5
	courseRatings["Python"] = 4.7
	courseRatings["JavaScript"] = 4.3
	courseRatings.output()

	for index, value := range userNames {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	for _, rating := range courseRatings {
		fmt.Println(rating)
	}
}
