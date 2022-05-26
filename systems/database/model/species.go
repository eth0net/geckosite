package model

import "database/sql"

// Species stores the details of a species.
type Species struct {
	UUID
	Common
	Name        string         `json:"name" gorm:"unique;not null"`
	LatinName   string         `json:"latinName" gorm:"unique;not null"`
	Description sql.NullString `json:"description"`
	Order       string         `json:"-" gorm:"not null"`
	Type        string         `json:"-" gorm:"not null"`
	Traits      []*Trait       `json:"traits"`
}
