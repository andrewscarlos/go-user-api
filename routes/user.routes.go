package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-web/controllers"
)

type userRoutes struct {
	userController controllers.UserControllerInterface
}

func NewUserRoutes(userController controllers.UserControllerInterface) RoutesInterface {
	return &userRoutes{userController}
}

func (r *userRoutes) RegisterRoutes(app *fiber.App) {
	app.Post("/user/create", r.userController.Create)
	app.Get("/user", r.userController.FindAll)
	app.Get("/user/:id", r.userController.Find)
	app.Put("user/:id", r.userController.Update)
	app.Delete("user/:id", r.userController.Delete)
}
