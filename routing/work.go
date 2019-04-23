package routing

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// WorkHandler is the router handler for content.
type WorkHandler struct {
	options *Options
}

// NewWorkHandler returns an instance of `WorkHandler`.
func NewWorkHandler(opts ...Option) *WorkHandler {
	options := newOptions(opts...)
	return &WorkHandler{
		options: options,
	}
}

// Work handles `GET /work/:company/:project`.
func (h *WorkHandler) Work(ctx *gin.Context) {
	company := ctx.Param("company")
	project := ctx.Param("project")
	view := fmt.Sprintf("%s/%s", company, project)
	log.Printf("view: %s", view)
	ctx.HTML(http.StatusOK, view, gin.H{})
}
