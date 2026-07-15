package main

import "fmt"

func main() {
	// website := []string{"google.com", "facebook.com", "youtube.com"}
	website := map[string]string{
		"Google":   "google.com",
		"Facebook": "facebook.com",
		"YouTube":  "youtube.com",
	}

	fmt.Println(website)
	fmt.Println(website["Google"])

	website["Twitter"] = "twitter.com"
	fmt.Println(website)

	delete(website, "Facebook")
	fmt.Println(website)
}
