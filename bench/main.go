package main

import (
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()
    app.Get("/", func(c *fiber.Ctx) error {
        return nil
    })
    app.Listen("0.0.0.0:8080")
}
