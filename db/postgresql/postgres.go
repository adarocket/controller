package postgresql

import (
	"adarocket/controller/config"
	"database/sql"
	"fmt"
	"log"
)

// PostgreSQL ...
var Postg postgresql

// InitDatabase ...
func InitDatabase(config config.DBConfig) {
	/*connStr := fmt.Sprintf(`user=%s password=%s dbname=%s sslmode=%s`,
	"postgres", "postgresql", "postgres", "disable")*/
	connStr := fmt.Sprintf(`user=%s password=%s dbname=%s sslmode=%s`,
		config.User, config.Password, config.Dbname, config.Sslmode)
	// connStr := "user = postgres password=postgresql dbname=crypto sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	Postg.dbConn = db
}

type postgresql struct {
	dbConn *sql.DB
}
