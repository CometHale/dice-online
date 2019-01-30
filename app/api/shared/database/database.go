package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

//POSTGRESQL is a global variable holding the database for the currently running app instance
var POSTGRESQL *sql.DB

// PostgreSQLInfo is a struct type holding conneciton information about a PostgreSQL database
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
		"sslmode=" + ci.SSLMode
}

// LoadDatabase connects to the database described by the given PostgreSQLInfo
func LoadDatabase(db PostgreSQLInfo) {

	log.Println(db)

	POSTGRESQL, err := sql.Open("postgres", postgresqldsn(db))

	if err != nil {
		log.Println("Postgres SQL Driver Error", err)
	}

	// Check if the db is alive
	if err = POSTGRESQL.Ping(); err != nil {
		log.Println("Database Error", err)
	}
}
