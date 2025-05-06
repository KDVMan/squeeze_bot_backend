package services_bot

import (
	"backend/internal/enums"
	"log"
	"sync"
	"time"
)

type guardState struct {
	active         bool
	stopDirection  enums.TradeDirection
	resetTimestamp int64 // Время, когда защита может быть снята
}

var (
	guardStates        = make(map[string]*guardState)
	guardMutex         sync.Mutex
	guardPercentLimit  = 1.0
	guardCooldownTicks = 2
)

func (object *botServiceImplementation) RunGuardChannel() {
	tickDuration := int64(time.Minute.Seconds()) * 1000

	for symbol := range object.guardChannel {
		quotes := object.quoteRepositoryService().GetBySymbol(symbol)
		currentQuote := quotes[len(quotes)-1]
		openPrice := currentQuote.PriceOpen
		highPrice := currentQuote.PriceHigh
		lowPrice := currentQuote.PriceLow

		if openPrice <= 0 {
			continue
		}

		highPercent := ((highPrice - openPrice) / openPrice) * 100
		lowPercent := ((openPrice - lowPrice) / openPrice) * 100
		now := time.Now().UnixMilli()

		guardMutex.Lock()

		state, exists := guardStates[symbol]
		if !exists {
			state = &guardState{}
			guardStates[symbol] = state
		}

		priceSpiked := false
		var direction enums.TradeDirection

		if highPercent >= guardPercentLimit {
			priceSpiked = true
			direction = enums.TradeDirectionShort
		} else if lowPercent >= guardPercentLimit {
			priceSpiked = true
			direction = enums.TradeDirectionLong
		}

		if priceSpiked {
			if !state.active || state.stopDirection != direction {
				// log.Printf("[GUARD] %s: spike detected %.2f%% %s — stop %s until %s",
				// 	symbol,
				// 	math.Max(highPercent, lowPercent),
				// 	direction,
				// 	direction,
				// 	time.UnixMilli(now+int64(guardCooldownTicks)*tickDuration).Format(time.RFC3339),
				// )
			} else {
				// log.Printf("[GUARD] %s: spike again %.2f%% — extended stop %s",
				// 	symbol,
				// 	math.Max(highPercent, lowPercent),
				// 	direction)
			}

			state.active = true
			state.stopDirection = direction
			state.resetTimestamp = now + int64(guardCooldownTicks)*tickDuration
		} else if state.active && now >= state.resetTimestamp {
			log.Printf("[GUARD] %s: market normalized — stop %s removed", symbol, state.stopDirection)

			state.active = false
			state.stopDirection = ""
			state.resetTimestamp = 0
		}

		guardMutex.Unlock()
	}
}

func (object *botServiceImplementation) isGuardActive(symbol string, direction enums.TradeDirection) bool {
	guardMutex.Lock()
	defer guardMutex.Unlock()

	state, exists := guardStates[symbol]

	return exists && state.active && state.stopDirection == direction
}

func (object *botServiceImplementation) GetGuardChannel() chan string {
	return object.guardChannel
}
