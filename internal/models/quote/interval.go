package models_quote

type Interval struct {
	Name     string `json:"name" validate:"required"`
	Seconds  int    `json:"seconds" validate:"required"`
	Active   bool   `json:"active"`
	Favorite bool   `json:"favorite"`
}

func IntervalLoadDefault() []*Interval {
	return []*Interval{
		{Name: "1m", Seconds: 60, Active: true, Favorite: true},
		{Name: "3m", Seconds: 60 * 3, Active: false, Favorite: true},
		{Name: "5m", Seconds: 60 * 5, Active: false, Favorite: true},
		{Name: "15m", Seconds: 60 * 15, Active: false, Favorite: true},
		{Name: "30m", Seconds: 60 * 30, Active: false, Favorite: true},
		{Name: "1h", Seconds: 60 * 60, Active: false, Favorite: true},
		{Name: "2h", Seconds: 60 * 60 * 2, Active: false, Favorite: false},
		{Name: "4h", Seconds: 60 * 60 * 4, Active: false, Favorite: false},
		{Name: "6h", Seconds: 60 * 60 * 6, Active: false, Favorite: false},
		{Name: "8h", Seconds: 60 * 60 * 8, Active: false, Favorite: false},
		{Name: "12h", Seconds: 60 * 60 * 12, Active: false, Favorite: false},
		{Name: "1d", Seconds: 60 * 60 * 24, Active: false, Favorite: true},
		{Name: "3d", Seconds: 60 * 60 * 24 * 3, Active: false, Favorite: false},
		{Name: "1w", Seconds: 60 * 60 * 24 * 7, Active: false, Favorite: false},
		{Name: "1M", Seconds: 60 * 60 * 24 * 30, Active: false, Favorite: false},
	}
}
