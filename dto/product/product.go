package productdto

type CreateProductRequest struct {
	Name        string `json:"name" form:"name" validate:"required"`
	Price       string `json:"price" form:"price" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	Stock       string `json:"stock" form:"stock" validate:"required"`
}

type UpdateProductRequest struct {
	Name        string `json:"name" form:"name"`
	Price       string `json:"price" form:"price"`
	Description string `json:"description" form:"description"`
	Stock       string `json:"stock" form:"stock"`
}