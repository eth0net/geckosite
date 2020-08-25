package model

import "github.com/google/uuid"

// Trait of a species, such as colouring or pattern change.
// Could be a single gene or a combination of genes.
type Trait struct {
	UUID
	Common
	Name        string     `json:"name" gorm:"uniqueIndex:species_trait"`
	Description string     `json:"description"`
	Species     *Species   `json:"species"`
	SpeciesID   *uuid.UUID `json:"-" gorm:"uniqueIndex:species_trait"`
}
