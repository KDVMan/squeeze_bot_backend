package enums_bot

type DealStatus string

const (
	DealStatusNull          DealStatus = "null"
	DealStatusSendOpenLimit DealStatus = "send_open_limit"
	DealStatusOpenLimit     DealStatus = "open_limit"
	DealStatusOpen          DealStatus = "open"
	DealStatusSendClose     DealStatus = "send_close"
	DealStatusClose         DealStatus = "close"
)

func (object DealStatus) String() string {
	return string(object)
}
