package main

import "fmt"

func main() {
	// 1)
	hobbies := [3]string{"Reading", "Traveling", "Cooking"}
	fmt.Println("Hobbies:", hobbies)

	// 2)
	fmt.Println("Hobbies:", hobbies[0])
	fmt.Println("Hobbies:", hobbies[1:3])

	// 3)
	mainHobbies := hobbies[:2]
	fmt.Println("Main Hobbies:", mainHobbies)
}
