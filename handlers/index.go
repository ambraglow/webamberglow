package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name   string
	Social string
	Quote  string
}

func IndexRoute(c *gin.Context) {
	people := []Person{
		{"luxploit.net", "https://luxploit.net", "MY WIFE"},
		{"Nora", "https://github.com/wandering-nora", "Gay ass bottom"},
		{"Arya", "https://0xarya.gay", "Many such cases"},
		{"WifiCable", "https://github.com/a-little-wifi/", "Remove ribs"},
		{"NSG650", "https://nsg650.github.io/", "Planes are manufactured by Boeing, I have no idea what I'm doing"},
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":  "./Ambraglow/hideout",
		"style":  "index.less",
		"posts":  Posts,
		"people": people,
	})
}

func HobbiesRoute(c *gin.Context) {
	parameter := c.Param("hobby")
	switch parameter {
	case "photography":
		c.HTML(http.StatusOK, "photography.html", gin.H{
			"title": "./Ambraglow/Photography",
			"style": "post.less",
		})
	default:
		HandleNotFound(c)
	}

}
