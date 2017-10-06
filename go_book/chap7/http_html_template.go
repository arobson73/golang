package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type dollars float64

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

var readTemp = template.Must(template.New("read").Parse(`
<html>
<body>
{{range $key, $val := .ItemMap }}
<p>{{$key}}: ${{$val}}</p>
{{ end }}
</body>
</html>
`))

//!+main

type database map[string]dollars
type TemplateList struct {
	ItemMap database
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/read", http.HandlerFunc(db.read))
	mux.Handle("/price", http.HandlerFunc(db.price))
	mux.Handle("/create", http.HandlerFunc(db.create))
	mux.Handle("/update", http.HandlerFunc(db.update))
	mux.Handle("/delete", http.HandlerFunc(db.delete))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

//in browser do localhost:8000/read
func (db database) read(w http.ResponseWriter, req *http.Request) {
	dbsync.Lock()
	if err := readTemp.Execute(w, &TemplateList{db}); err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprintf(w, "Failed to show table:%q\n", err)
	}
	/*
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}*/
	dbsync.Unlock()
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	//check if item already exists
	dbsync.Lock()
	price, ok := db[item]
	dbsync.Unlock()
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

//global
var dbsync sync.Mutex

//notice the quotes which are required since maybe shell thinks & is a command or something
//curl -X GET http://0.0.0.0:8000/update?item=shoes"&"price=7
func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	priceS := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(priceS, 64)
	//check if item already exists
	dbsync.Lock()
	if _, ok := db[item]; ok {
		w.WriteHeader(http.StatusAlreadyReported)
		fmt.Fprintf(w, "%s is already in the database\n", item)

	} else if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Price is not float %s\n", priceS)
	} else {
		db[item] = dollars(price)
		fmt.Fprintf(w, "item %s with price %g\n", item, price)
	}
	dbsync.Unlock()
}
func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	priceS := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(priceS, 64)
	//check if item already exists
	dbsync.Lock()
	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusAlreadyReported)
		fmt.Fprintf(w, "%s is not in the database\n", item)

	} else if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Price is not float %s\n", priceS)
	} else {
		db[item] = dollars(price)
		fmt.Fprintf(w, "item %s with price %g\n", item, price)
	}
	dbsync.Unlock()
}

func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	//check if item already exists
	dbsync.Lock()
	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusAlreadyReported)
		fmt.Fprintf(w, "%s is not in the database\n", item)
	} else {
		delete(db, item)
		fmt.Fprintf(w, "item %s has been deleted \n", item)
	}
	dbsync.Unlock()
}
