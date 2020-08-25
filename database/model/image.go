package model

// Image stores the details of an image.
type Image struct {
	Common
	UUID

	FileName string `json:"fileName"`
	FileType string `json:"fileType"`
	FilePath string `json:"filePath"`
}
