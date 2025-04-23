package enums_bot

type DealStatus string

const (
	DealStatusNull            DealStatus = "null"
	DealStatusSendOpenLimit   DealStatus = "send_open_limit"
	DealStatusSendOpenLimitWs DealStatus = "send_open_limit_ws"
	DealStatusOpenLimit       DealStatus = "open_limit"
	DealStatusOpen            DealStatus = "open"
	DealStatusClose           DealStatus = "close"
)

func (object DealStatus) String() string {
	return string(object)
}
