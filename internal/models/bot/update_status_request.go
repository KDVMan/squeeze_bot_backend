package models_bot

import (
	enums_bot "backend/internal/enums/bot"
)

type UpdateStatusRequestModel struct {
	ID     uint             `json:"id" validate:"required,gt=0"`
	Status enums_bot.Status `json:"status" validate:"required,botStatus"`
}
