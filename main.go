package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alter334/go_bot_template/base"
	"github.com/joho/godotenv"
)

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	bottoken, err := getEnv("TRAQ_BOT_TOKEN")
	if err != nil {
		return
	}
	base.Setup(bottoken)

}

func getEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("environment variable %s not set", key)
	}
	return value, nil
}
