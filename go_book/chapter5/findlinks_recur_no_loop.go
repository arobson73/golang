// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 122.
//!+main

// Findlinks1 prints the links in an HTML document read from standard input. 
//REmoved the loop and used recursion instead of loop
package main

import (
	"fmt"
	"os"
	"strconv"

	"golang.org/x/net/html"
)

//var nest int

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	sta, no := visit(nil, doc, 0)
	fmt.Println("no is ", no)
	for _, link := range sta {
		fmt.Println(link)
	}
}

//!-main

//!+visit
// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node, nin int) (res []string, nout int) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				ns := strconv.Itoa(nin)
				links = append(links, a.Val+".Nest="+ns)
			}
		}
	}
	if n.FirstChild != nil {
		nin++
		links, _ = visit(links, n.FirstChild, nin)
		nin--

	}
	if n.NextSibling != nil {
		nin++
		links, _ = visit(links, n.NextSibling, nin)
		nin--
	}
	return links, nin
}

//!-visit

/*
//!+html
package html

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

func Parse(r io.Reader) (*Node, error)
//!-html
*/
