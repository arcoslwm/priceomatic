package products

import (
    "github.com/gofiber/fiber/v2"
)

func GetProducts (c *fiber.Ctx) error {
    return c.SendString("list all products")
}

func GetProduct (c *fiber.Ctx) error {
    return c.SendString("return one product")
}

func AddProduct (c *fiber.Ctx) error {
    return c.SendString("add product")
}

func UpdateProduct (c *fiber.Ctx) error {
    return c.SendString("updated product")
}

func DeleteProduct (c *fiber.Ctx) error {
    return c.SendString("updated product")
}
