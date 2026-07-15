package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"exampl.com/note/note"
	"exampl.com/note/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display() string
}

type outputtable interface {
	saver
	displayer
}

// type outputtable interface {
// 	Save() error
// 	Display() string
// }

func main() {
	title, content := getNoteData()
	n, err := note.New(title, content)
	if err != nil {
		printSomething(err)
		return
	}

	text := getTodoData()
	t, err := todo.New(text)
	if err != nil {
		printSomething(err)
		return
	}

	err = outputData(n)
	if err != nil {
		return
	}

	err = outputData(t)
	if err != nil {
		return
	}

	printSomething("Saved successfully.")
}

func printSomething(value any) {
	// intVal, ok := value.(int)

	// if ok {
	// 	fmt.Println("Integer value:", intVal)
	// }

	// strVal, ok := value.(string)

	// if ok {
	// 	fmt.Println("String value:", strVal)
	// }

	switch v := value.(type) {
	case string:
		fmt.Println(v)
	case error:
		fmt.Println("Error:", v.Error())
	default:
		fmt.Println("Unknown type")
	}
}

func outputData(o outputtable) error {
	fmt.Println(o.Display())
	return saveData(o)
}

func saveData(s saver) error {
	err := s.Save()
	if err != nil {
		printSomething(err)
		return err
	}
	return nil
}

func getTodoData() string {
	text := getUserInput("Todo text: ")
	return text
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
