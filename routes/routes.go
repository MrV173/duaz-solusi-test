package routes

import (
	"test-duaz-solusi/handlers"
	"test-duaz-solusi/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {

	//User Route
	r.Get("/users", handlers.UserHandlerGetAll)
	r.Get("/user/:id", handlers.UserHandlerGetByID)
	r.Post("/user", handlers.UserHandlerCreate)
	r.Put("/user/:id", handlers.UserHandlerUpdate)
	r.Delete("/user/:id", handlers.UserDeleteById)

	//Product Route
	r.Get("/products", handlers.ProductHandlerGetAll)
	r.Get("/product/:id", handlers.ProductHandlerGetByID)
	r.Post("/product",middleware.AuthMiddleware, handlers.ProductHandlerCreate)
	r.Put("/product/:id",middleware.AuthMiddleware, handlers.ProductHandlerUpdate)
	r.Delete("/product/:id",middleware.AuthMiddleware, handlers.ProductDeleteById)

	//Login Route
	r.Post("/login", handlers.LoginHandler)

}
