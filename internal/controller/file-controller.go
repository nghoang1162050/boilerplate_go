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
// @Summary Upload File
// @Description Uploads a file using multipart form data.
// @Tags Files
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File to upload"
// @Success 200 {object} dto.BaseResponse "Success"
// @Failure 400 {object} dto.BaseResponse "Bad Request"
// @Failure 500 {object} dto.BaseResponse "Internal Server Error"
// @Router /files/upload [post]
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
// @Summary Delete File
// @Description Deletes a file with the specified object name.
// @Tags Files
// @Accept json
// @Produce json
// @Param object_name query string true "Object name"
// @Success 200 {object} dto.BaseResponse "File deleted successfully"
// @Failure 400 {object} dto.BaseResponse "Bad Request"
// @Failure 500 {object} dto.BaseResponse "Internal Server Error"
// @Router /files/delete [delete]
func (f *fileController) Delete(ctx echo.Context) error {
	objectName := ctx.QueryParam("object_name")
	if objectName == "" {
		return ctx.JSON(http.StatusBadRequest, dto.NewBaseResponse(400, "Object name is required", nil))
	}

	if err := f.fileUseCase.Delete(ctx.Request().Context(), objectName); err != nil {
		log.Error("Error deleting file: ", err)
		return ctx.JSON(http.StatusInternalServerError, dto.NewBaseResponse(500, "Error deleting file", err.Error()))
	}

	return ctx.JSON(http.StatusOK, dto.NewBaseResponse(200, "File deleted successfully", nil))
}

// Download implements FileController.
// @Summary Download File
// @Description Downloads a file identified by the object name.
// @Tags Files
// @Accept json
// @Produce octet-stream
// @Param object_name query string true "Object name"
// @Success 200 {file} file "File content"
// @Failure 400 {object} dto.BaseResponse "Bad Request"
// @Failure 404 {object} dto.BaseResponse "File Not Found"
// @Failure 500 {object} dto.BaseResponse "Internal Server Error"
// @Router /files/download [get]
func (f *fileController) Download(ctx echo.Context) error {
	objectName := ctx.QueryParam("object_name")
	if objectName == "" {
		return ctx.JSON(http.StatusBadRequest, dto.NewBaseResponse(400, "Object name is required", nil))
	}

	file, err := f.fileUseCase.Download(ctx.Request().Context(), objectName)
	if err != nil {
		log.Error("Error downloading file: ", err)
		return ctx.JSON(http.StatusInternalServerError, dto.NewBaseResponse(500, "Error downloading file", err.Error()))
	}
	if file == nil {
		return ctx.JSON(http.StatusNotFound, dto.NewBaseResponse(404, "File not found", nil))
	}

	return ctx.Stream(http.StatusOK, "application/octet-stream", file)
}

// Search implements FileController.
// @Summary Search Files
// @Description Searches and retrieves a list of files that match the given prefix.
// @Tags Files
// @Accept json
// @Produce json
// @Param prefix query string true "Prefix for file search"
// @Success 200 {object} dto.BaseResponse "Success"
// @Failure 400 {object} dto.BaseResponse "Bad Request"
// @Failure 404 {object} dto.BaseResponse "No files found"
// @Failure 500 {object} dto.BaseResponse "Internal Server Error"
// @Router /files/search [get]
func (f *fileController) Search(ctx echo.Context) error {
	prefix := ctx.QueryParam("prefix")
	if prefix == "" {
		return ctx.JSON(http.StatusBadRequest, dto.NewBaseResponse(400, "Prefix is required", nil))
	}

	files, err := f.fileUseCase.Search(ctx.Request().Context(), prefix)
	if err != nil {
		log.Error("Error searching files: ", err)
		return ctx.JSON(http.StatusInternalServerError, dto.NewBaseResponse(500, "Error searching files", err.Error()))
	}
	if len(files) == 0 {
		return ctx.JSON(http.StatusNotFound, dto.NewBaseResponse(404, "No files found", nil))
	}

	return ctx.JSON(http.StatusOK, dto.NewBaseResponse(200, "Success", files))
}
