package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BlogRoute(c *gin.Context) {
	c.HTML(http.StatusOK, "blog.html", gin.H{
		"title": "./Ambraglow/blog",
		"style": "blog.less",
		"posts": Posts,
	})
}

func GetPost(c *gin.Context) {
	postid, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		FourfourRoute(c)
		panic(err.Error())
	}

	if !(postid >= len(Posts)) {
		data := map[string]interface{}{
			"title": "./Ambraglow/blog",
			"style": "post.less",
			"stuff": Posts[postid].Content,
		}

		c.HTML(http.StatusOK, "blogpost.html", data)
	} else {
		FourfourRoute(c)
	}
}
