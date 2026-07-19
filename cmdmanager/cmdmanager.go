package cmdmanager

import "fmt"

type CommandManager struct{}

func NewCommandManager() *CommandManager {
	return &CommandManager{}
}

func (cm *CommandManager) ReadLinesFromFile() ([]string, error) {
	fmt.Println("Please enter your prices. Confirm every price with ENTER")

	prices := []string{}

	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)
		if price == "0" {
			fmt.Println("End")
			break
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (cm *CommandManager) WriteLinesToFile(data interface{}) error {
	fmt.Println("Tax-included prices:")
	fmt.Println(data)
	return nil
}
