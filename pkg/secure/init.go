package secure

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	secret string
	shield *Shield
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	key := os.Getenv("SECRET_KEY")
	if key == "" {
		log.Fatal("SECRET_KEY not found in .env file")
	}

	secret = key
	shield = NewShield(secret)
}
