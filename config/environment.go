package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DBUsername string
	DBPassword string
	DBPort     string
	DBHost     string
	DBName     string
}

var JWTSECRET = ""

func AssignEnv(c AppConfig) (AppConfig, bool) {
	var missing = false
	if val, found := os.LookupEnv("DBUsername"); found {
		c.DBUsername = val
	} else {
		missing = true
	}
	if val, found := os.LookupEnv("DBPassword"); found {
		c.DBPassword = val
	} else {
		missing = true
	}
	if val, found := os.LookupEnv("DBPort"); found {
		c.DBPort = val
	} else {
		missing = true
	}
	if val, found := os.LookupEnv("DBHost"); found {
		c.DBHost = val
	} else {
		missing = true
	}
	if val, found := os.LookupEnv("DBName"); found {
		c.DBName = val
	} else {
		missing = true
	}
	if val, found := os.LookupEnv("JWT_SECRET"); found {
		JWTSECRET = val
	} else {
		missing = true
	}
	return c, missing
}

func InitConfig() AppConfig {
	var result AppConfig
	var missing = false
	result, missing = AssignEnv(result)
	if missing {
		godotenv.Load(".env")
		result, _ = AssignEnv(result)
	}

	return result
}
