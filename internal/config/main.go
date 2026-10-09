package config

import (
	"os"

	"github.com/joho/godotenv"
)

//capital c is for accessing in publically in another files 
type Config struct {
	Port string
	Env string 
}


func MustLoad() Config{
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT is required")// crashing app if port number is not there
	}

	env := os.Getenv("ENV")
	if env == "" {
		panic("ENV is required")
	}


	return Config{
		Port: port,
		Env:  env,
	}

}