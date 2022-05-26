package model

import (
	"github.com/google/uuid"
)

// UUID field for gorm models.
type UUID struct {
	ID *uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;unique;default:gen_random_uuid();not null"`
}
