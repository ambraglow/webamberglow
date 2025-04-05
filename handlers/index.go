package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexRoute(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Welcome to Ambra's website!",
		"style": "index.less",
	})
}
