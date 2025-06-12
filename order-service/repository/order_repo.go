package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/larryokubasu5460/order-service/model"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *model.Order) error
	GetOrderByID(ctx context.Context, id uuid.UUID) (*model.Order, error)
	ListOrdersByUser(ctx context.Context, userID uuid.UUID) ([]model.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) CreateOrder(ctx context.Context, order *model.Order) error {
	return r.db.WithContext(ctx).Create(order).Error
}

func (r *orderRepository) GetOrderByID(ctx context.Context, id uuid.UUID) (*model.Order, error) {
	var order model.Order
	err := r.db.WithContext(ctx).Preload("OrderItems").First(&order, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *orderRepository) ListOrdersByUser(ctx context.Context, userID uuid.UUID) ([]model.Order, error) {
	var orders []model.Order
	err := r.db.WithContext(ctx).
		Preload("OrderItems").
		Where("user_id = ?", userID).
		Order("created_at desc").
		Find(&orders).Error

	return orders, err
}
