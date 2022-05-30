package config

import (
	"github.com/joho/godotenv"
	"main/internal/model"
	"main/internal/storage"
	"os"
)

var Config model.ConfigType

func Init() {
	Config = model.ConfigType{}

	loadEnv()
	initializeStorage()
}

func loadEnv() {
	errEnvLocal := godotenv.Load(".env.local")
	if errEnvLocal != nil {
		errEnv := godotenv.Load(".env")
		if errEnv != nil {
			panic("File .env and .enn.local is missing.")
		}
	}
}

func initializeStorage() {
	dsn := os.Getenv("DATABASE_DSN")
	storage.Initialize(dsn)
	Config.Storage = storage.GetDB()
}

func GetConfig() model.ConfigType {
	return Config
}
