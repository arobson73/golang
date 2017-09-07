package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

//run like this:
//progname htmlfile id

//note i used an html file as input instead of fetching html from a url
//my computer has intermittent internet issues. pain in neck
func main() {
	if len(os.Args[1:]) != 2 {
		fmt.Println("Must enter html file and element to search")
		os.Exit(1)
	}
	for _, arg := range os.Args[1:] {
		fmt.Println("Enter arg: ", arg)
	}
	var r io.Reader
	var err error
	htmlin := os.Args[1]
	id := os.Args[2]
	r, err = os.Open(htmlin)
	if err != nil {
		fmt.Println("Problem reading ", htmlin)
		os.Exit(1)
	}
	doc, err := html.Parse(r)
	if err != nil {
		fmt.Println("Problems parsing html")
		os.Exit(1)
	}
	e := findElement(doc, id)
	if e != nil {
		fmt.Println("Found element: %+v", e)
	} else {
		fmt.Println("Not found element with id", id)
	}

}

func findElement(doc *html.Node, id string) *html.Node {
	// forEachNode calls the functions pre(x) and post(x) for each node
	// x in the tree rooted at n. Both functions are optional.
	// pre is called before the children are visited (preorder) and
	// post is called after (postorder).

	var forEachNode func(n *html.Node, pre, post func(n *html.Node) bool)
	forEachNode = func(n *html.Node, pre, post func(n *html.Node) bool) {
		if pre != nil {
			if pre(n) {
				return
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			forEachNode(c, pre, post)
		}
		if post != nil {
			post(n)
		}
	}

	var ele *html.Node
	var startElement func(n *html.Node) bool
	startElement = func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == id {
					ele = n
					return true
				}
			}
		}
		return false
	}

	forEachNode(doc, startElement, startElement) // last param can be nil

	return ele
}
