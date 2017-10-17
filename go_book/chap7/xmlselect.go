// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 214.
//!+

// Xmlselect prints the text of selected elements of an XML document.
//./fetch http://www.w3.org/TR/2006/REC-xml11-20060816 | ./xmlselect div h3
package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	//fmt.Println(len(os.Args[1:]))
	var stack [][]string // stack of element names and attributes
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			newE := make([]string, 0)
			newE = append(newE, tok.Name.Local) // push
			for _, v := range tok.Attr {
				if v.Name.Local == "id" || v.Name.Local == "class" {
					newE = append(newE, v.Value)
				}
			}
			stack = append(stack, newE)
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop

		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				for _, v := range stack {
					fmt.Printf("%s ", strings.Join(v, " | "))
				}
				fmt.Printf("tok:%s\n", tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x [][]string, y []string) bool {
	for len(y) <= len(x) { //len(x) will be number of rows
		if len(y) == 0 {
			return true
		}
		for _, v := range x[0] {
			if v == y[0] {
				y = y[1:]
				break
			}
		}
		x = x[1:] //move to the next row
	}
	return false
}

//!-
