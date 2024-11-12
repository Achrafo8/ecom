package config

import (
	"fmt"
	"os"

	"github.com/lpernett/godotenv"
)

type Config struct {
	PublishHost string
	Port        string
	DBUser      string
	DBPassword  string
	DBAddress   string
	DBName      string
}

var Envs = initConfig()

func initConfig() Config {

	godotenv.Load()
	return Config{

		PublishHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:        getEnv("PORT", "8080"),
		DBUser:      getEnv("DB_USER", "root"),
		DBPassword:  getEnv("DB_PASSWORD", "panda_88"),
		DBAddress:   fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:      getEnv("DB_NAME", "ECOM"),
	}
}

func getEnv(key, fallback string) string {

	if value, ok := os.LookupEnv(key); ok {

		return value
	}
	return fallback
}
