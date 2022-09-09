package main

import (
	"bytes"
	"fmt"

	"github.com/worzeel/markdown2json/json"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/util"
)

func main() {
	fmt.Println("markdown2json")

	md := goldmark.New(
		goldmark.WithRenderer(renderer.NewRenderer(renderer.WithNodeRenderers(util.Prioritized(json.NewRenderer(), 1000)))),
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
	)

	var buf bytes.Buffer
	if err := md.Convert([]byte("# Testing\nThis is a testing"), &buf); err != nil {
		panic(err)
	}

	fmt.Println(buf.String())
}
