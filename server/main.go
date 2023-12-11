package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)



func main() {

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/dashboard", func(c * fiber.Ctx) error {
		return c.SendString("dashboard")
	})

	app.Get("/login", func(c * fiber.Ctx) error {
		return c.SendString("login")
	})

	app.Get("/signup", func(c * fiber.Ctx) error {
		return c.SendString("signup")
	})

	app.Post("/login", func(c * fiber.Ctx) error {
		return c.SendString("login post")
	})

	app.Post("/signup", func(c * fiber.Ctx) error {
		return c.SendString("signup post")
	})


	app.Get("/player-list", func(c * fiber.Ctx) error {
		return c.SendString("player list")
	})

	app.Get("/head-to-head", func(c * fiber.Ctx) error {
		return c.SendString("head to head")
	})

	log.Fatal(app.Listen(":4000"))
	
}