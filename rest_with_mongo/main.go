package main

import (
	"flag"
	"fmt"
	"log"
	//will need to edit these paths depending on your path location
	"github.com/rest_with_mongo/myevents/src/eventsservice/rest"
	"github.com/rest_with_mongo/myevents/src/lib/configuration"
	"github.com/rest_with_mongo/myevents/src/lib/persistence/dblayer"
)

func main() {

	confPath := flag.String("conf", `.\configuration\config.json`, "flag to set the path to the configuration json file")
	flag.Parse()
	//extract configuration
	config, _ := configuration.ExtractConfiguration(*confPath)

	fmt.Println("Connecting to database")
	dbhandler, _ := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)
	//RESTful API start
	fmt.Println("Here")
	log.Fatal(rest.ServeAPI(config.RestfulEndpoint, dbhandler))
}
