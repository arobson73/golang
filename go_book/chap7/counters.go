package main //package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

type WordCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
	r := bytes.NewReader(p)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*w += WordCounter(1)
	}
	err := scanner.Err()
	if err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	return len(p), err
}

type LineCounter int

func (lc *LineCounter) Write(p []byte) (int, error) {
	lines := bytes.Split(p, []byte("\n"))
	*lc += LineCounter(len(lines))
	return len(p), nil
}

type ByteCounter struct {
	count int64
	w     io.Writer
}

func (bc *ByteCounter) Write(p []byte) (int, error) {
	n, err := bc.w.Write(p)
	bc.count += int64(n)
	return n, err
}

//Writer is an interface to the Write method, so we need
// to return a pointer to a ByteCounter which is the interface for Write
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	newbc := ByteCounter{0, w}
	return &newbc, &newbc.count
}

func main() {

	var c WordCounter
	r, _ := c.Write([]byte("one two three four five size seven eight"))
	fmt.Println(c, r)
	r, _ = c.Write([]byte("one two three four"))
	fmt.Println(c, r)
	//	c.words = 0 // reset
	c = 0
	var names = "Hi Ho jo done what why when now today then them me my we when ten gen ken"
	fmt.Fprintf(&c, "fee fi %s", names)
	fmt.Println(c)

	var lc LineCounter
	r, _ = lc.Write([]byte("hello\ngood\nbye\nfast\nslow\nquick"))
	fmt.Println(lc, r)
	wo, count := CountingWriter(os.Stdout)
	fmt.Fprintf(wo, "Hello Golang\n")
	fmt.Println(*count)

}
