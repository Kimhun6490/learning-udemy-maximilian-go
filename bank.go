package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var accountBalance float64 = readBalanceFromFile()

	fmt.Println("Welcome to Go Bank!")

	showMenu(&accountBalance)

	fmt.Println("Bye!")
}

func showMenu(accountBalance *float64) {
	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

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
	writeBalanceToFile(accountBalance)
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
		writeBalanceToFile(accountBalance)
	}
}

func writeBalanceToFile(accountBalance *float64) {
	balanceStr := fmt.Sprintf("%.2f", *accountBalance)

	// 0644 = owner can read/write, group can read, others can read
	os.WriteFile("balance.txt", []byte(balanceStr), 0644)
}

func readBalanceFromFile() float64 {
	data, err := os.ReadFile("balance.txt")
	if err != nil {
		fmt.Println("Error reading balance from file:", err)
		return 0
	}

	// var balance float64
	// fmt.Sscanf(string(data), "%f", &balance)
	// return balance

	balanceStr := string(data)
	balance, err := strconv.ParseFloat(balanceStr, 64)
	if err != nil {
		fmt.Println("Error parsing balance:", err)
		return 0
	}
	return balance
}
