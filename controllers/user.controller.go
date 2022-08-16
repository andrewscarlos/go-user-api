package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-web/models"
	"go-web/services"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

type UserControllerInterface interface {
	Find(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type userController struct {
	userService services.UserServiceInterface
}

func NewUserController(userService services.UserServiceInterface) UserControllerInterface {
	return &userController{userService}
}

func (c *userController) Find(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, err := c.userService.Find(id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}
	return ctx.Status(http.StatusOK).JSON(user)
}

func (c *userController) FindAll(ctx *fiber.Ctx) error {
	users, err := c.userService.FindAll()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}
	return ctx.Status(http.StatusOK).JSON(users)
}

func (c *userController) Create(ctx *fiber.Ctx) error {
	var newUser models.User
	err := ctx.BodyParser(&newUser)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(err)
	}
	newUser.Id = bson.NewObjectId()
	newUser.CreateAt = time.Now()
	err = c.userService.Create(&newUser)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(err)
	}
	return ctx.Status(http.StatusOK).JSON(newUser)
}

func (c *userController) Update(ctx *fiber.Ctx) error {
	var userUpdate models.User
	_ = ctx.BodyParser(&userUpdate)
	id := ctx.Params("id")
	user, err := c.userService.Find(id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}
	user.Email = userUpdate.Email
	user.Password = userUpdate.Password
	user.UpdateAt = time.Now()
	err = c.userService.Update(user)
	if err != nil {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(err)
	}
	return ctx.Status(http.StatusOK).JSON(user)
}

func (c *userController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := c.userService.Delete(id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(err)
	}
	return ctx.Status(http.StatusOK).JSON("User deleted")
}
