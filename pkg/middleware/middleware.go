package middleware

import (
	jwtToken "test-duaz-solusi/pkg/jwt"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(ctx *fiber.Ctx) error {
	token := ctx.Get("x-token")
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Please Login ",
		})
	}

	_, err := jwtToken.VerifyToken(token)

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Please Login",
		})
	}

	return ctx.Next()
}