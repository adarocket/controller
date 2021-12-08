package save

import (
	"github.com/adarocket/controller/db/structs"
	"github.com/adarocket/controller/repository/informer"
	"github.com/adarocket/proto/proto-gen/cardano"
	"log"
	"time"
)

func AutoSave(server *informer.CardanoInformServer, db structs.Database, minutes int) {
	go func() {
		for _ = range time.Tick(time.Minute * time.Duration(minutes)) {
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
			log.Println("GetStatistic() returned nil")
			continue
		}

		dataNodes := structs.Node{
			NodeAuthData: *value.NodeAuthData,
		}
		if stats.NodeBasicData != nil {
			dataNodes.NodeBasicData = *stats.NodeBasicData
		}

		if err := db.CreateNodeData(dataNodes); err != nil {
			log.Println(err)
			log.Println("lost data", dataNodes)
			if err := db.Ping(); err != nil {
				log.Println(err)
			}
			continue
		}

		dataServer := structs.ServerData{Uuid: value.NodeAuthData.Uuid}
		if stats.Updates != nil {
			dataServer.Updates = *stats.Updates
		}
		if stats.CpuState != nil {
			dataServer.CPUState = *stats.CpuState
		}
		if stats.Online != nil {
			dataServer.Online = *stats.Online
		}
		if stats.MemoryState != nil {
			dataServer.MemoryState = *stats.MemoryState
		}
		if stats.Security != nil {
			dataServer.Security = *stats.Security
		}
		if stats.ServerBasicData != nil {
			dataServer.ServerBasicData = *stats.ServerBasicData
		}

		if err := db.CreateNodeServerData(dataServer); err != nil {
			log.Println(err)
			log.Println("lost data", dataServer)
			if err := db.Ping(); err != nil {
				log.Println(err)
			}
		}

		dataCardano := structs.CardanoData{Uuid: value.NodeAuthData.Uuid}
		if stats.Epoch != nil {
			dataCardano.Epoch = *stats.Epoch
		}
		if stats.KesData != nil {
			dataCardano.KESData = *stats.KesData
		}
		if stats.Blocks != nil {
			dataCardano.Blocks = *stats.Blocks
		}
		if stats.StakeInfo != nil {
			dataCardano.StakeInfo = *stats.StakeInfo
		}
		if stats.NodePerformance != nil {
			dataCardano.NodePerformance = *stats.NodePerformance
		}
		if stats.NodeState != nil {
			dataCardano.NodeState = *stats.NodeState
		}

		if err := db.CreateCardanoData(dataCardano); err != nil {
			log.Println(err)
			log.Println("lost data", dataCardano)
			if err := db.Ping(); err != nil {
				log.Println(err)
			}
		}
	}

	server.NodeStatistics = map[string]*cardano.SaveStatisticRequest{}
}
