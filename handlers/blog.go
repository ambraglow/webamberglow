package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Blog posts list page
func BlogRoute(c *gin.Context) {
	c.HTML(http.StatusOK, "blog.html", gin.H{
		"title": "./Ambraglow/blog",
		"style": "blog.less",
		"posts": Posts,
	})
}

// Serve blog post by identifying number in request
func GetPost(c *gin.Context) {
	postid, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		HandleNotFound(c)
		panic(err.Error())
	}

	if postid <= len(Posts) {
		var content template.HTML

		for _, post := range Posts {
			if postid == post.Id {
				content = Posts[post.Id].Content
			}
		}

		data := map[string]interface{}{
			"title": "./Ambraglow/blog",
			"style": "post.less",
			"stuff": content,
		}
		/*
			fmt.Println("post id:")
			fmt.Println(postid)
			fmt.Println("post's post id:")
			fmt.Println(Posts[postid].Id)
		*/

		c.HTML(http.StatusOK, "blogpost.html", data)
	} else {
		HandleNotFound(c)
	}
}
