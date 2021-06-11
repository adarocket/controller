package save

import (
	"adarocket/controller/db/structs"
	"adarocket/controller/informer"
	"log"
	"time"
)

const timeout = 10

var db structs.Database

func InitDb(newDb structs.Database) {
	db = newDb
}

func AutoSave(server *informer.InformServer) {
	go func() {
		for _ = range time.Tick(time.Minute * timeout) {
			saveToDb(server)
		}
	}()
}

func saveToDb(server *informer.InformServer) {
	for _, value := range server.NodeStatistics {
		if value.NodeInfo == nil {
			continue
		}
		if value.NodeInfo.NodeAuthData != nil {
			if err := db.CreateNodeAuthData(*value.NodeInfo.NodeAuthData); err != nil {
				log.Println(err)
			}
		}
		if value.NodeInfo.Statistic == nil {
			continue
		}
		if value.NodeInfo.Statistic.ServerBasicData != nil {
			if err := db.CreateServerBasicData(*value.NodeInfo.Statistic.ServerBasicData,
				value.NodeInfo.NodeAuthData.Uuid); err != nil {
				log.Println(err)
			}
		}
		if value.NodeInfo.Statistic.NodeState != nil {
			if err := db.CreateNodeStateData(*value.NodeInfo.Statistic.NodeState,
				value.NodeInfo.NodeAuthData.Uuid); err != nil {
				log.Println(err)
			}
		}
		if value.NodeInfo.Statistic.NodePerformance != nil {
			if err := db.CreateNodePerformanceData(*value.NodeInfo.Statistic.NodePerformance,
				value.NodeInfo.NodeAuthData.Uuid); err != nil {
				log.Println(err)
			}
		}
		if value.NodeInfo.Statistic.MemoryState != nil {
			if err := db.CreateMemoryStateData(*value.NodeInfo.Statistic.MemoryState,
				value.NodeInfo.NodeAuthData.Uuid); err != nil {
				log.Println(err)
			}
		}
		if value.NodeInfo.Statistic.Online != nil {
			if err := db.CreateOnlineData(*value.NodeInfo.Statistic.Online,
				value.NodeInfo.NodeAuthData.Uuid); err != nil {
				log.Println(err)
			}
		}
		if value.NodeInfo.Statistic.StakeInfo != nil {
			if err := db.CreateStakeInfoData(*value.NodeInfo.Statistic.StakeInfo,
				value.NodeInfo.NodeAuthData.Uuid); err != nil {
				log.Println(err)
			}
		}
		if value.NodeInfo.Statistic.Security != nil {
			if err := db.CreateSecurityData(*value.NodeInfo.Statistic.Security,
				value.NodeInfo.NodeAuthData.Uuid); err != nil {
				log.Println(err)
			}
		}
		if value.NodeInfo.Statistic.Updates != nil {
			if err := db.CreateUpdatesData(*value.NodeInfo.Statistic.Updates,
				value.NodeInfo.NodeAuthData.Uuid); err != nil {
				log.Println(err)
			}
		}
		if value.NodeInfo.Statistic.Blocks != nil {
			if err := db.CreateBlocksData(*value.NodeInfo.Statistic.Blocks,
				value.NodeInfo.NodeAuthData.Uuid); err != nil {
				log.Println(err)
			}
		}
		if value.NodeInfo.Statistic.Epoch != nil {
			if err := db.CreateEpochData(*value.NodeInfo.Statistic.Epoch,
				value.NodeInfo.NodeAuthData.Uuid); err != nil {
				log.Println(err)
			}
		}
		if value.NodeInfo.Statistic.NodeBasicData != nil {
			if err := db.CreateNodeBasicData(*value.NodeInfo.Statistic.NodeBasicData,
				value.NodeInfo.NodeAuthData.Uuid); err != nil {
				log.Println(err)
			}
		}
		if value.NodeInfo.Statistic.CpuState != nil {
			if err := db.CreateCpuStateData(*value.NodeInfo.Statistic.CpuState,
				value.NodeInfo.NodeAuthData.Uuid); err != nil {
				log.Println(err)
			}
		}
		if value.NodeInfo.Statistic.KesData != nil {
			if err := db.CreateKesData(*value.NodeInfo.Statistic.KesData,
				value.NodeInfo.NodeAuthData.Uuid); err != nil {
				log.Println(err)
			}
		}
	}
}
