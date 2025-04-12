package usecase

import (
	"boilerplate_go/internal/dto"
	"boilerplate_go/internal/utils"
	"context"
	"io"
	"mime/multipart"
)

type FileUseCase interface {
	Upload(ctx context.Context, src multipart.File, fileDto *dto.FileDto) error
	Download(ctx context.Context, objectName string) (io.Reader, error)
	Delete(ctx context.Context, objectName string) error
	Search(ctx context.Context, prefix string) ([]string, error)
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
func (f *fileUseCase) Delete(ctx context.Context, objectName string) error {
	return utils.FileClient.Delete(ctx, objectName)
}

// Download implements FileUseCase.
func (f *fileUseCase) Download(ctx context.Context, objectName string) (io.Reader, error) {
	fileContent, err := utils.FileClient.Download(ctx, objectName)
	if err != nil {
		return nil, err
	}

	return fileContent, nil
}

// Search implements FileUseCase.
func (f *fileUseCase) Search(ctx context.Context, prefix string) ([]string, error) {
	objectCh := utils.FileClient.List(ctx, prefix)

	var files []string
	for object := range objectCh {
		if object.Err != nil {
			return nil, object.Err
		}
		files = append(files, object.Key)
	}

	return files, nil
}
