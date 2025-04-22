package enums_bot

type DealStatus string

const (
	DealStatusNull  DealStatus = "null"
	DealStatusOpen  DealStatus = "open"
	DealStatusClose DealStatus = "close"
)

func (object DealStatus) String() string {
	return string(object)
}
