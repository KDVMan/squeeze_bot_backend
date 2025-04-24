package services_helper

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/shopspring/decimal"
	"math"
	"strconv"
	"time"
)

func MustConvertStringToFloat64(value string) float64 {
	if result, err := strconv.ParseFloat(value, 64); err == nil {
		return result
	}

	return 0
}

func MustConvertStringToDecimal(value string) decimal.Decimal {
	d, err := decimal.NewFromString(value)
	if err != nil {
		return decimal.Zero
	}

	return d
}

func MustConvertByteToMd5(text []byte) string {
	hash := md5.Sum(text)

	return hex.EncodeToString(hash[:])
}

func MustConvertStringToMd5(text string) string {
	return MustConvertByteToMd5([]byte(text))
}

func GetPercentFromMinMax(min float64, max float64, fix int) float64 {
	if min == 0 {
		return 0
	}

	result := ((max / min) * 100) - 100

	if fix > 0 {
		return Round(result, fix)
	}

	return result
}

func Round(value float64, decimal int) float64 {
	if decimal == 0 {
		return math.Round(value)
	}

	multiplier := math.Pow(10, float64(decimal))

	return math.Round(value*multiplier) / multiplier
}

func Floor(value float64, decimal int) float64 {
	if decimal == 0 {
		return math.Floor(value)
	}

	multiplier := math.Pow(10, float64(decimal))

	return math.Floor(value*multiplier) / multiplier
}

func MustConvertUnixMillisecondsToString(value int64) string {
	return time.UnixMilli(value).Format("02.01.2006 15:04:05.000")
}
