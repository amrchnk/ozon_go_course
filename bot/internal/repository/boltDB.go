package repository

import (
	"github.com/amrchnk/ozon_go_course/bot/internal/models"
)

type Bucket string

const (
	Product Bucket = "product"
)

type ProductRepository interface {
	CreateProduct(product models.Product) error
	GetProductById(productId int64)(string,error)
	GetProductList()([]models.Product,error)
}
