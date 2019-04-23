package routing

import (
	"fmt"
	"html/template"
	"strings"
	"time"

	commonConfig "common/config"
)

func templateFuncMap(config *commonConfig.Config) template.FuncMap {
	return template.FuncMap{
		"url":  url(config),
		"year": year,
	}
}

func url(config *commonConfig.Config) func(string) string {
	return func(path string) string {
		path = strings.TrimPrefix(path, "/")
		hostname := config.HTTP.Address(true)
		return fmt.Sprintf("%s/%s", hostname, path)
	}
}

func year() int {
	return time.Now().Year()
}
