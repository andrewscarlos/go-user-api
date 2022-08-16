package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go-web/controllers"
	"go-web/db"
	"go-web/repository"
	"go-web/routes"
	"go-web/services"
	"log"
)

func main() {
	conn := db.NewConnection()
	defer conn.Close()

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())
	userRepository := repository.NewUserRepository(conn)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)
	userRouter := routes.NewUserRoutes(userController)
	userRouter.RegisterRoutes(app)
	log.Fatal(app.Listen(":3000"))

}
