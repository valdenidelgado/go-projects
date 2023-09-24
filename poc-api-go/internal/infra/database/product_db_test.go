package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valdenidelgado/poc-api-go/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestFindAllProducts(t *testing.T) {
	// Arrange
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.NoError(t, err)
		db.Create(product)
	}
	// Act
	productDB := NewProduct(db)
	products, err := productDB.FindAll(1, 10, "asc")

	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")

	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")

	assert.NoError(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 23", products[2].Name)
}

func TestUpdateProduct(t *testing.T) {
	// Arrange
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10.00)
	assert.NoError(t, err)
	db.Create(product)
	// Act
	productDB := NewProduct(db)
	product.Name = "Product 2"
	err = productDB.Update(product)
	// Assert
	assert.NoError(t, err)

	product, err = productDB.FindById(product.ID.String())
	fmt.Printf("aaaaaaaaaaaaaaaaaaaa" + product.Name)

	assert.NoError(t, err)
	assert.Equal(t, "Product 2", product.Name)
}

func TestDeleteProduct(t *testing.T) {
	// Arrange
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10.00)
	assert.NoError(t, err)
	db.Create(product)
	// Act
	productDB := NewProduct(db)
	err = productDB.Delete(product.ID.String())
	// Assert
	assert.NoError(t, err)

	_, err = productDB.FindById(product.ID.String())

	assert.Error(t, err)
}
