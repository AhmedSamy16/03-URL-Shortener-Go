package main

import (
	"github.com/AhmedSamy16/03-url-shortener-Go/application"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	app := application.New()

	app.Start()
}
