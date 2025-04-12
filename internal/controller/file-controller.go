package controller

import (
	"boilerplate_go/internal/dto"
	"boilerplate_go/internal/usecase"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
)

type FileController interface {
	Upload(ctx echo.Context) error
	Download(ctx echo.Context) error
	Delete(ctx echo.Context) error
	Search(ctx echo.Context) error
}

type fileController struct {
	fileUseCase usecase.FileUseCase
}

func NewFileController(fileUseCase usecase.FileUseCase) FileController {
	return &fileController{
		fileUseCase: fileUseCase,
	}
}

// Upload implements FileController.
func (f *fileController) Upload(ctx echo.Context) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		log.Error("Error getting file from form: ", err)
		return ctx.JSON(400, dto.NewBaseResponse(400, "Error getting file from form", err.Error()))
	}

	src, err := file.Open()
	if err != nil {
        return ctx.JSON(400, dto.NewBaseResponse(400, "Error getting file from form", err.Error()))
    }
    defer src.Close()

	fileDto := dto.FileDto{
		FileName: file.Filename,
		FileSize: file.Size,
		FileType: file.Header.Get("Content-Type"),
	}
	
	if err := f.fileUseCase.Upload(ctx.Request().Context(), src, &fileDto); err != nil {
		log.Error("Error uploading file: ", err)
		return ctx.JSON(500, dto.NewBaseResponse(500, "Error uploading file", err.Error()))
	}

	return ctx.JSON(http.StatusOK, dto.NewBaseResponse(200, "Success", nil))
}

// Delete implements FileController.
func (f *fileController) Delete(ctx echo.Context) error {
	panic("unimplemented")
}

// Download implements FileController.
func (f *fileController) Download(ctx echo.Context) error {
	panic("unimplemented")
}

// Search implements FileController.
func (f *fileController) Search(ctx echo.Context) error {
	return nil
}
