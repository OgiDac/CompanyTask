package controllers

import (
	"io"
	"net/http"
	"strconv"

	"github.com/OgiDac/CompanyTask/domain"
	"github.com/gin-gonic/gin"
)

type FileController struct {
	FileUseCase domain.FileUseCase
}

// UploadFile godoc
// @Summary      Upload a file for a user
// @Description  Uploads a file linked to the user ID
// @Tags         files
// @Accept       multipart/form-data
// @Produce      json
// @Param        id path int true "User ID"
// @Param        file formData file true "File to upload"
// @Success      200 {object} map[string]string
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /public/api/files/{id} [post]
// @Security     BearerAuth
func (fc *FileController) UploadFile(c *gin.Context) {
	idParam := c.Param("id")
	userID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to get file"})
		return
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to open file"})
		return
	}
	defer src.Close()

	data, err := io.ReadAll(src)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read file"})
		return
	}

	err = fc.FileUseCase.UploadFile(c.Request.Context(), uint(userID), file.Filename, file.Header.Get("Content-Type"), data)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "file uploaded successfully"})
}


// DownloadFile godoc
// @Summary      Download a user file
// @Description  Downloads a file by its ID
// @Tags         files
// @Produce      application/octet-stream
// @Param        id path string true "File ID"
// @Success      200 {file} file
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Router       /public/api/files/{id} [get]
// @Security     BearerAuth
func (fc *FileController) DownloadFile(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "missing id"})
		return
	}

	ctx := c.Request.Context()
	file, err := fc.FileUseCase.GetFileByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Header("Content-Disposition", "attachment; filename="+file.Filename)
	c.Data(http.StatusOK, "application/octet-stream", file.Data)
}

// GetFilesByUser godoc
// @Summary      Get all files for a user
// @Description  Returns file IDs and names for a user ID
// @Tags         files
// @Produce      json
// @Param        id path int true "User ID"
// @Success      200 {array} domain.UserFileMeta
// @Failure      400 {object} map[string]string
// @Router       /public/api/files/user/{id} [get]
// @Security     BearerAuth
func (fc *FileController) GetFilesByUser(c *gin.Context) {
	idParam := c.Param("id")
	userID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	files, err := fc.FileUseCase.GetFilesByUserID(c.Request.Context(), uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, files)
}

// DeleteFilesByUser godoc
// @Summary      Delete all files for a user
// @Description  Deletes all files linked to a user ID
// @Tags         files
// @Produce      json
// @Param        id path int true "User ID"
// @Success      200 {object} map[string]string
// @Failure      400 {object} map[string]string
// @Router       /public/api/files/user/{id} [delete]
// @Security     BearerAuth
func (fc *FileController) DeleteFilesByUser(c *gin.Context) {
	idParam := c.Param("id")
	userID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	err = fc.FileUseCase.DeleteFilesByUserID(c.Request.Context(), uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "all files deleted"})
}
