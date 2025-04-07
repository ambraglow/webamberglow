package main

import (
	"log"
	"path/filepath"
	"website/handlers"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	handlers.BlogMarkdownInit()

	router.Static("/assets", "./assets")
	router.HTMLRender = loadTemplates("templates")

	router.GET("/", handlers.IndexRoute)
	router.GET("/blog", handlers.BlogRoute)

	if err := router.Run(":3000"); err != nil {
		log.Fatal(err)
	}
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/pages/*.html")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}
