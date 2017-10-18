package main

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

//./fetch http://www.w3.org/TR/2006/REC-xml11-20060816 | ./xmlTree
type Node interface { //can be a *Element or CharData type
	String() string
}
type CharData string

func (c CharData) String() string {
	return string(c)
}

func (e *Element) String() string {
	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	var visit func(n Node, w io.Writer, depth int)
	visit = func(n Node, w io.Writer, depth int) {
		switch n := n.(type) {
		case *Element:
			fmt.Fprintf(w, "%*s%s %s\n", depth*2, "", n.Type.Local, n.Attr)
			for _, ne := range n.Children {
				visit(ne, w, depth+1)
			}
		case CharData:
			fmt.Fprintf(w, "%*s%q\n", depth*2, "", n)
		default:
			panic(fmt.Sprintf("recieved %T", n))

		}

	}
	visit(e, writer, 0)
	return b.String()

}

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func createTree(r io.Reader) (Node, error) {
	dec := xml.NewDecoder(r)
	var stack []*Element
	var root Node
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			//fmt.Print("S")
			ele := &Element{tok.Name, tok.Attr, nil}
			if len(stack) != 0 {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, ele)
			} else {
				root = ele
			}
			stack = append(stack, ele) // push
		case xml.EndElement:
			//fmt.Print("E")
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			//fmt.Print("C")
			if len(stack) != 0 {
				parent := stack[len(stack)-1]
				parent.Children = append(parent.Children, CharData(tok))
			} else {
				//	fmt.Print(">*H*<")
			}

		}
	}
	return root, nil
}

func main() {
	node, err := createTree(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
	fmt.Println(node.String())

}
