package handlers

import (
	authdto "test-duaz-solusi/dto/auth"
	"test-duaz-solusi/models"
	"test-duaz-solusi/pkg/bcrypt"
	jwtToken "test-duaz-solusi/pkg/jwt"
	"test-duaz-solusi/pkg/mysql"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func LoginHandler(ctx *fiber.Ctx) error {
	loginRequest := new(authdto.LoginRequest)

	if err := ctx.BodyParser(loginRequest); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(loginRequest)

	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map {
			"message" : "login failed",
			"error" : errValidate.Error(),
		})
	}

	var user models.User
	err := mysql.DB.First(&user, "email = ?", loginRequest.Email).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map {
			"message": "login failed",
		})
	}

	isValid := bcrypt.CheckPasswordHash(loginRequest.Password, user.Password)

	if !isValid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map {
			"message" : "login failed",
		})
	}

	claims := jwt.MapClaims{}
	claims["email"] = user.Email
	claims["fullname"] = user.FullName
	claims["phone"] = user.Phone
	claims["gender"] = user.Gender
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
		if errGenerateToken != nil {
			if err != nil {
				return ctx.Status(404).JSON(fiber.Map {
					"message": "login failed",
				})
			}
		}


	return ctx.JSON(fiber.Map {
		"token" : token,
	})
}