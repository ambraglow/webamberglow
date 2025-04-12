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
		{"luxploit.net", "https://bsky.app/profile/luxploit.net", "MY WIFE"},
		{"Nora", "", ""},
		{"Arya", "https://0xarya.gay", ""},
		{"WifiCable", "", ""},
		{"NSG650", "https://bsky.app/profile/nsg650.bsky.social", "Planes are manufactured by Boeing, I have no idea what I'm doing"},
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
			"title": "./Ambraglow/hideout",
			"style": "post.less",
		})
	default:
		FourfourRoute(c)
	}

}
