package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UUID field for gorm models.
type UUID struct {
	ID *uuid.UUID `json:"id" gorm:"primaryKey;unique;not null"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (u *UUID) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID != nil {
		return
	}

	id, err := uuid.NewUUID()
	if err != nil {
		return
	}
	u.ID = &id

	return
}
