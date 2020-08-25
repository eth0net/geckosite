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
func (u *UUID) BeforeCreate(tx *gorm.DB) error {
	if u.ID != nil {
		return nil
	}

	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	u.ID = &id

	return nil
}
