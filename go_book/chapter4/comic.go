package main

import (
	"bufio"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type WordIndex map[string]map[int]bool
type NumIndex map[int]Comic

type Comic struct {
	Alt        string `json:"alt"`
	Day        string `json:"day"`
	Img        string `json:"img"`
	Link       string `json:"link"`
	Month      string `json:"month"`
	News       string `json:"news"`
	Num        int    `json:"num"`
	Safe_Title string `json:"safe_title"`
	Title      string `json:"title"`
	Transcript string `json:"transcript"`
	Year       string `json:"year"`
}

const MAX_COMICS = 20

func getComic(n int) (Comic, error) {
	var comic Comic
	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", n)
	resp, err := http.Get(url)
	if err != nil {
		return comic, fmt.Errorf("failed to get url")

	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return comic, fmt.Errorf("status failed with %s\n", resp.StatusCode)
	}

	if err = json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return comic, fmt.Errorf("failed to decode to json")
	}

	return comic, nil

}
func comicCount() (int, error) {
	resp, err := http.Get("https://xkcd.com/info.0.json")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("can't get main page: %s", resp.Status)
	}
	var comic Comic
	if err = json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return 0, err
	}
	return comic.Num, nil
}

func indexComics(comics chan Comic) (WordIndex, NumIndex) {
	wordIndex := make(WordIndex)
	numIndex := make(NumIndex)
	for comic := range comics {
		numIndex[comic.Num] = comic
		scanner := bufio.NewScanner(strings.NewReader(comic.Transcript))
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			token := strings.ToLower(scanner.Text())
			if _, ok := wordIndex[token]; !ok {
				wordIndex[token] = make(map[int]bool, 1)
			}
			wordIndex[token][comic.Num] = true
		}
	}
	return wordIndex, numIndex
}

func index(filename string) error {
	comicChan, err := fetchComics(MAX_COMICS)
	if err != nil {
		return err
	}
	wordIndex, numIndex := indexComics(comicChan)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	enc := gob.NewEncoder(file)
	fmt.Println(wordIndex)
	err = enc.Encode(wordIndex)
	if err != nil {
		return err
	}
	err = enc.Encode(numIndex)
	if err != nil {
		return err
	}
	fmt.Println(numIndex)
	return nil
}
func fetchComics(amount int) (chan Comic, error) {
	if amount > MAX_COMICS {
		amount = MAX_COMICS
	}

	comicC := make(chan Comic, MAX_COMICS)
	var erre error
	var count int
	for i := 1; i <= amount; i++ {
		comic, err := getComic(i)
		if err != nil {
			fmt.Printf("Can't get comic %d", i)
			erre = err
			continue
		}
		count++
		comicC <- comic
	}
	close(comicC)
	if count != amount {
		return comicC, fmt.Errorf("Could not get the requested amount of comics, only %d of %d", count, amount)
	} else {

		return comicC, erre
	}

}

func readIndex(filename string) (WordIndex, NumIndex, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	dec := gob.NewDecoder(file)
	var wordIndex WordIndex
	var numIndex NumIndex
	err = dec.Decode(&wordIndex)
	if err != nil {
		return nil, nil, err
	}
	dec.Decode(&numIndex)
	if err != nil {
		return nil, nil, err
	}
	return wordIndex, numIndex, nil
}

func showWordIndex(filename string) {
	wi, _, err := readIndex(filename)
	if err != nil {
		fmt.Printf("error is %s", err)
	}
	for k, v := range wi {
		fmt.Printf("1:key is %t, val is %t\n", k, v)
		fmt.Printf("2:key is %+v, val is %+v\n", k, v)
		fmt.Printf("3:val is %+v\n", wi[k])

	}
}
func comicsContainingWords(words []string, wordIndex WordIndex, numIndex NumIndex) []Comic {
	found := make(map[int]int) // comic Num -> count words found
	comics := make([]Comic, 0)
	for _, word := range words {
		for num := range wordIndex[word] {
			found[num]++
		}
	}
	for num, nfound := range found {
		if nfound == len(words) {
			comics = append(comics, numIndex[num])
		}
	}
	return comics
}

func search(query string, filename string) {
	wordIndex, numIndex, err := readIndex(filename)
	if err != nil {
		fmt.Printf("error occured searching %s", err)
		return
	}
	comics := comicsContainingWords(strings.Fields(query), wordIndex, numIndex)
	if len(comics) == 0 {
		fmt.Printf("No comics found with those words\n")
		return
	}
	for _, comic := range comics {
		fmt.Printf("%+v\n\n", comic)
	}

}

func showWordIndexForString(filename, str string) {
	wi, _, err := readIndex(filename)
	if err != nil {
		fmt.Printf("Error occured in showWordIndexForString")
		return
	}
	for k, v := range wi[str] {
		fmt.Printf("key is %+v, val is %+v\n", k, v)
	}

}

func printAll(in <-chan Comic) {
	for c := range in {
		fmt.Println(c.Title)
	}
}

const usage = "input must be as follows\ncomic get N\ncomic count\ncomic printAll\ncomic index filename\ncomic showWordIndex filename\ncomic showWordIndexForString filename string\ncomic search filename string"

func usageMessage() {
	fmt.Println(usage)
	os.Exit(1)
}

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}

	cmd := os.Args[1]
	switch cmd {
	case "get":
		if len(os.Args) != 3 {
			usageMessage()
		}
		n, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "N %s must be convertable to integer", os.Args[1])
		}
		comic, err := getComic(n)
		if err != nil {
			fmt.Printf("error reading comic %s", err)
		}
		fmt.Printf("Title is %s\n", comic.Title)
	case "count":
		ccount, err := comicCount()
		if err != nil {
			fmt.Printf("error reading total comic count %s", err)
		}
		fmt.Printf("Total comic count is %d\n", ccount)
	case "printAll":
		cc, _ := fetchComics(MAX_COMICS)
		printAll(cc)
	case "index":
		if len(os.Args) != 3 {
			usageMessage()
		}
		err := index(os.Args[2])
		if err != nil {
			log.Fatal("Error serializing indexes", err)
		}
	case "showWordIndex":
		if len(os.Args) != 3 {
			usageMessage()
		}
		showWordIndex(os.Args[2])
	case "showWordIndexForString":
		if len(os.Args) != 4 {
			usageMessage()

		}
		showWordIndexForString(os.Args[2], os.Args[3])

	case "search":
		if len(os.Args) != 4 {
			usageMessage()
		}
		search(os.Args[3], os.Args[2])

	default:
		fmt.Println("Unrecognised input command")

	}

	os.Exit(0)

}
