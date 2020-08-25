package model

import (
	"time"

	"github.com/google/uuid"
)

// Animal stores the details for an animal.
type Animal struct {
	UUID
	Common

	// Details specific to the animal.
	Reference   string     `json:"reference,omitEmpty"`
	Name        string     `json:"name,omitEmpty"`
	Description string     `json:"description,omitEmpty"`
	Images      []*Image   `json:"images" gorm:"many2many:animal_images;"`
	Species     *Species   `json:"species"`
	SpeciesID   *uuid.UUID `json:"-" gorm:"not null"`
	Sex         string     `json:"sex" gorm:"default:Unknown;not null;check:sex IN ('Male','Female','Unknown')"`
	Status      string     `json:"status" gorm:"default:Pet;not null;check:status IN ('Pet','Breeder','Sale')"`

	// Important dates for our records.
	DateLaid    *time.Time `json:"dateLaid"`
	DateHatched *time.Time `json:"dateHatched"`
	DateBought  *time.Time `json:"dateBought"`
	DateSold    *time.Time `json:"dateSold"`

	// Relations to other animals.
	Father   *Animal    `json:"father"`
	FatherID *uuid.UUID `json:"-"`
	Mother   *Animal    `json:"mother"`
	MotherID *uuid.UUID `json:"-"`

	// PossibleFathers []*Animal `json:"possibleFathers" gorm:"many2many:animal_parents;foreignKey:id,sex;"`
	// PossibleMothers []*Animal `json:"possibleMothers" gorm:"many2many:animal_parents;foreignKey:id,sex;"`
	// Children []Animal

	// Traits         []*Gene `json:"traits"`
	// PossibleTraits []*Gene `json:"possibleTraits"`
}
