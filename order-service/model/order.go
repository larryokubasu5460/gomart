package model

import (
	"time"

	"github.com/google/uuid"
)

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "PENDING"
	OrderStatusConfirmed OrderStatus = "CONFIRMED"
	OrderStatusFailed    OrderStatus = "FAILED"
)

type Order struct {
	ID          uuid.UUID   `gorm:"type:uuid;primary_key;" json:"id"`
	UserID      uuid.UUID   `gorm:"type:uuid;not null" json:"user_id"`
	TotalAmount float64     `gorm:"not null;" json:"total_amount"`
	Status      OrderStatus `gorm:"type:varchar(20);not null" json:"status"`
	CreatedAt   time.Time   `json:"created_at"`
	OrderItems  []OrderItem `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE;" json:"order_items"`
}

type OrderItem struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	OrderID    uuid.UUID `gorm:"type:uuid;not null" json:"order_id"`
	ProductID  uuid.UUID `gorm:"type:uuid;not null" json:"product_id"`
	Quantity   int       `gorm:"not null" json:"quantity"`
	UnitPrice  float64   `gorm:"not null" json:"unit_price"`
	TotalPrice float64   `gorm:"not null" json:"total_price"`
}
