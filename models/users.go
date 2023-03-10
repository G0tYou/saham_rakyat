package resource

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Id         int            `json:"id" gorm:"primaryKey;not null"`
	FullName   string         `json:"full_name" gorm:"not null"`
	FirstOrder time.Time      `json:"first_order" gorm:"default:"`
	CreatedAt  time.Time      `json:"created_at" gorm:"default:CURRENT_TIMESTAMP()"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP()"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"default:"`
}
