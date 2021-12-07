package save

import (
	"github.com/adarocket/controller/db/structs"
	"github.com/adarocket/controller/informer"
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

		if value.NodeAuthData != nil {
			if err := db.CreateNodeAuthData(*value.NodeAuthData); err != nil {
				log.Println(err)
			}
		}

		var stats *cardano.Statistic
		if stats = value.GetStatistic(); stats == nil {
			continue
		}

		if stats.ServerBasicData != nil {
			if err := db.CreateServerBasicData(*stats.ServerBasicData,
				value.NodeAuthData.Uuid); err != nil {
				log.Println(err)
			}
		}
		if stats.NodeState != nil {
			if err := db.CreateNodeStateData(*stats.NodeState,
				value.NodeAuthData.Uuid); err != nil {
				log.Println(err)
			}
		}
		if stats.NodePerformance != nil {
			if err := db.CreateNodePerformanceData(*stats.NodePerformance,
				value.NodeAuthData.Uuid); err != nil {
				log.Println(err)
			}
		}
		if stats.MemoryState != nil {
			if err := db.CreateMemoryStateData(*stats.MemoryState,
				value.NodeAuthData.Uuid); err != nil {
				log.Println(err)
			}
		}
		if stats.Online != nil {
			if err := db.CreateOnlineData(*stats.Online,
				value.NodeAuthData.Uuid); err != nil {
				log.Println(err)
			}
		}
		if stats.StakeInfo != nil {
			if err := db.CreateStakeInfoData(*stats.StakeInfo,
				value.NodeAuthData.Uuid); err != nil {
				log.Println(err)
			}
		}
		if stats.Security != nil {
			if err := db.CreateSecurityData(*stats.Security,
				value.NodeAuthData.Uuid); err != nil {
				log.Println(err)
			}
		}
		if stats.Updates != nil {
			if err := db.CreateUpdatesData(*stats.Updates,
				value.NodeAuthData.Uuid); err != nil {
				log.Println(err)
			}
		}
		if stats.Blocks != nil {
			if err := db.CreateBlocksData(*stats.Blocks,
				value.NodeAuthData.Uuid); err != nil {
				log.Println(err)
			}
		}
		if stats.Epoch != nil {
			if err := db.CreateEpochData(*stats.Epoch,
				value.NodeAuthData.Uuid); err != nil {
				log.Println(err)
			}
		}
		if stats.NodeBasicData != nil {
			if err := db.CreateNodeBasicData(*stats.NodeBasicData,
				value.NodeAuthData.Uuid); err != nil {
				log.Println(err)
			}
		}
		if stats.CpuState != nil {
			if err := db.CreateCpuStateData(*stats.CpuState,
				value.NodeAuthData.Uuid); err != nil {
				log.Println(err)
			}
		}
		if stats.KesData != nil {
			if err := db.CreateKesData(*stats.KesData,
				value.NodeAuthData.Uuid); err != nil {
				log.Println(err)
			}
		}
	}
}
