package enums_bot

type BotEvent string

const (
	BotEventStatusOpenLimit BotEvent = "status_open_limit"
	BotEventOpenLimit       BotEvent = "open_limit"
	BotEventCancelLimit     BotEvent = "cancel_limit"
	BotEventCloseLimit      BotEvent = "close_limit"
)

func (object BotEvent) String() string {
	return string(object)
}
