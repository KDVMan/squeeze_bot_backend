package models_order

import (
	enums_exchange "backend/internal/enums/exchange"
)

type OrderModel struct {
	OrderID          string                              `gorm:"unique;not null" json:"orderID"`
	Symbol           string                              `gorm:"not null" json:"symbol"`
	SideType         enums_exchange.SideType             `gorm:"not null" json:"sideType"`
	OrderType        enums_exchange.OrderType            `gorm:"not null" json:"orderType"`
	PositionType     enums_exchange.PositionType         `gorm:"not null" json:"positionType"`
	ExecutionStatus  enums_exchange.OrderExecutionStatus `gorm:"not null" json:"executionStatus"`
	Status           enums_exchange.OrderStatus          `gorm:"not null" json:"status"`
	OriginalPrice    float64                             `json:"originalPrice"`
	AveragePrice     float64                             `json:"averagePrice"`
	OriginalQuantity float64                             `json:"originalQuantity"`
	FilledQuantity   float64                             `json:"filledQuantity"`
	Commission       float64                             `json:"commission"`
	Amount           float64                             `json:"amount"`
}

func (object *OrderModel) UpdateAmount() {
	if object.SideType == enums_exchange.SideTypeBuy {
		object.Amount = object.OriginalPrice * object.OriginalQuantity
	} else if object.SideType == enums_exchange.SideTypeSell {
		object.Amount = object.OriginalQuantity
	}
}
