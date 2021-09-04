package main

import "github.com/gofiber/fiber/v2"



func main() {
	app := fiber.New()

    app.Get("/",index)

	app.Listen(":8080")
}

func index(c *fiber.Ctx) error {
    return c.SendString("func index real Hola Go 👋!!!! ")
}
