package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var (
	ServerAddress      string
	TokenSecretKey     string
	TokenDuration      time.Duration
	DbSource           string
	DbName             string
	MinioEndpoint      string
	MinioAccessKeyId   string
	MinioAccessKeyPass string
	EnvMode            string
)

var (
	DevelopmentMode = "development"
	ProductionMode  = "production"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("couldn't load .env file")
	}

	ServerAddress = os.Getenv("SERVER_ADDRESS")
	TokenSecretKey = os.Getenv("TOKEN_SECRET_KEY")
	TokenDuration, _ = time.ParseDuration(os.Getenv("TOKEN_DURATION"))
	DbSource = os.Getenv("DB_SOURCE")
	DbName = os.Getenv("DB_NAME")
	MinioEndpoint = os.Getenv("MINIO_ENDPOINT")
	MinioAccessKeyId = os.Getenv("MINIO_ACCESS_KEY_ID")
	MinioAccessKeyPass = os.Getenv("MINIO_ACCESS_KEY_PASS")
	EnvMode = os.Getenv("ENV_MODE")
}
