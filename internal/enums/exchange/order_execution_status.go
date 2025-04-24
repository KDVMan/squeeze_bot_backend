package enums_exchange

type OrderExecutionStatus string

const (
	OrderExecutionStatusNew         OrderExecutionStatus = "NEW"
	OrderExecutionStatusPartialFill OrderExecutionStatus = "PARTIAL_FILL"
	OrderExecutionStatusFill        OrderExecutionStatus = "FILL"
	OrderExecutionStatusCanceled    OrderExecutionStatus = "CANCELED"
	OrderExecutionStatusCalculated  OrderExecutionStatus = "CALCULATED"
	OrderExecutionStatusExpired     OrderExecutionStatus = "EXPIRED"
	OrderExecutionStatusTrade       OrderExecutionStatus = "TRADE"
	OrderExecutionStatusAmendment   OrderExecutionStatus = "AMENDMENT"
)

func (object OrderExecutionStatus) String() string {
	return string(object)
}
