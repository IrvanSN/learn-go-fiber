package main

import (
	"fmt"
	"github.com/IrvanSN/learn-go-fiber/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.SendString("OK")
	})

	app.Post("/api/products", handlers.CreateProduct)
	app.Get("/api/products", handlers.GetAllProduct)

	err := app.Listen(":8080")

	if err != nil {
		return
	}

	fmt.Println("cok")
}
