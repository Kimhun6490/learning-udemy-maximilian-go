package fileops

import (
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(fileName string, value float64) {
	valueStr := fmt.Sprintf("%.2f", value)

	os.WriteFile(fileName, []byte(valueStr), 0644)
}

func GetFloatFromFile(fileName string) float64 {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading balance from file:", err)
		return 0
	}

	valueStr := string(data)
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		fmt.Println("Error parsing value:", err)
		return 0
	}
	return value
}
