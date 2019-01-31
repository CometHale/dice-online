package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
	"runtime"

	"github.com/comethale/dice-online/app/api/route"
	"github.com/comethale/dice-online/app/api/shared/database"
	"github.com/comethale/dice-online/app/api/shared/session"
	"google.golang.org/appengine"

	"github.com/comethale/dice-online/app/api/shared/server"
)

func main() {

	// Logging, verbose with file name and line number
	log.SetFlags(log.Lshortfile)

	// use all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())

	// retrieve and unmarshal the configuration information
	LoadConfig("config"+string(os.PathSeparator)+"config.json", &config)

	// Configure the session cookie store
	session.Configure(config.Session)

	// initiate database connection
	database.ConnectPostgreSQL(config.PostgreSQL)

	// start the server
	server.Run(route.LoadRoutes(), nil, config.Server)

	appengine.Main()
}

// *****************************************************************************
// Application Settings
// *****************************************************************************

// config is the settings variable
var config configuration

// configuration contains the application settings
type configuration struct {
	PostgreSQL database.PostgreSQLInfo `json:"PostgreSQL"`
	Server     server.Server           `json:"Server"`
	Session    session.Session         `json:"Session"`
}

// LoadConfig loads the confiugration file and unmarshals it into the settings variable
func LoadConfig(configPath string, config *configuration) {

	var input = io.ReadCloser(os.Stdin)
	input, err := os.Open(configPath)

	if err != nil {
		log.Fatalln(err)
	}

	jsonBytes, err := ioutil.ReadAll(input)
	input.Close()

	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(jsonBytes, &config)

	if err != nil {
		log.Fatalln(err)
	}
}
