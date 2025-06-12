package service_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/larryokubasu5460/order-service/client"
	"github.com/larryokubasu5460/order-service/model"
	"github.com/larryokubasu5460/order-service/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockUserClient struct{ mock.Mock }

func (m *mockUserClient) UserExists(userID uuid.UUID) (bool, error) {
	args := m.Called(userID)
	return args.Bool(0), args.Error(1)
}

type mockProductClient struct{ mock.Mock }

func (m *mockProductClient) GetProduct(productID uuid.UUID) (*client.ProductDTO, error) {
	args := m.Called(productID)
	return args.Get(0).(*client.ProductDTO), args.Error(1)
}

type mockOrderRepo struct{ mock.Mock }

func (m *mockOrderRepo) CreateOrder(ctx context.Context, order *model.Order) error {
	args := m.Called(ctx, order)
	return args.Error(0)
}

func (m *mockOrderRepo) GetOrderByID(ctx context.Context, id uuid.UUID) (*model.Order, error) {
	return nil, nil
}

func (m *mockOrderRepo) ListOrdersByUser(ctx context.Context, userID uuid.UUID) ([]model.Order, error) {
	return nil, nil
}

func TestCreateOrder_Success(t *testing.T) {
	userID := uuid.New()
	productID := uuid.New()

	mockUser := new(mockUserClient)
	mockProduct := new(mockProductClient)
	mockRepo := new(mockOrderRepo)

	mockUser.On("UserExists", userID).Return(true, nil)
	mockProduct.On("GetProduct", productID).Return(&client.ProductDTO{
		ID: productID,
		Name: "Phone",
		Price: 100.0,
	}, nil)
	mockRepo.On("CreateOrder", mock.Anything, mock.Anything).Return(nil)
	
	orderService := service.NewOrderService(mockRepo, mockUser, mockProduct)

	req := service.CreateOrderRequest{
		UserID: userID,
		Items:[]service.OrderItemRequest{
			{ProductID:productID, Quantity: 2},
		},
	}

	order, err := orderService.CreateOrder(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, order)
	assert.Equal(t, 200.0, order.TotalAmount)
	mockUser.AssertExpectations(t)
	mockProduct.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}
