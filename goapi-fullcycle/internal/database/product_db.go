package database

import (
	"database/sql"

	"github.com/valdenidelgado/go-projects/goapi-fullcycle/internal/entity"
)

type ProductDB struct {
	db *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{db: db}
}

func (p *ProductDB) GetProducts() ([]*entity.Product, error) {
	rows, err := p.db.Query("SELECT id, name, description, price, category_id FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product

	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.ImageURL, &product.CategoryID); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (p *ProductDB) GetProduct(id string) (*entity.Product, error) {
	var product entity.Product
	err := p.db.QueryRow("SELECT id, name, description, price, image_url, category_id FROM products WHERE id = ?", id).
		Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.ImageURL, &product.CategoryID)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *ProductDB) GetProductByCategory(categoryID string) ([]*entity.Product, error) {
	rows, err := p.db.Query("SELECT id, name, description, price, image_url, category_id FROM products WHERE category_id = ?", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entity.Product
	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.ImageURL, &product.CategoryID); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (p *ProductDB) CreateProduct(product *entity.Product) (string, error) {
	_, err := p.db.Exec("INSERT INTO products (id, name, description, price, image_url, category_id) VALUES (?, ?, ?, ?, ?, ?)", product.ID, product.Name, product.Description, product.Price, product.ImageURL, product.CategoryID)
	if err != nil {
		return "", err
	}
	return product.ID, nil
}
