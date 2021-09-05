package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/KlareTeam/interview-challenges/go/pricematic/products"
    // "./products"
)



func main() {
	app := fiber.New()

    app.Get("/",index)
    setRoutes(app)

	app.Listen(":8080")
}

func index(c *fiber.Ctx) error {
    return c.SendString("func index Hola Go👋!! ")
}

func setRoutes (app *fiber.App) {

    // * GET `/api/v1/products` <-- Lista todos los productos
    // * GET `/api/v1/products/:id` <-- Obtiene el producto :id
    // * POST `/api/vi/products` <-- Inserta un producto en la base de datos

    // * [BONUS] `PATCH /api/vi/products/:id` <-- Actualiza uno o mas campos del producto :id
    // * [BONUS] `DELETE /api/vi/products/:id` <-- Elimina el producto :id

    app.Get("/api/v1/products", products.GetProducts)
    app.Get("/api/v1/products/:id", products.GetProduct)
    app.Post("/api/v1/products", products.AddProduct)
}
