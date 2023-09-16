package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var MongodbUri string
var ServerUrl string
var CertFile string
var KeyFile string
var TimeToReopenSession int64

func LoadDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	MongodbUri = os.Getenv("MONGODB_URI")
	ServerUrl = os.Getenv("SERVER_URL")
	CertFile = os.Getenv("CERT_FILE")
	KeyFile = os.Getenv("KEY_FILE")
	TimeToReopenSession = 10 * 60 * 1000 // 10 minutes
}
