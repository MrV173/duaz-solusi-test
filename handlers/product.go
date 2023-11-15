package handlers

import (
	productdto "test-duaz-solusi/dto/product"
	"test-duaz-solusi/models"
	"test-duaz-solusi/pkg/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ProductHandlerGetAll(ctx *fiber.Ctx) error {
	var products []models.Product

	err := mysql.DB.Find(&products).Error
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map {
			"message" : err.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map {
		"message" : "success",
		"data" : products,
	})
}

func ProductHandlerCreate(ctx *fiber.Ctx) error {
	product := new(productdto.CreateProductRequest)
	if err := ctx.BodyParser(product); err != nil {
		return ctx.Status(500).JSON(fiber.Map {
			"message" : err.Error(),
		})
	}

	validation := validator.New()
	errValidate := validation.Struct(product)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map {
			"message" : "failed",
			"error" : errValidate.Error(),
		})
	}

	newProduct := models.Product{
		Name: product.Name,
		Price: product.Price,
		Description: product.Description,
		Stock: product.Stock,
	}

	err := mysql.DB.Create(&newProduct).Error
		if err != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"message" : "failed to store data",
			})
		}
	
	return ctx.JSON(fiber.Map{
		"message" : "success",
		"data" : newProduct,
	})
}

func ProductHandlerGetByID (ctx *fiber.Ctx ) error {
	id := ctx.Params("id")

	var product models.Product

	err := mysql.DB.First(&product, id).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map {
			"message" : "product not found",
		})
	}

	return ctx.Status(200).JSON(fiber.Map {
		"message" : "success",
		"data" : product,
	})
}

func ProductHandlerUpdate (ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var product models.Product

	err := mysql.DB.First(&product, id).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map {
			"message" : "product not found",
		})
	}

	request := new(productdto.UpdateProductRequest)
	if err := ctx.BodyParser(request); err != nil {
		return ctx.Status(500).JSON(fiber.Map {
			"message" : "internal server error",
		})
	}



	if request.Name != "" {
		product.Name = request.Name
	}

	if request.Price != "" {
		product.Price = request.Price
	}

	if request.Description != "" {
		product.Description = request.Description
	}

	if request.Stock != "" {
		product.Stock = request.Stock
	}

	errTemp := mysql.DB.Save(&product).Error

	if errTemp != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message" : "failed to store data",
		})
	}
	
	return ctx.Status(200).JSON(fiber.Map {
		"message" : "success",
		"data" : product,
	})
}

func ProductDeleteById (ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var product models.Product

	err := mysql.DB.Delete(&product, id).Error

	if err != nil {
		return ctx.Status(404).JSON(fiber.Map {
			"message" : "product not found",
		})
	}

	return ctx.Status(200).JSON(fiber.Map {
		"message" : "success",
	})
}
