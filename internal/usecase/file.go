package usecase

import (
	"boilerplate_go/internal/dto"
	"boilerplate_go/internal/utils"
	"context"
	"mime/multipart"
)

type FileUseCase interface {
	Upload(ctx context.Context, src multipart.File, fileDto *dto.FileDto) error
	Download() error
	Delete() error
	Search() error
}

type fileUseCase struct {
	// fileRepository repository.FileRepository
}

func NewFileUseCase() FileUseCase {
	return &fileUseCase{
		// fileRepository: fileRepository,
	}
}

// Upload implements FileUseCase.
func (f *fileUseCase) Upload(ctx context.Context, src multipart.File, fileDto *dto.FileDto) error {
	// TODO: gaurd against file size and type
	return utils.FileClient.Upload(ctx, fileDto.FileName, src, fileDto.FileSize, fileDto.FileType)
}

// Delete implements FileUseCase.
func (f *fileUseCase) Delete() error {
	panic("unimplemented")
}

// Download implements FileUseCase.
func (f *fileUseCase) Download() error {
	panic("unimplemented")
}

// Search implements FileUseCase.
func (f *fileUseCase) Search() error {
	panic("unimplemented")
}
