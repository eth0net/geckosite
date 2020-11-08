package model

import (
	"strconv"
	"time"

	"github.com/google/uuid"
)

// Measurement stores a measurement of an animal.
type Measurement struct {
	AnimalID *uuid.UUID `json:"-" gorm:"not null"`
	Date     *time.Time `json:"date" gorm:"default:Now();not null"`
	Type     string     `json:"type" gorm:"not null;check:type in ('Length','Weight')"`
	Value    uint16     `json:"value" gorm:"not null"`
	Unit     string     `json:"unit" gorm:"not null"`
}

func (m Measurement) String() (s string) {
	val, base := uint64(m.Value), 10
	s = strconv.FormatUint(val, base) + m.Unit
	return
}
