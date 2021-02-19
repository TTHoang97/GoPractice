// main.go

package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)


func main() {

	//Reading environment variables from a .env file and loading them into the app
	err := godotenv.Load("database.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	a.Run(":8010")
}