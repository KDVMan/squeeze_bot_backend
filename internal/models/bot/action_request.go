package models_bot

import enums_bot "backend/internal/enums/bot"

type ActionRequestModel struct {
	Action enums_bot.Action `json:"action" validate:"required,botAction"`
}
