package routing

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomeHandler is the router handler for home.
type HomeHandler struct {
	options *Options
}

// NewHomeHandler returns an instance of `HomeHandler`.
func NewHomeHandler(opts ...Option) *HomeHandler {
	options := newOptions(opts...)
	return &HomeHandler{
		options: options,
	}
}

// Index is the handler for `GET /`.
func (h *HomeHandler) Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "home/index", gin.H{
		"title": "Home",
	})
}
