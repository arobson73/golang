//crawl with a depth limit
package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
)

type linkDepth struct {
	depth int
	url   string
}

func crawl(url linkDepth) []linkDepth {
	fmt.Println(url)
	list, err := links.Extract(url.url)
	if err != nil {
		log.Print(err)
	}
	return convert2LinkDepth(list, url.depth)
}

func convert2LinkDepth(args []string, inc int) []linkDepth {
	var lda []linkDepth
	for _, v := range args {
		lda = append(lda, linkDepth{1 + inc, v})
	}
	return lda
}

//!+
func main() {
	worklist := make(chan []linkDepth)  // lists of URLs, may have duplicates
	unseenLinks := make(chan linkDepth) // de-duplicated URLs

	ld := convert2LinkDepth(os.Args[1:], 0)

	// Add command-line arguments to worklist.
	go func() { worklist <- ld }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link.url] && link.depth < 4 {
				seen[link.url] = true
				unseenLinks <- link
			}
		}
	}
}

//!-
