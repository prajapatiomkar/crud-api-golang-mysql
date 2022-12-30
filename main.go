package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/prajapatiomkar/crud-api-golang-mysql/database"
	"github.com/prajapatiomkar/crud-api-golang-mysql/router"
)

func main() {
	// Connect with database
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	router.SetupRoutes(app)
	app.Listen(":8080") 
}
