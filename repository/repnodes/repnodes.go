package repnodes

import (
	"github.com/adarocket/controller/db/postgresql"
	"github.com/adarocket/controller/db/structs"
	"github.com/adarocket/controller/repository/config"
	"log"
)

type RepoNodes struct {
	db structs.Database
}

func InitController(conf config.Config) RepoNodes {
	dbPostgres, err := postgresql.InitDatabase(conf)
	if err != nil {
		log.Fatal(err)
	}

	return RepoNodes{db: dbPostgres}
}

func (c *RepoNodes) GetNodeData(uuid string) (structs.Node, error) {
	return c.db.GetNodeData(uuid)
}
