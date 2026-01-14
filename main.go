package main

import (
	"fmt"
	"os"
)

func main() {
	token := os.Getenv("DISCORD_TOKEN")
	if token == "" {
		fmt.Println("No DISCORD_TOKEN found. Please set DISCORD_TOKEN environment variable.")
		return
	}

	bot, err := NewUwuBot(token)
	if err != nil {
		fmt.Printf("Error creating Discord session: %v\n", err)
		return
	}

	if err := bot.Run(); err != nil {
		fmt.Printf("Error running bot: %v\n", err)
	}
}
