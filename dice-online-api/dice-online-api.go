package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"

	"dice-online-api/route"
	"dice-online-api/shared/database"
	"dice-online-api/shared/session"

	"dice-online-api/shared/server"
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
	database.ConnectPostgreSQL()

	// start the server
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	log.Fatal(http.ListenAndServe(":"+port, route.LoadRoutes()))
	// server.Run(route.LoadRoutes(), nil, config.Server)
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
