package service

import "github.com/google/uuid"

type CreateOrderRequest struct {
	UserID uuid.UUID          `json:"user_id"`
	Items  []OrderItemRequest `json:"items"`
}

type OrderItemRequest struct {
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
}
