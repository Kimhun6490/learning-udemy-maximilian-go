package main

import "fmt"

func main() {
	age := 32      //regular variable
	agePtr := &age //pointer variable

	fmt.Println("Age:", age)
	fmt.Println("AgePtr:", agePtr)
	fmt.Println("Value at AgePtr:", *agePtr)

	// adultYears := getAdultYears(agePtr)
	// fmt.Println("Adult years:", adultYears)

	editAgeToAdultYears(agePtr)
	fmt.Println("Adult years:", age)
}

func editAgeToAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18 // modify the value at the pointer
}
