package services_interface_order

import models_order "backend/internal/models/order"

type OrderService interface {
	Update(*models_order.OrderModel)
	RunOrderChannel()
	GetOrderChannel() chan *models_order.OrderModel
	GetAmount() float64
}
