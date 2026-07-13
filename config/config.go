package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	OpenAiKey string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env not found ")
	}
	return &Config{
		Port:      os.Getenv("PORT"),
		OpenAiKey: os.Getenv("OPENAI_API_KEY"),
	}
}
