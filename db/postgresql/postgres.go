package postgresql

import (
	"database/sql"
	"fmt"
	"github.com/adarocket/controller/repository/config"
	_ "github.com/lib/pq"
)

func InitDatabase(config config.Config) (Postgresql, error) {
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
