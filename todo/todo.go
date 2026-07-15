package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Todo struct {
	Text string `json:"text"`
}

func (n Todo) Display() string {
	return fmt.Sprintf("Text: %s", n.Text)
}

func (n Todo) Save() error {
	filename := strings.ReplaceAll(n.Text, " ", "_")
	filename = strings.ToLower(filename) + ".json"

	jsonData, err := json.Marshal(n)
	if err != nil {
		fmt.Println("Error marshalling todo:", err)
		return err
	}

	return os.WriteFile(filename, jsonData, 0644)
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("Text cannot be empty")
	}

	return Todo{
		Text: text,
	}, nil
}
