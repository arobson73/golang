package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

func main() {
	data := ` 
<!DOCTYPE html>
<html>
<body>

<p>The image is a link. You can click on it.</p>

<a href="default.asp">
  <img src="smiley.gif" alt="HTML tutorial" style="width:42px;height:42px;border:0">
</a>

<p>We have added "border:0" to prevent IE9 (and earlier) from displaying a border around the image.</p>

</body>
</html>
`
	sr := NewstrReader(data)
	doc, err := html.Parse(sr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "otherlinks: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}

}

type StrReader string

//see io file. Interface Reader implements Read method
// bytes are copied into p. number of bytes copied are return
//note here i just looked into Reader see what that did
func (s *StrReader) Read(p []byte) (n int, err error) {

	n = copy(p, []byte(*s))
	err = io.EOF
	return
}

func NewstrReader(s string) *StrReader {
	var str StrReader
	str = StrReader(s)
	return &str
}
func visit(links []string, n *html.Node) []string {

	if n.Type == html.ElementNode && (n.Data == "a" || n.Data == "img" || n.Data == "script") {
		for _, a := range n.Attr {
			if a.Key == "src" || a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
