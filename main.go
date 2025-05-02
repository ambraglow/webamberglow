package main

import (
	"flag"
	"log"
	"path/filepath"
	"website/handlers"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

var (
	bindPort = flag.String("p", "port", "port to bind webserver to")
)

func main() {
	flag.Parse()
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	handlers.BlogMarkdownInit()

	router.Static("/assets", "./assets")
	router.HTMLRender = loadTemplates("templates")

	router.NoRoute(handlers.HandleNotFound)

	router.GET("/", handlers.IndexRoute)
	router.GET("/blog", handlers.BlogRoute)
	router.GET("/blog/:Id", handlers.GetPost)
	router.GET("/hobbies/:hobby", handlers.HobbiesRoute)

	if err := router.Run(":" + *bindPort); err != nil {
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
