package handler_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/larryokubasu5460/product-service/handler"
	"github.com/larryokubasu5460/product-service/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock ProductService

type MockService struct {
	mock.Mock
}

func (m *MockService) Create(p *model.Product) error {
	args := m.Called(p)
	return args.Error(0)
}

func (m *MockService) GetAll() ([]model.Product, error) {
	args := m.Called()
	return args.Get(0).([]model.Product), args.Error(1)
}

func (m *MockService) GetByID(id uint) (*model.Product, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Product), args.Error(1)
}

func TestGetAllProducts(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := new(MockService)
	mockSvc.On("GetAll").Return([]model.Product{
		{Name: "Sample"}, {Name: "Another"},
	}, nil)

	h := handler.ProductHandler{Service: mockSvc}
	r := gin.Default()
	h.RegisterRoutes(r)

	req, _ := http.NewRequest(http.MethodGet, "/products", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockSvc.AssertExpectations(t)
}

func TestCreateProduct_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockSvc := new(MockService)
	productJSON := `{"name":"New Product", "description":"Test", "price":100}`

	mockSvc.On("Create", mock.AnythingOfType("model.Product")).Return(nil)

	h := handler.ProductHandler{Service: mockSvc}
	r := gin.Default()
	h.RegisterRoutes(r)

	req, _ := http.NewRequest(http.MethodPost, "/products",strings.NewReader(productJSON))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp,req)

	assert.Equal(t,http.StatusCreated, resp.Code)
	mockSvc.AssertExpectations(t)
}