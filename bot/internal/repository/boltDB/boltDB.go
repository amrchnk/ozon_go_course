package boltDB

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/amrchnk/ozon_go_course/bot/internal/models"
	"github.com/amrchnk/ozon_go_course/bot/internal/repository"
	"github.com/boltdb/bolt"
)

type ProductRepository struct{
	db *bolt.DB
}

func NewProductRepository(db *bolt.DB)*ProductRepository{
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository)CreateProduct(product *models.Product) error{
	err:=r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(repository.Products))

		id, _ := b.NextSequence()
		product.Id = int(id)-1
		buf, err := json.Marshal(product)
		if err != nil {
			return err
		}
		b.Put(intToByte(product.Id), buf)
		fmt.Println(product)
		return nil
	})
	return err
}

// intToByte returns an 8-byte big endian representation of v.
func intToByte(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}


func (r *ProductRepository)GetProductById(productId int) (models.Product, error){
	product:=models.Product{}
	err:=r.db.View(func(tx *bolt.Tx) error {
		b:=tx.Bucket([]byte(repository.Products))
		data:=b.Get(intToByte(productId))
		json.Unmarshal(data,&product)
		return nil
	})
	if err!=nil{
		return product,err
	}

	return product,nil
}

func (r *ProductRepository)GetProductList()([]models.Product,error){
	var products []models.Product
	err:=r.db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(repository.Products))

		b.ForEach(func(key, value []byte) error {
			product:=models.Product{}
			json.Unmarshal(value,&product)
			products=append(products,product)
			return nil
		})

		return nil
	})
	return products,err
}

func (r *ProductRepository)DeleteProductById(productId int)error{
	return r.db.Update(func(tx *bolt.Tx) error {
		b:=tx.Bucket([]byte(repository.Products))
		if err:=b.Get(intToByte(productId));err==nil{
			return errors.New(fmt.Sprint("Товара нет в списке"))
		}

		return b.Delete(intToByte(productId))
	})
}