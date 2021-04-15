package config

import (
	"log"
	"os"
	"strconv"
)

// Config : struct
type Config struct {
	Env        string
	Port       int
	DBHost     string
	DBName     string
	DBUserName string
	DBPassword string
	DBPort     string
}

// GetConfig : ...
func GetConfig() *Config {
	env := os.Getenv("app_env")
	port, _ := strconv.Atoi(os.Getenv("app_port"))
	dbHost := os.Getenv("db_host")
	dbName := os.Getenv("db_name")
	dbUserName := os.Getenv("db_user")
	dbPassword := os.Getenv("db_password")
	dbPort := os.Getenv("db_port")

	log.Printf(`
    env: %s | port: %d
    dbHost: %s
    dbName: %s
    dbUserName: %s
    dbPassword: %s
    dbPort: %s`, env, port, dbHost, dbName, dbUserName, dbPassword, dbPort)

	return &Config{
		Env:        env,
		Port:       port,
		DBHost:     dbHost,
		DBName:     dbName,
		DBUserName: dbUserName,
		DBPassword: dbPassword,
		DBPort:     dbPort,
	}
}
