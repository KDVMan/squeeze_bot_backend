package enums_exchange

type SideType string

const (
	SideTypeBuy  SideType = "BUY"
	SideTypeSell SideType = "SELL"
)

func (object SideType) String() string {
	return string(object)
}
