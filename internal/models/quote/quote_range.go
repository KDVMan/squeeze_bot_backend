package models_quote

import (
	"backend/internal/enums"
	"time"
)

type QuoteRangeModel struct {
	QuotesLimit int64
	TimeFrom    int64
	TimeTo      int64
	TimeStep    int64
	Iterations  int
}

func GetTimeRange(interval enums.Interval, limit int64) (int64, int64) {
	currentTime := time.Now().UnixMilli()
	timeFrom := currentTime - int64(limit*60*1000)
	timeTo := currentTime

	// currentTime := time.Now().UnixMilli()
	// timeTo := ((currentTime/60000)+1)*60000 - 1
	// intervalMinutes := enums.IntervalSeconds(interval) / 60
	// neededMinutes := limit*intervalMinutes + 1
	// timeFrom := timeTo - neededMinutes*60000 + 1

	return timeFrom, timeTo
}

func GetRange(limit int64, timeFrom int64, timeTo int64, milliseconds int64) *QuoteRangeModel {
	timeRange := (timeTo - timeFrom) + 1000
	total := timeRange / milliseconds

	if limit > total {
		total = limit
		// limit = total
	}

	return &QuoteRangeModel{
		QuotesLimit: limit,
		TimeFrom:    timeFrom,
		TimeTo:      timeTo,
		TimeStep:    milliseconds * limit,
		Iterations:  int((total + limit - 1) / limit),
	}
}
