package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

//POSTGRESQL is a global variable holding the database for the currently running app instance
var POSTGRESQL *sql.DB

//Database organizes information about a given database
type Database struct {
	Type     string
	Postgres PostgreSQLInfo
}

//PostgreSQLInfo collects all the data necessary to connect to a database
type PostgreSQLInfo struct {
	Username string
	Password string
	Name     string
	Host     string
	Port     string
	SSLMode  string
}

// DSN returns the Data Source Name
func postgresqldsn() string {
	// Example: root:@tcp(localhost:3306)/ipaddressservices

	return "user=" + os.Getenv("Username") + " " +
		"password=" + os.Getenv("Password") + " " +
		"dbname=" + os.Getenv("Name") + " " +
		"host=" + os.Getenv("Host") + " " +
		"port=" + os.Getenv("Port") + " " +
		//"connection_timeout=" + ci.ConnectTimeout + " " +
		"sslmode=" + os.Getenv("SSLMode")
}

//ConnectPostgreSQL connects to a given database
func ConnectPostgreSQL() {
	var err error

	// log.Println(d)

	//Connect to PostgreSQL
	if POSTGRESQL, err = sql.Open("postgres", postgresqldsn()); err != nil {
		log.Println("Postgres SQL Driver Error", err)
	}

	// Check if the db is alive
	if err = POSTGRESQL.Ping(); err != nil {
		log.Println("Database Error", err)
	}
}
