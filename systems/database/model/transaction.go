package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// Transaction stores the details of a single transaction.
type Transaction struct {
	UUID
	Common

	Date  time.Time      `json:"date" gorm:"type:date;not null"`
	Note  sql.NullString `json:"note"`
	Type  string         `json:"type" gorm:"not null;check type in ('Purchase','Sale')"`
	Value uint32         `json:"value" gorm:"not null"`

	Animals   []*Animal  `json:"animals" gorm:"many2many:transaction_animals"`
	Contact   *Contact   `json:"contact"`
	ContactID *uuid.UUID `json:"-" gorm:"type:uuid;not null"`
}
