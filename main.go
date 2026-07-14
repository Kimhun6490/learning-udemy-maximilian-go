package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"exampl.com/note/note"
)

func main() {
	title, content := getNoteData()

	n, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	n.Display()

	err = n.Save()
	if err != nil {
		fmt.Println("Error saving note:", err)
	}

	fmt.Println("Note saved successfully.")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%s", prompt)

	value, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	value = strings.TrimSuffix(value, "\n")
	value = strings.TrimSuffix(value, "\r")

	return value
}
