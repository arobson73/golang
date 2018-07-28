package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func init() {
	//	log.SetOutput(ioutil.Discard)

}

func main() {
	var buf bytes.Buffer
	var logger = log.New(&buf, "LoggerBuf: ", log.Lshortfile)
	//logger.SetOutput(ioutil.Discard)

	f, err := os.OpenFile("testlogfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	var flogger = log.New(f, "LoggerFile: ", log.Lshortfile)

	defer f.Close()

	logger.Printf("hello")
	logger.Println("123")
	logger.Print("abc")
	fmt.Print(&buf)
	log.Print("what")

	flogger.Print("Hello file logging")

}
