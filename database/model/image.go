package model

import "github.com/google/uuid"

// Image stores the details of an image.
type Image struct {
	Common
	UUID

	Animal   *Animal    `json:"animal"`
	AnimalID *uuid.UUID `json:"-"`
	FileName string     `json:"fileName"`
	FileType string     `json:"fileType"`
	FilePath string     `json:"filePath"`
}
