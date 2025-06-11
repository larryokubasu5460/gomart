package service

import (
	"github.com/larryokubasu5460/product-service/model"
	"github.com/larryokubasu5460/product-service/repository"
)

type ProductService interface {
	Create(product *model.Product) error
	GetAll() ([]model.Product, error)
	GetByID(id uint) (*model.Product, error)
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo}
}

func (s *productService) Create(product *model.Product) error {
	return s.repo.Create(product)
}

func (s *productService) GetAll() ([]model.Product, error) {
	return s.repo.FindAll()
}

func (s *productService) GetByID(id uint) (*model.Product, error) {
	return s.repo.FindByID(id)
}