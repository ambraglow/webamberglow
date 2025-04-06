package handlers

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

var md goldmark.Markdown

type Blogpost struct {
	Title   string
	Summary string
	Tags    []string
}

func BlogMarkdownInit() {
	md = goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
			extension.Linkify,
		),
	)
}

func BlogPostMetadata() []Blogpost {
	posts, err := filepath.Glob("markdown/*.md")
	if err != nil {
		panic(err.Error())
	}
	var blogposts []Blogpost

	for _, post := range posts {
		var buf bytes.Buffer
		context := parser.NewContext()

		file, err := os.ReadFile(post)
		if err != nil {
			panic(err.Error())
		}

		if err := md.Convert(file, &buf, parser.WithContext(context)); err != nil {
			panic(err.Error())
		}

		metadata := meta.Get(context)

		var blogpost Blogpost

		blogpost.Title = fmt.Sprintf("%v", metadata["Title"])
		if metadata["Summary"] != nil {
			blogpost.Summary = fmt.Sprintf("%v", metadata["Summary"])
		}
		if metadata["Tags"] != nil {
			blogpost.Tags = append(blogpost.Tags, fmt.Sprintf("%v", metadata["Tags"]))
		}

		blogposts = append(blogposts, blogpost)
	}

	return blogposts
}
