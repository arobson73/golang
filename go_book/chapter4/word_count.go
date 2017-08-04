package main

import (
	"bufio"
	"fmt"
	"os"
)

//./prog < file

func main() {
	wordFreq := make(map[string]int)
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		word := scan.Text()
		wordFreq[word]++

	}
	if scan.Err() != nil {
		fmt.Fprintln(os.Stderr, scan.Err())
		os.Exit(1)
	}
	for word, n := range wordFreq {
		fmt.Printf("%-20s %d\n", word, n)
	}

}
