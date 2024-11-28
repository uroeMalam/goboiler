package main

import (
	"fmt"
	"log"
	"markopillow/init/internal/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	viperConfig := config.NewViper()
	app := config.NewFiber(viperConfig)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello astronout!")
	})

	webPort := 3000
	err := app.Listen(fmt.Sprintf(":%d", webPort))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
