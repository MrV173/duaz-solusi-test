package models

type Product struct {
	ID          int    `json:"id" gorm:"primaryKey:autoIncrement"`
	Name        string `json:"name" gorm:"type: varchar(255)"`
	Price       string `json:"price"`
	Description string `json:"description" gorm:"type: varchar(255)"`
	Stock       string `json:"stock"`
}

type ProductResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       string `json:"price"`
	Description string `json:"description"`
	Stock       string `json:"stock"`
}

func (ProductResponse) TableName() string {
	return "products"
}