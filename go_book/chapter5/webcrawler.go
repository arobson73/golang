package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/andy/gopl.io/ch5/links"
)

func crawler(urls []string) {

	breadthFirst := func(f func(item string) []string, worklist []string) {
		seen := make(map[string]bool)
		for len(worklist) > 0 {
			items := worklist
			worklist = nil
			for _, item := range items {
				if !seen[item] {
					seen[item] = true
					worklist = append(worklist, f(item)...)
				}
			}
		}
	}

	var origHost string //original host
	saveLink := func(urlin string) error {
		url, err := url.Parse(urlin)
		if err != nil {
			return fmt.Errorf("Problem with url : %s", err)
		}
		if origHost == "" {
			origHost = url.Host
		} else if origHost != url.Host { //reject a new host
			return nil
		}
		fmt.Println("Host:", url.Host)
		fmt.Println("Path:", url.Path)

		dir := url.Host
		dir = filepath.Join(dir, url.Path)
		filename := filepath.Join(dir, "data.html")

		fmt.Println("Filename:", filename)
		err = os.MkdirAll(dir, 0777)
		if err != nil {
			return fmt.Errorf("Problem with mkdir : %s", err)
		}
		resp, err := http.Get(urlin)
		if err != nil {
			return fmt.Errorf("Probem with http get : %s", err)
		}
		defer resp.Body.Close()
		file, err := os.Create(filename)
		if err != nil {
			return fmt.Errorf("Problem with create file : %s", err)
		}

		_, err = io.Copy(file, resp.Body)
		if err != nil {
			return fmt.Errorf("Problem with copy file : %s", err)
		}

		err = file.Close()
		if err != nil {
			return fmt.Errorf("Problem with close file : %s", err)
		}
		return nil

	}
	crawl := func(url string) []string {
		fmt.Println(url)
		saveLink(url)
		list, err := links.Extract(url)
		if err != nil {
			log.Print(err)
		}
		return list
	}

	breadthFirst(crawl, urls)

}

//!+main
func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	crawler(os.Args[1:])
}
