package main

import (
	"github.com/PrinceNarteh/go-auth/database"
	"github.com/PrinceNarteh/go-auth/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// connecting to database
	database.Connect()

	// instantiate fiber app
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// setting up routes
	routes.Setup(app)

	// app listening on port 3000
	app.Listen("127.0.0.1:3000")
}
