package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var (
	ServerAddress string
	EnvMode       string

	// token
	TokenSecretKey string
	TokenDuration  time.Duration

	// database
	DbSource string
	DbName   string

	// minio
	MinioEndpoint      string
	MinioAccessKeyId   string
	MinioAccessKeyPass string
	MinioBucketName    string
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
	EnvMode = os.Getenv("ENV_MODE")

	TokenSecretKey = os.Getenv("TOKEN_SECRET_KEY")
	TokenDuration, _ = time.ParseDuration(os.Getenv("TOKEN_DURATION"))

	DbSource = os.Getenv("DB_SOURCE")
	DbName = os.Getenv("DB_NAME")

	MinioEndpoint = os.Getenv("MINIO_ENDPOINT")
	MinioAccessKeyId = os.Getenv("MINIO_ACCESS_KEY_ID")
	MinioAccessKeyPass = os.Getenv("MINIO_ACCESS_KEY_PASS")
	MinioBucketName = os.Getenv("MINIO_BUCKET_NAME")
}
