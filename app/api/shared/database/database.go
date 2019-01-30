package database

import (
	"database/sql"
	"log"

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
func postgresqldsn(ci PostgreSQLInfo) string {
	// Example: root:@tcp(localhost:3306)/ipaddressservices

	return "user=" + ci.Username + " " +
		"password=" + ci.Password + " " +
		"dbname=" + ci.Name + " " +
		"host=" + ci.Host + " " +
		"port=" + ci.Port + " " +
		//"connection_timeout=" + ci.ConnectTimeout + " " +
		"sslmode=" + ci.SSLMode
}

//ConnectPostgreSQL connects to a given database
func ConnectPostgreSQL(d PostgreSQLInfo) {
	var err error

	log.Println(d)

	//Connect to PostgreSQL
	if POSTGRESQL, err = sql.Open("postgres", postgresqldsn(d)); err != nil {
		log.Println("Postgres SQL Driver Error", err)
	}

	// Check if the db is alive
	if err = POSTGRESQL.Ping(); err != nil {
		log.Println("Database Error", err)
	}
}
