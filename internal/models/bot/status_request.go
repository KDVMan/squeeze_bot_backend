package models_bot

import (
	enums_bot "backend/internal/enums/bot"
)

type StatusRequestModel struct {
	ID     uint             `json:"id" validate:"required,gt=0"`
	Status enums_bot.Status `json:"status" validate:"required,botStatus"`
}
