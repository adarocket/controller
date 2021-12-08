package save

import (
	"github.com/adarocket/controller/repository/db/structs"
	"github.com/adarocket/controller/repository/informer"
	"github.com/adarocket/proto/proto-gen/cardano"
	"log"
	"time"
)

const timeout = 5

func AutoSave(server *informer.CardanoInformServer, db structs.Database) {
	go func() {
		for _ = range time.Tick(time.Minute * timeout) {
			saveToDb(server, db)
		}
	}()
}

func saveToDb(server *informer.CardanoInformServer, db structs.Database) {
	for _, value := range server.NodeStatistics {
		if value.NodeAuthData == nil {
			continue
		}

		var stats *cardano.Statistic
		if stats = value.GetStatistic(); stats == nil {
			continue
		}

		if err := db.Ping(); err != nil {
			log.Println(err)
		}

		dataNodes := structs.Node{
			NodeAuthData:  *value.NodeAuthData,
			NodeBasicData: *stats.NodeBasicData,
		}
		if err := db.CreateNodeData(dataNodes); err != nil {
			log.Println(err)
			log.Println("lost data", dataNodes)
			continue
		}

		dataServer := structs.ServerData{
			Uuid:            value.NodeAuthData.Uuid,
			Updates:         *stats.Updates,
			CPUState:        *stats.CpuState,
			Online:          *stats.Online,
			MemoryState:     *stats.MemoryState,
			Security:        *stats.Security,
			ServerBasicData: *stats.ServerBasicData,
		}
		if err := db.CreateNodeServerData(dataServer); err != nil {
			log.Println(err)
			log.Println("lost data", dataServer)
		}

		dataCardano := structs.CardanoData{
			Uuid:            value.NodeAuthData.Uuid,
			Epoch:           *stats.Epoch,
			KESData:         *stats.KesData,
			Blocks:          *stats.Blocks,
			StakeInfo:       *stats.StakeInfo,
			NodePerformance: *stats.NodePerformance,
			NodeState:       *stats.NodeState,
		}
		if err := db.CreateCardanoData(dataCardano); err != nil {
			log.Println(err)
			log.Println("lost data", dataCardano)
		}

		server.NodeStatistics = map[string]*cardano.SaveStatisticRequest{}
	}
}
