package enums_bot

type EntryStatus string

const (
	EntryStatusNormal EntryStatus = "normal" // нормальная работа
	EntryStatusNext   EntryStatus = "next"   // ждем следующую свечу (не смогли зайти)
)

func (object EntryStatus) String() string {
	return string(object)
}
