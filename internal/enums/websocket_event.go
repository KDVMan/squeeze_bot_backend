package enums

type WebsocketEvent string

const (
	WebsocketEventExchangeLimits WebsocketEvent = "exchangeLimits"
	WebsocketEventSymbolList     WebsocketEvent = "symbolList"
	WebsocketEventCurrentPrice   WebsocketEvent = "currentPrice"
	WebsocketEventUser           WebsocketEvent = "user"
	WebsocketEventLeverage       WebsocketEvent = "leverage"
	WebsocketEventBot            WebsocketEvent = "bot"
)
