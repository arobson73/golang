package main

import "os"
import "fmt"
import "bufio"
import "log"
import "strings"

func keepLine(line string, fils []string) bool {
	for _, fil := range fils {
		if strings.Contains(line, fil) {
			return true
		}
	}
	return false
}

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Need to supply input file and filter params")
    fmt.Println("go run filter_file.go infile.txt keep INFO WARN")
	}else{
		fn := os.Args[1]
		op := os.Args[2]
		filt := os.Args[3:]
		file,err := os.Open(fn)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			nxtLine := scanner.Text()
			if strings.Compare(op,"keep") == 0 {
				if keepLine(nxtLine,filt) {
					fmt.Println(nxtLine)
				}
			}else {
				if !keepLine(nxtLine,filt) {
					fmt.Println(nxtLine)
				}
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
	

}
	
