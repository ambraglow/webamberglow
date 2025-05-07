package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strconv"

	fences "github.com/stefanfritsch/goldmark-fences"
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
	// initialize goldmark context
	MD = goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
			extension.Linkify,
			&fences.Extender{},
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithXHTML(),
		),
	)
	// returns a struct containing the blog posts
	Posts, _ = BlogPosts()
}

func BlogPosts() ([]Blogpost, error) {
	files, err := filepath.Glob("markdown/*.md")
	if err != nil {
		return nil, fmt.Errorf("Path not found: %w", err)
	}
	var posts []Blogpost

	for _, file := range files {
		var buf bytes.Buffer
		var blogpost Blogpost

		context := parser.NewContext()

		fileread, err := os.ReadFile(file)
		if err != nil {
			return nil, fmt.Errorf("couldn't read blog post file: %w", err)
		}

		if err := MD.Convert(fileread, &buf, parser.WithContext(context)); err != nil {
			panic(err.Error())
		}

		//	Fill out the header information
		metadata := meta.Get(context)

		blogpost.Title = fmt.Sprintf("%v", metadata["Title"])

		if metadata["Summary"] != nil {
			blogpost.Summary = fmt.Sprintf("%v", metadata["Summary"])
		}

		if metadata["Tags"] != nil {
			blogpost.Tags = append(blogpost.Tags, fmt.Sprintf("%v", metadata["Tags"]))
		}

		if metadata["Id"] != nil {
			blogpost.Id, _ = strconv.Atoi(fmt.Sprintf("%v", metadata["Id"]))
			fmt.Print("post Id logged:")
			fmt.Println(blogpost.Id)
		}

		// Insert the blog post post content
		blogpost.Content = template.HTML(buf.String()) // fuck go, fuck stackoverflow, fuck gin, fuck goldmark
		// "can i also get fucked?" - Laura
		// ":3" - Ambra
		posts = append(posts, blogpost)
	}

	return posts, nil
}
