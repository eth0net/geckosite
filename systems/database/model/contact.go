package model

import "database/sql"

// Contact stores the details of a sales contact.
// Could be a shop or individual we have bought from or sold to.
type Contact struct {
	UUID
	Common

	Name    string         `json:"name" gorm:"unique;not null"`
	Phone   sql.NullString `json:"phone"`
	Email   sql.NullString `json:"email"`
	Website sql.NullString `json:"website"`
	Note    sql.NullString `json:"note"`

	Transactions []*Transaction `json:"transactions"`
}
