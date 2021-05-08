package model

import "github.com/google/uuid"

// AnimalParent is the custom join table to store parent-child relations for Animals.
type AnimalParent struct {
	Common
	ChildID   *uuid.UUID `json:"childID" gorm:"type:uuid;primaryKey;not null;uniqueIndex:idx_parent_child"`
	ParentID  *uuid.UUID `json:"parentID" gorm:"type:uuid;primaryKey;not null;uniqueIndex:idx_parent_child"`
	ParentSex string     `json:"parentSex" gorm:"primaryKey;not null;uniqueIndex:idx_parent_child;check:parent_sex IN ('Male','Female')"`
}
