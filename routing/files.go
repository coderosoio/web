package routing

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// FileHandler handles static files.
type FileHandler struct {
	options *Options
}

// NewFileHandler returns an instance of `FileHandler`.
func NewFileHandler(opts ...Option) *FileHandler {
	options := newOptions(opts...)
	return &FileHandler{
		options: options,
	}
}

// Download files.
func (h *FileHandler) Download(ctx *gin.Context) {
	var err error
	name := ctx.Param("filename")
	filename := fmt.Sprintf("static/%s", name)
	if filename, err = filepath.Abs(filename); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	if _, err = os.Stat(filename); os.IsNotExist(err) {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.Header("Content-Description", "File Transfer")
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", name))
	ctx.Data(http.StatusOK, "application/octet-stream", b)
}
