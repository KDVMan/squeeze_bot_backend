package services_exchange_limit

import (
	"net/http"
	"strconv"
	"sync"
)

type Transport struct {
	Value http.RoundTripper
}

var (
	limits = make(map[string]int)
	lock   sync.RWMutex
)

func (object *Transport) RoundTrip(request *http.Request) (*http.Response, error) {
	response, err := object.Value.RoundTrip(request)
	if err != nil {
		return nil, err
	}

	lock.Lock()
	defer lock.Unlock()

	update(response.Header)

	return response, nil
}

func update(headers http.Header) {
	names := []string{
		"x-mbx-used-weight",
		"x-mbx-used-weight-1m",
		"x-sapi-used-ip-weight-1m",
		"x-mbx-order-count-1s",
		"x-mbx-order-count-1m",
		"x-mbx-order-count-1h",
		"x-mbx-order-count-1d",
	}

	for _, name := range names {
		if valueStr := headers.Get(name); valueStr != "" {
			if value, err := strconv.Atoi(valueStr); err == nil {
				limits[name] = value
			}
		}
	}
}

func GetLimits() map[string]int {
	lock.RLock()
	defer lock.RUnlock()

	return limits
}
