package main

import (
	"gofiber/database"
	"gofiber/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func init() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Eror in loading .env file")
	}

	database.ConnectDB()
}

func main() {

	sqlDB, err := database.DBConn.DB()

	if err != nil {
		panic("Error in sql connection")
	}

	defer sqlDB.Close()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(logger.New())

	router.SetupRouters(app)

	app.Listen(":8082")
}
