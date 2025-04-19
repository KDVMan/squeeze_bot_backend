package models_channel

import "backend/internal/enums"

type BroadcastChannelModel struct {
	Event enums.WebsocketEvent `json:"event"`
	Data  interface{}          `json:"data"`
}
