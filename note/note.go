package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (n Note) Display() string {
	return fmt.Sprintf("Title: %s\nContent: %s\nCreated At: %s", n.Title, n.Content, n.CreatedAt.Format(time.RFC1123))
}

func (n Note) Save() error {
	filename := strings.ReplaceAll(n.Title, " ", "_")
	filename = strings.ToLower(filename) + ".json"

	jsonData, err := json.Marshal(n)
	if err != nil {
		fmt.Println("Error marshalling note:", err)
		return err
	}

	return os.WriteFile(filename, jsonData, 0644)
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Title and content cannot be empty")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
