package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

//TODO use a struct to store string and depth. then string
//can be stored without any leading whitespace (simpler to parse)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}
	// forEachNode calls the functions pre(x) and post(x) for each node
	// x in the tree rooted at n. Both functions are optional.
	// pre is called before the children are visited (preorder) and
	// post is called after (postorder).

	var forEachNode func(n *html.Node, pre, post func(n *html.Node))
	forEachNode = func(n *html.Node, pre, post func(n *html.Node)) {
		if pre != nil {
			pre(n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			forEachNode(c, pre, post)
		}
		if post != nil {
			post(n)
		}
	}

	var depth int
	var same []string

	var startElement func(n *html.Node)
	startElement = func(n *html.Node) {
		switch t := n.Type; t {
		case html.ElementNode:
			temp := fmt.Sprintf("%*s<%s>\n", depth*2, "", n.Data)
			same = append(same, temp)
			depth++
		case html.TextNode:
			if len(n.Data) > 0 {
				temp := fmt.Sprintf("%*s%s\n", depth*2, "", n.Data)
				same = append(same, temp)

			}
		case html.CommentNode:
			temp := fmt.Sprintf("<!--%s-->\n", n.Data)
			same = append(same, temp)

		}
	}

	var endElement func(n *html.Node)
	endElement = func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth--
			temp := fmt.Sprintf("%*s</%s>\n", depth*2, "", n.Data)
			same = append(same, temp)
			//		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
	var useShortForms func(in []string) []string
	useShortForms = func(in []string) []string {
		//check for match of <img></img> and replace with <img/>
		var result []string
		for i, j := 0, 1; j < len(in); {
			str := in[i]
			idx1 := strings.Index(str, "<")
			idx2 := strings.Index(str, ">")
			if idx1 == -1 || idx2 == -1 {
				i += 1
				j = i + 1
				continue
			}
			sl := str[idx1+1 : idx2] //get value between <>
			slc := "</" + sl + ">"   // see if next one is this, if so replace
			if strings.Contains(in[j], slc) {
				//	fmt.Printf("yes in[j] =%s, slc= %s\n", in[j], slc)
				newstr := str[:idx1+1] + sl + "/>" + "\n"
				result = append(result, newstr)
				i += 2
			} else {
				//fmt.Printf("no in[j] =%s, slc= %s\n", in[j], slc)
				result = append(result, str)
				i += 1
			}
			j = i + 1

		}
		return result

	}
	forEachNode(doc, startElement, endElement)
	fmt.Println("\ntest results")
	res := useShortForms(same)
	for _, val := range res {
		fmt.Printf(val)
	}

	return nil
}
