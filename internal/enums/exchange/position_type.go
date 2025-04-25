package enums_exchange

type PositionType string

const (
	PositionTypeLong  PositionType = "LONG"
	PositionTypeShort PositionType = "SHORT"
	PositionTypeBoth  PositionType = "BOTH"
)

func (object PositionType) String() string {
	return string(object)
}
