package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/larryokubasu5460/order-service/handler"
	"github.com/larryokubasu5460/order-service/model"
	"github.com/larryokubasu5460/order-service/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockOrderService struct{ mock.Mock }

func (m *mockOrderService) CreateOrder(ctx context.Context, req service.CreateOrderRequest) (*model.Order, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*model.Order), args.Error(1)
}

func (m *mockOrderService) GetOrder(ctx context.Context, id uuid.UUID) (*model.Order, error) {
	return nil, nil
}

func (m *mockOrderService) ListUserOrders(ctx context.Context, userID uuid.UUID) ([]model.Order, error) {
	return nil, nil
}

func TestCreateOrderHandler(t *testing.T) {
	userID := uuid.New()
	productID := uuid.New()

	mockService := new(mockOrderService)
	handler := handler.NewOrderHandler(mockService)

	order := &model.Order{
		ID:          uuid.New(),
		UserID:      userID,
		TotalAmount: 100,
	}

	mockService.On("CreateOrder", mock.Anything, mock.Anything).Return(order, nil)

	// prepare request
	body := service.CreateOrderRequest{
		UserID: userID,
		Items: []service.OrderItemRequest{
			{ProductID: productID, Quantity: 1},
		},
	}

	jsonBody, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "/orders", bytes.NewReader(jsonBody))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.CreateOrder(c)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockService.AssertExpectations(t)
}
