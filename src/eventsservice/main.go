package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/souravsk/myevent/src/eventsservice/rest"
	"github.com/souravsk/myevent/src/lib/configuration"
	"github.com/souravsk/myevent/src/lib/persistence/dblayer"
)

func main() {
	confPath := flag.String("conf", `.\configuration\config.json`, "flag to set the path to the configuration json file") // flag to set the path to the configuration json file
	flag.Parse()
	//extract configuration
	config, _ := configuration.ExtractConfiguration(*confPath) // extract configuration

	fmt.Println("Connecting to database")
	dbhandler, _ := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection) // connect to database
	//RESTful API start
	log.Fatal(rest.ServeAPI(config.RestfulEndpoint, dbhandler)) // start the RESTful API
}
