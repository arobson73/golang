package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(1)
	}
	words, images := countWordsAndImages(doc)
	fmt.Println(words, images)
}

func countWordsAndImages(n *html.Node) (words, images int) {

	type counter struct {
		images int
		words  int
	}

	var f func(ctr *counter, n *html.Node)
	f = func(ctr *counter, n *html.Node) {
		if n == nil {
			return
		}
		switch n.Type {
		case html.TextNode:
			ctr.words += wordCount(n.Data)
		case html.ElementNode:
			if n.Data == "img" {
				ctr.images++
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(ctr, c)
		}
	}

	var mycounter counter
	f(&mycounter, n)
	return mycounter.words, mycounter.images
}

func wordCount(s string) int {
	n := 0
	scan := bufio.NewScanner(strings.NewReader(s))
	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		n++
	}
	return n
}
