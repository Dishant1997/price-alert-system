package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoDBURI string
	ChainlinkNode string
	ContractAddress string
}

var AppConfig Config

func LoadConfig() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found. Using system environment variables.")
	}

	AppConfig = Config{
		MongoDBURI:     getEnv("MONGODB_URI", "mongodb://localhost:27017"),
		ChainlinkNode:  getEnv("CHAINLINK_NODE_URL", ""),
		ContractAddress: getEnv("CONTRACT_ADDRESS", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}