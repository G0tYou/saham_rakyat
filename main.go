package main

import (
	"log"
	"os"
	routes "saham_rakyat/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Error loading log file")
	}
	log.SetOutput(file)

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	e := echo.New()
	routes.Routes(e)
	e.Logger.Fatal(e.Start(os.Getenv("APP_PORT")))

}
