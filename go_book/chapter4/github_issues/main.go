// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 112.
//!+

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/andy/gopl.io/ch4/github"
)

func printResults(cont []*github.Issue, msg string) {
	if len(cont) > 0 {
		fmt.Printf("\n%s\n", msg)
		for _, item := range cont {
			fmt.Printf("#%-5d %9.9s %.45s %.55s\n",
				item.Number, item.User.Login, item.CreatedAt, item.Title)

		}
	}

}

//"gopl.io/ch4/github"
//!+
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	now := time.Now()
	dayOld := make([]*github.Issue, 0)
	monthOld := make([]*github.Issue, 0)
	yearOld := make([]*github.Issue, 0)
	older := make([]*github.Issue, 0)

	day := now.AddDate(0, 0, -1)
	month := now.AddDate(0, -1, 0)
	year := now.AddDate(-1, 0, 0)

	for _, item := range result.Items {
		switch {
		case item.CreatedAt.After(day):
			dayOld = append(dayOld, item)
		case item.CreatedAt.After(month):
			monthOld = append(monthOld, item)
		case item.CreatedAt.After(year):
			yearOld = append(yearOld, item)
		default:
			older = append(older, item)

		}
	}

	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.45s %.55s\n",
			item.Number, item.User.Login, item.CreatedAt, item.Title)
	}
	printResults(dayOld, "less than a day")
	printResults(monthOld, "less than a month")
	printResults(yearOld, "less than a year")
	printResults(older, "older than a year")
}

//!-

/*
//!+textoutput
$ go build gopl.io/ch4/issues
$ ./issues repo:golang/go is:open json decoder
13 issues:
#5680    eaigner encoding/json: set key converter on en/decoder
#6050  gopherbot encoding/json: provide tokenizer
#8658  gopherbot encoding/json: use bufio
#8462  kortschak encoding/json: UnmarshalText confuses json.Unmarshal
#5901        rsc encoding/json: allow override type marshaling
#9812  klauspost encoding/json: string tag not symmetric
#7872  extempora encoding/json: Encoder internally buffers full output
#9650    cespare encoding/json: Decoding gives errPhase when unmarshalin
#6716  gopherbot encoding/json: include field name in unmarshal error me
#6901  lukescott encoding/json, encoding/xml: option to treat unknown fi
#6384    joeshaw encoding/json: encode precise floating point integers u
#6647    btracey x/tools/cmd/godoc: display type kind of each named type
#4237  gjemiller encoding/base64: URLEncoding padding is optional
//!-textoutput
*/
