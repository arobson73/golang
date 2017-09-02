//scans over the elements of the html and displays
//the frequency of the occurance of elementNodes
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

//!+
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "elem_freq: %v\n", err)
		os.Exit(1)
	}
	ef := make(map[string]int, 0)
	visit(ef, doc)
	for k, v := range ef {
		fmt.Println(k, v)
	}
}

func visit(ef map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		ef[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(ef, c)
	}
}

//!-
