package repository

import (
	"github.com/amrchnk/ozon_go_course/bot/internal/models"
)

type Bucket string

const (
	Products Bucket = "products"
)

type ProductRepository interface {
	CreateProduct(product *models.Product) error
	GetProductById(productId int)(models.Product,error)
	GetProductList()([]models.Product,error)
}
