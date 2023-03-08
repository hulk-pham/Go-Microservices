package common

import (
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadLocal(c *gin.Context, file *multipart.FileHeader) {
	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
		return
	}
}

func RemoveFile(filepath string) (err error) {
	return os.Remove(filepath)
}
