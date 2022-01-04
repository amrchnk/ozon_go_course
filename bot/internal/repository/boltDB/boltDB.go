package boltDB

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"github.com/amrchnk/ozon_go_course/bot/internal/models"
	"github.com/amrchnk/ozon_go_course/bot/internal/repository"
	"github.com/boltdb/bolt"
	"strconv"
)

type ProductRepository struct{
	db *bolt.DB
}

func NewProductRepository(db *bolt.DB)*ProductRepository{
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository)CreateProduct(product models.Product) error{
	return r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(repository.Product))

		id, _ := b.NextSequence()
		product.Id = int(id)

		buf, err := json.Marshal(product)
		if err != nil {
			return err
		}

		return b.Put(itob(product.Id), buf)
	})
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}


func (r *ProductRepository)GetProductById(productID int64) (string, error){
	var token string
	err:=r.db.View(func(tx *bolt.Tx) error {
		b:=tx.Bucket([]byte(repository.Product))
		data:=b.Get(strconv.AppendInt(nil, productID, 10))
		token=string(data)
		return nil
	})
	if err!=nil{
		return "",err
	}

	if token==""{
		return "",errors.New("product is not found")
	}

	return token,nil
}

func (r *ProductRepository)GetProductList()([]models.Product,error){
	var products []models.Product
	err:=r.db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(repository.Product))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			product:=models.Product{}
			json.Unmarshal(v,&product)
			products=append(products,product)
		}
		return nil
	})
	return products,err
}