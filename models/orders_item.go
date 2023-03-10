package resource

import (
	"time"

	"gorm.io/gorm"
)

type OrdersItem struct {
	Id        int            `json:"id" gorm:"primaryKey;not null"`
	Name      string         `json:"name" gorm:"not null"`
	Price     float64        `json:"price" gorm:"not null"`
	ExpiredAt time.Time      `json:"expired_at" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at" gorm:"default:CURRENT_TIMESTAMP()"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP()"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"default:"`
}
