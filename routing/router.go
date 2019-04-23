package routing

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"

	commonConfig "common/config"
)

// NewRouter returns a new Gin router.
func NewRouter(config *commonConfig.Config) (*gin.Engine, error) {
	router := gin.Default()
	router.RedirectTrailingSlash = true

	if commonConfig.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}

	renderer, err := setupRenderer(config)
	if err != nil {
		return nil, err
	}

	router.HTMLRender = renderer
	router.StaticFS("/assets", http.Dir("static"))
	router.StaticFile("/browserconfig.xml", "static/browserconfig.xml")
	router.StaticFile("/apple-touch-icon.png", "static/images/apple-touch-icon.png")
	router.StaticFile("/apple-touch-icon-precomposed.png", "static/images/apple-touch-icon-precomposed.png")
	router.StaticFile("/favicon.ico", "static/images/favicon.ico")

	homeHandler := NewHomeHandler(
		Config(config),
	)
	workHandler := NewWorkHandler(
		Config(config),
	)

	home := router.Group("/")
	{
		home.GET("/", homeHandler.Index)
	}

	work := router.Group("/work")
	{
		work.GET("/:company/:project", workHandler.Work)
	}

	return router, nil
}

func setupRenderer(config *commonConfig.Config) (*gintemplate.TemplateEngine, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	templatePath, err := filepath.Abs(fmt.Sprintf("%s/template", wd))
	if err != nil {
		return nil, err
	}

	renderer := gintemplate.New(gintemplate.TemplateConfig{
		Root:         templatePath,
		Extension:    ".html",
		Master:       "layout/master",
		Funcs:        templateFuncMap(config),
		DisableCache: !commonConfig.IsProduction(),
	})

	return renderer, nil
}
