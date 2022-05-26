package model

import (
	"strconv"
	"time"

	"github.com/google/uuid"
)

// Measurement stores a measurement of an animal.
type Measurement struct {
	Common
	AnimalID *uuid.UUID `json:"-" gorm:"type:uuid;not null"`
	Date     time.Time  `json:"date" gorm:"default:now();not null"`
	Type     string     `json:"type" gorm:"not null;check:type in ('Length','Weight')"`
	Unit     string     `json:"unit" gorm:"not null"`
	Value    uint16     `json:"value" gorm:"not null"`
}

// String implements Stringer interface.
func (m Measurement) String() string {
	val, base := uint64(m.Value), 10
	return strconv.FormatUint(val, base) + m.Unit
}
