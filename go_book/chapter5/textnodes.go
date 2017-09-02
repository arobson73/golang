package main

import (
	"bytes"
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(1)
	}
	text := getText(doc)
	fmt.Println(text)
}

func getText(n *html.Node) string {
	var f func(text *bytes.Buffer, n *html.Node)
	f = func(text *bytes.Buffer, n *html.Node) {
		if n == nil {
			return
		}
		if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
			return
		}

		if n.Type == html.TextNode {
			text.WriteString(n.Data)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(text, c)
		}
	}
	text := new(bytes.Buffer)
	f(text, n)
	return text.String()
}
