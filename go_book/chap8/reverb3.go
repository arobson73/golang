// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.

//use this with the netcat4 code.
//run reverb3 first, then run this.
//while its echoing the data hit, CTRL-D. the echo data will
//continue and then close
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration, wgp *sync.WaitGroup) {
	defer wgp.Done()
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))

}

//!+
func handleConn(c net.Conn) {
	//	fmt.Println("Connection")
	var wg sync.WaitGroup
	input := bufio.NewScanner(c)
	for input.Scan() { // note it loops inside Scan until there are tokens available
		wg.Add(1)
		go echo(c, input.Text(), 1*time.Second, &wg)

	}
	// NOTE: ignoring potential errors from input.Err()
	go func() {
		wg.Wait()
		if tcpconn, ok := c.(*net.TCPConn); ok {
			//fmt.Println("Closing Write")
			tcpconn.CloseWrite()
		}
	}()
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
