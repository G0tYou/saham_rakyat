package resource

import (
	"time"

	"gorm.io/gorm"
)

type OrderHistories struct {
	Id          int            `json:"id" gorm:"primaryKey;not null"`
	UserId      int            `json:"user_id" gorm:"not null"`
	OrderItem   int            `json:"order_item" gorm:"not null"`
	Description string         `json:"description" gorm:"not null"`
	CreatedAt   time.Time      `json:"created_at" gorm:"default:CURRENT_TIMESTAMP()"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP()"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"default:"`
}
