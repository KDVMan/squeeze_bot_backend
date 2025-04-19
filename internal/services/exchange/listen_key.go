package services_exchange

import (
	"context"
	"time"
)

func (object *exchangeServiceImplementation) GetListenKey() (string, error) {
	if object.listenKey == "" {
		listenKey, err := object.client.NewStartUserStreamService().Do(context.Background())
		if err != nil {
			return "", err
		}

		object.listenKey = listenKey
		object.renewListenKey()
	}

	return object.listenKey, nil
}

func (object *exchangeServiceImplementation) renewListenKey() {
	if object.stopRenewListenKey != nil {
		close(object.stopRenewListenKey)
	}

	object.stopRenewListenKey = make(chan struct{})

	go func() {
		ticker := time.NewTicker(30 * time.Minute)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				err := object.client.NewKeepaliveUserStreamService().ListenKey(object.listenKey).Do(context.Background())
				if err != nil {
					object.listenKey = ""
					object.stopRenewListenKey = nil
					return
				}
			case <-object.stopRenewListenKey:
				return
			}
		}
	}()
}

func (object *exchangeServiceImplementation) DeleteListenKey() error {
	if object.listenKey == "" {
		return nil
	}

	err := object.client.NewCloseUserStreamService().ListenKey(object.listenKey).Do(context.Background())
	if err != nil {
		return err
	}

	if object.stopRenewListenKey != nil {
		close(object.stopRenewListenKey)
		object.stopRenewListenKey = nil
	}

	object.listenKey = ""

	return nil
}
