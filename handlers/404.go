package handlers

import (
	"github.com/gin-gonic/gin"
)

func FourfourRoute(c *gin.Context) {
	data := map[string]interface{}{
		"title": "./Ambraglow/blog",
		"style": "index.less",
	}
	c.HTML(404, "404.html", data)
}
