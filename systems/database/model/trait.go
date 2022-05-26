package model

import (
	"database/sql"

	"github.com/google/uuid"
)

// Trait of a species, such as colouring or pattern change.
// Could be a single gene or a combination of genes.
type Trait struct {
	UUID
	Common
	Name        string         `json:"name" gorm:"not null;uniqueIndex:idx_species_trait"`
	Description sql.NullString `json:"description"`
	Species     *Species       `json:"species"`
	SpeciesID   *uuid.UUID     `json:"-" gorm:"type:uuid;not null;uniqueIndex:idx_species_trait"`
}
