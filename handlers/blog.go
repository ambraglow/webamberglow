package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BlogRoute(c *gin.Context) {
	Posts := BlogPostMetadata()

	c.HTML(http.StatusOK, "blog.html", gin.H{
		"title": "./Ambraglow/blog",
		"style": "blog.less",
		"posts": Posts,
	})
}
