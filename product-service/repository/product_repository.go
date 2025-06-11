package repository

import (
	"github.com/larryokubasu5460/product-service/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(prodct *model.Product) error
	FindAll() ([]model.Product, error)
	FindByID(id uint) (*model.Product, error)
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepo{db}
}

func (r *productRepo) Create(product *model.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepo) FindAll() ([]model.Product, error) {
	var products []model.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *productRepo) FindByID(id uint) (*model.Product, error) {
	var product model.Product
	err := r.db.First(&product,id).Error
	return &product, err
}