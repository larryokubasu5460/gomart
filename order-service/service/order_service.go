package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/larryokubasu5460/order-service/client"
	"github.com/larryokubasu5460/order-service/model"
	"github.com/larryokubasu5460/order-service/repository"
)

type OrderService interface {
	CreateOrder(ctx context.Context, req CreateOrderRequest) (*model.Order, error)
	GetOrder(ctx context.Context, id uuid.UUID) (*model.Order, error)
	ListUserOrders(ctx context.Context, userID uuid.UUID) ([]model.Order, error)
}

type orderService struct {
	repo          repository.OrderRepository
	userClient    client.UserClient
	productClient client.ProductClient
}

func NewOrderService(repo repository.OrderRepository, userClient client.UserClient, productClient client.ProductClient) OrderService {
	return &orderService{
		repo:          repo,
		userClient:    userClient,
		productClient: productClient,
	}
}

func (s *orderService) CreateOrder(ctx context.Context, req CreateOrderRequest) (*model.Order, error) {
	// validate user
	userExists, err := s.userClient.UserExists(req.UserID)
	if err != nil {
		return nil, fmt.Errorf("error checking user: %w", err)
	}
	if !userExists {
		return nil, fmt.Errorf("user does not exist")
	}

	// prepare order
	order := &model.Order{
		ID:          uuid.New(),
		UserID:      req.UserID,
		Status:      model.OrderStatusPending,
		CreatedAt:   time.Now(),
		OrderItems:  make([]model.OrderItem, 0),
		TotalAmount: 0,
	}

	// validate each product, calculate price
	for _, item := range req.Items {
		product, err := s.productClient.GetProduct(item.ProductID)
		if err != nil {
			return nil, fmt.Errorf("product %s error: %w", item.ProductID, err)
		}

		totalPrice := float64(item.Quantity) * product.Price
		orderItem := model.OrderItem{
			ID:         uuid.New(),
			OrderID:    order.ID,
			ProductID:  item.ProductID,
			Quantity:   item.Quantity,
			UnitPrice:  product.Price,
			TotalPrice: totalPrice,
		}

		order.OrderItems = append(order.OrderItems, orderItem)
		order.TotalAmount += totalPrice
	}

	// save to DB
	if err := s.repo.CreateOrder(ctx, order); err != nil {
		return nil, fmt.Errorf("failed to create order: %w", err)
	}

	return order, nil
}

func (s *orderService) GetOrder(ctx context.Context, id uuid.UUID) (*model.Order, error) {
	return s.repo.GetOrderByID(ctx, id)
}

func (s *orderService) ListUserOrders(ctx context.Context, userID uuid.UUID) ([]model.Order, error) {
	return s.repo.ListOrdersByUser(ctx, userID)
}
