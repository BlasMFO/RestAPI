package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/v1/movies", Home)
	app.Get("/v1/movies/:user_movie", SelectMovie)
}

func main() {

	fmt.Println("Running gofiber-app")

	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}
