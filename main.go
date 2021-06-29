// main.go

package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// load .env file from given path
	// we keep it empty it will load .env from current directory
	godotenv.Load(".env")
	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
		os.Getenv("APP_DB_SCHEMA"))
	a.Run(":8010")
}
