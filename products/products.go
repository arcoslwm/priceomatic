package products

import (
    "log"
    "time"
    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
    "github.com/KlareTeam/interview-challenges/go/pricematic/database"
)

type Product struct {
    gorm.Model
    // ``
    // Id uint `gorm:"primaryKey" json:"id"`
    ID        uint       `gorm:"primary_key"; json:"id"`
    Name string `gorm:"size:120"; json:"name"`
    Price uint `json:"price"`
    CreatedAt time.Time  `json:"-"`
    UpdatedAt time.Time  `json:"-"`
    DeletedAt *time.Time `json:"-"`
    // CreatedAt time.Time `gorm:"column:fecha_de_insercion" json:"createdAt"`
}

func GetProducts (c *fiber.Ctx) error {
    db := database.DBConn
    var products []Product
    db.Find(&products)
    return c.JSON(products)
}

func GetProduct (c *fiber.Ctx) error {
    id := c.Params("id")
    db := database.DBConn
    var product Product
    db.Find(&product,id)
    if product.ID == 0 {
        return c.Status(404).SendString("product not found")
    }
    return c.JSON(product)
}

/**
 * // TODO: -validaciones data nombre required por ejemplo.
 */
func AddProduct (c *fiber.Ctx) error {
    db := database.DBConn
    product := new(Product)
    if err := c.BodyParser(product); err != nil {
        log.Print("BodyParser error: ",err)
        return c.Status(400).SendString("error on request")
    }

    db.Create(&product)
    return c.JSON(product)
}

func UpdateProduct (c *fiber.Ctx) error {
    return c.SendString("updated product")
}

func DeleteProduct (c *fiber.Ctx) error {
    id := c.Params("id")
    db := database.DBConn
    var product Product
    db.First(&product,id)
    if product.ID == 0 {
        return c.Status(404).SendString("product not found")
    }
    db.Delete(&product)
    return c.SendString("deleted product id: ")
}
