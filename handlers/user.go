package handlers

import (
	"log"
	userdto "test-duaz-solusi/dto/user"
	"test-duaz-solusi/models"
	"test-duaz-solusi/pkg/bcrypt"
	"test-duaz-solusi/pkg/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)


func UserHandlerGetAll(ctx *fiber.Ctx) error {
	var users []models.User
	
	err := mysql.DB.Find(&users).Error
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map {
			"message" : err.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map {
		"message" : "success",
		"data" : users,
	})
}

func UserHandlerCreate(ctx *fiber.Ctx) error {
	user := new(userdto.CreateUserRequest)
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(500).JSON(fiber.Map {
			"message" : err.Error(),
		})
	}

	validation := validator.New()
	errValidate := validation.Struct(user)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map {
			"message" : "failed",
			"error" : errValidate.Error(),
		})
	}

	newUser := models.User{
		Email: user.Email,
		Password: user.Password,
		FullName: user.FullName,
		Gender: user.Gender,
		Phone: user.Phone,
	}

	hashedPassword, err := bcrypt.HashingPassword(user.Password)

	if err != nil {
		log.Printf(hashedPassword)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message" : "internal server error",
		})
	}

	newUser.Password = hashedPassword

	errCreate := mysql.DB.Create(&newUser).Error
		if errCreate != nil {
			return ctx.Status(400).JSON(fiber.Map{
				"message" : "failed to store data",
			})
		}
	
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message" : "success",
		"data" : newUser,
	})

}

func UserHandlerGetByID (ctx *fiber.Ctx ) error {
	id := ctx.Params("id")

	var user models.User

	err := mysql.DB.First(&user, id).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map {
			"message" : "user not found",
		})
	}

	return ctx.Status(200).JSON(fiber.Map {
		"message" : "success",
		"data" : user,
	})
}

func UserHandlerUpdate (ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var user models.User

	err := mysql.DB.First(&user, id).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map {
			"message" : "user not found",
		})
	}

	request := new(userdto.UpdateUserRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(500).JSON(fiber.Map {
			"message" : "internal server error",
		})
	}

	if request.Email != "" {
		user.Email = request.Email
	}

	if request.Password != "" {
		hashedPassword, err := bcrypt.HashingPassword(request.Password)

	if err != nil {
		log.Printf(hashedPassword)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message" : "internal server error",
			})
		}
		user.Password = hashedPassword
	}

	if request.FullName != "" {
		user.FullName = request.FullName
	}

	if request.Gender != "" {
		user.Gender = request.Gender
	}

	if request.Phone != "" {
		user.Phone = request.Phone
	}

	errTemp := mysql.DB.Save(&user).Error

	if errTemp != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message" : "failed to store data",
		})
	}
	
	return ctx.Status(200).JSON(fiber.Map {
		"message" : "success",
		"data" : user,
	})
}

func UserDeleteById (ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var user models.User

	err := mysql.DB.Delete(&user, id).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map {
			"message" : "user not found",
		})
	}

	return ctx.Status(200).JSON(fiber.Map {
		"message" : "success",
	})
}