package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func slice(a []Blogpost, start, end int) []Blogpost {
	return a[start:end]
}

func IndexRoute(c *gin.Context) {
	Posts := BlogPostMetadata()

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "./Ambraglow/hideout",
		"style": "index.less",
		"posts": slice(Posts, 0, 3),
	})
}
