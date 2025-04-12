package dto

type FileDto struct {
	FileName string `json:"file_name" validate:"required"`
	FileSize int64  `json:"file_size" validate:"required"`
	FileType string `json:"file_type" validate:"required"`
}
