package main

import (
	"fmt"
	"os"
	"github.com/NickG76/blog_aggregator/internal/config"
)
func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Error reading config:", err)
		os.Exit(1)
	}

	fmt.Println("Current DB URL:", cfg.Url)

	if len(os.Args) > 1 {
		newUser := os.Args[1]
		if err := cfg.SetUser(newUser); err != nil {
			fmt.Println("Error setting user:", err)
			os.Exit(1)
		}
		fmt.Println("Updated DB URL to:", cfg.Url)
	} else {
		fmt.Println("No new user provided, exiting.")
	}

} 
