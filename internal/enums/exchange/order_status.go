package enums_exchange

type OrderStatus string

const (
	OrderStatusNew             OrderStatus = "NEW"
	OrderStatusPartiallyFilled OrderStatus = "PARTIALLY_FILLED"
	OrderStatusFilled          OrderStatus = "FILLED"
	OrderStatusCanceled        OrderStatus = "CANCELED"
	OrderStatusRejected        OrderStatus = "REJECTED"
	OrderStatusExpired         OrderStatus = "EXPIRED"
	OrderStatusNewInsurance    OrderStatus = "NEW_INSURANCE"
	OrderStatusNewADL          OrderStatus = "NEW_ADL"
)

func (object OrderStatus) String() string {
	return string(object)
}
