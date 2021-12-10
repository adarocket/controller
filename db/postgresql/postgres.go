package postgresql

import (
	"database/sql"
	"fmt"
	"github.com/adarocket/controller/db/structs"
	"github.com/adarocket/controller/repository/config"
	_ "github.com/lib/pq"
	"log"
)

var dbInstance Postgresql

func InitDatabase(config config.Config) (structs.Database, error) {
	var err error

	if dbInstance.dbConn == nil {
		dbInstance, err = createDbInstance(config)
		if err != nil {
			log.Println(err)
		}
	} else if err = dbInstance.Ping(); err != nil {
		log.Println("Db ping returned error")
		dbInstance, err = createDbInstance(config)
		if err != nil {
			log.Println(err)
		}
	}

	return dbInstance, err
}

func createDbInstance(config config.Config) (Postgresql, error) {
	log.Println("Starting initialization db...")
	connStr := fmt.Sprintf(`user=%s password=%s dbname=%s sslmode=%s`,
		config.DBConfig.User, config.DBConfig.Password,
		config.DBConfig.Dbname, config.DBConfig.Sslmode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return Postgresql{}, err
	}

	if err = db.Ping(); err != nil {
		return Postgresql{}, err
	}

	return Postgresql{dbConn: db}, nil
}

type Postgresql struct {
	dbConn *sql.DB
}

func (p Postgresql) Ping() error {
	return p.dbConn.Ping()
}
