//call like this. make sure clocks are running first!
///clockwall US=localhost:8010 London=localhost:8030 Tokyo=localhost:8020
package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type placeHost struct {
	place []string
	hosts []string
}

//just simple parsing, no checking for bad input
func getParams() placeHost {
	var ph placeHost
	for i := 1; i < len(os.Args); i++ {
		temp := strings.Split(os.Args[i], "=")
		if len(temp) == 2 {
			ph.place = append(ph.place, temp[0])
			ph.hosts = append(ph.hosts, temp[1])
		}
	}
	return ph
}
func showInput(in placeHost) {
	for i := 0; i < len(in.place); i++ {
		fmt.Println(in.place[i], in.hosts[i])
	}

}
func getConns(hosts []string) []net.Conn {
	var c []net.Conn
	for _, h := range hosts {
		fmt.Println("Dialling ", h)
		conn, err := net.Dial("tcp", h)
		if err != nil {
			fmt.Println("Error Dialling", h)
			log.Fatal(err)
		}
		c = append(c, conn)
	}
	return c
}

func main() {
	ph := getParams()
	showInput(ph)
	cons := getConns(ph.hosts)
	for {
		b := make([]byte, 1024)
		for i, c := range cons {
			n, err := c.Read(b)
			//remove the newline char from the Read
			if bytes.Contains(b, []byte("\n")) {
				b = bytes.Replace(b, []byte("\n"), []byte(""), 1)

			}
			if err != nil {
				fmt.Println("Error Reading connection")
				log.Fatal(err)
				return
			}
			if n > 0 {
				os.Stdout.WriteString("\r")
				if i == 0 {
					os.Stdout.WriteString(ph.place[0] + "=")
				}
				for j := 0; j < i; j++ {

					os.Stdout.WriteString("\t\t\t")
				}
				if i != 0 {
					os.Stdout.WriteString(ph.place[i] + "=")

				}
			}
			os.Stdout.Write(b)
		}
		time.Sleep(1 * time.Second)
	}

}

//!-
