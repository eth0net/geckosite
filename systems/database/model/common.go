package model

import (
	"time"

	"gorm.io/gorm"
)

// Common fields for gorm models.
type Common struct {
	CreatedAt time.Time      `json:"createdAt" gorm:"default:now();not null;autoCreateTime"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"default:now();not null;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
