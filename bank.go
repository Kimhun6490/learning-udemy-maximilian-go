package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const balanceFileName = "balance.txt"

func main() {
	var accountBalance float64 = fileops.GetFloatFromFile(balanceFileName)

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())

	showMenu(&accountBalance)

	fmt.Println("Bye!")
}

func showMenu(accountBalance *float64) {
	for {
		presentOptions()

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			checkBalance(accountBalance)
		case 2:
			depositMoney(accountBalance)
		case 3:
			withdrawMoney(accountBalance)
		case 4:
			fmt.Println("Thank you for using Go Bank!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func checkBalance(accountBalance *float64) {
	fmt.Printf("Your current balance is $%.2f.\n", *accountBalance)
}

func depositMoney(accountBalance *float64) {
	var amount float64
	fmt.Println("Enter the amount to deposit:")
	fmt.Scanln(&amount)
	*accountBalance += amount
	fmt.Printf("$%.2f has been deposited to your account.\n", amount)
	fileops.WriteFloatToFile(balanceFileName, *accountBalance)
}

func withdrawMoney(accountBalance *float64) {
	var amount float64
	fmt.Println("Enter the amount to withdraw:")
	fmt.Scanln(&amount)

	if amount <= 0 {
		fmt.Println("Invalid amount. Please enter a positive value.")
		return
	} else if amount > *accountBalance {
		fmt.Println("Insufficient funds.")
	} else {
		*accountBalance -= amount
		fmt.Printf("$%.2f has been withdrawn from your account.\n", amount)
		fileops.WriteFloatToFile(balanceFileName, *accountBalance)
	}
}
