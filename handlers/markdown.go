package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strconv"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

var MD goldmark.Markdown

var Posts []Blogpost

type Blogpost struct {
	Title   string
	Summary string
	Tags    []string
	Id      int
	Content template.HTML
}

func BlogMarkdownInit() {
	MD = goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
			extension.Linkify,
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
	Posts = BlogPosts()
}

func BlogPosts() []Blogpost {
	files, err := filepath.Glob("markdown/*.md")
	if err != nil {
		panic(err.Error())
	}
	var posts []Blogpost

	for _, file := range files {
		var buf bytes.Buffer
		var blogpost Blogpost

		context := parser.NewContext()

		fileread, err := os.ReadFile(file)
		if err != nil {
			panic(err.Error())
		}

		if err := MD.Convert(fileread, &buf, parser.WithContext(context)); err != nil {
			panic(err.Error())
		}

		//	header
		metadata := meta.Get(context)

		blogpost.Title = fmt.Sprintf("%v", metadata["Title"])
		if metadata["Summary"] != nil {
			blogpost.Summary = fmt.Sprintf("%v", metadata["Summary"])
		}
		if metadata["Tags"] != nil {
			blogpost.Tags = append(blogpost.Tags, fmt.Sprintf("%v", metadata["Tags"]))
		}
		if metadata["Id"] != nil {
			blogpost.Id, err = strconv.Atoi(fmt.Sprintf("%v", metadata["Id"]))
		}
		// Blogpost content
		blogpost.Content = template.HTML(buf.String()) // fuck go, fuck stackoverflow, fuck gin, fuck goldmark
		// "can i also get fucked?" - Laura
		// ":3" - Ambra
		posts = append(posts, blogpost)
	}

	return posts
}
