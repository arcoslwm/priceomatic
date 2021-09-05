package main

import (
    "log"
    "os"
    "github.com/gofiber/fiber/v2"
    "github.com/joho/godotenv"
    "github.com/KlareTeam/interview-challenges/go/pricematic/products"
    "github.com/KlareTeam/interview-challenges/go/pricematic/database"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func main() {
	app := fiber.New()

    loadDotEnv()
    initDB()

    app.Get("/",index)
    setRoutes(app)

	app.Listen(":8080")
}

func index(c *fiber.Ctx) error {
    return c.SendString("func index Hola Go👋!! ")
}

func loadDotEnv()  {
    err := godotenv.Load("local.env")
    if err != nil {
        log.Fatal("Error loading local.env file: ",err)
        os.Exit(2)
    }
}

func initDB()  {

    envDsn := os.Getenv("DSN_POSTGRES")

    log.Print("Connecting to PostgreSQL DB...")
    dsn := envDsn

    var err error
    database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        log.Fatal("Failed to connect to database. \n", err)
        os.Exit(2)
    }
    log.Println("connected")
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
