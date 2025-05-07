package handlers

import (
	"github.com/gin-gonic/gin"
)

func HandleNotFound(c *gin.Context) {
	data := map[string]interface{}{
		"title": "./Ambraglow/404",
		"style": "index.less",
	}
	c.HTML(404, "404.html", data)
}
