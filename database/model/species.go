package model

// Species stores the details of a species.
type Species struct {
	UUID
	Common
	Name        string   `json:"name" gorm:"unique;not null"`
	LatinName   string   `json:"latinName" gorm:"unique;not null"`
	Description string   `json:"description"`
	Order       string   `json:"-"`
	Path        string   `json:"-"`
	Traits      []*Trait `json:"traits"`
	// Images      []string `json:"images"`
}
