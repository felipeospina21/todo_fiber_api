package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TodoItem struct {
	gorm.Model
	ItemId  uuid.UUID `json:"item_id"`
	Message string    `json:"message"`
	Done    bool      `json:"done"`
}
