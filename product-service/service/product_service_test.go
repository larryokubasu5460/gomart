package service_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/larryokubasu5460/product-service/model"
	"github.com/larryokubasu5460/product-service/service"
)

// Mock repo

type MockProductRepo struct {
	mock.Mock
}

func (m *MockProductRepo) Create(p *model.Product) error {
	args := m.Called(p)
	return args.Error(0)
}

func (m *MockProductRepo) FindAll() ([]model.Product, error) {
	args := m.Called()
	return args.Get(0).([]model.Product), args.Error(1)
}

func (m *MockProductRepo) FindByID(id uint) (*model.Product, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Product), args.Error(1)
}

func TestProductService_GetAll(t *testing.T) {
	mockRepo := new(MockProductRepo)
	mockProducts := []model.Product{
		{Name: "Product A", Price: 10.0},
		{Name: "Product B", Price: 20.0},
	}

	mockRepo.On("FindAll").Return(mockProducts, nil)

	service := service.NewProductService(mockRepo)
	result, err := service.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result))
	mockRepo.AssertExpectations(t)
}

func TestProductService_GetByID_NotFound(t *testing.T) {
	mockRepo := new(MockProductRepo)
	mockRepo.On("FindByID", uint(999)).Return(&model.Product{}, errors.New("not found"))

	svc := service.NewProductService(mockRepo)
	_, err := svc.GetByID(999)

	assert.Error(t, err)
}
