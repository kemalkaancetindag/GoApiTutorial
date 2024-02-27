package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln("Error while loading .env")
	}

	mongoURI := os.Getenv("MONGO_URI")

	return mongoURI
}
