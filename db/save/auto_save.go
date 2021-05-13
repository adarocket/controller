package save

import (
	"adarocket/controller/db/postgresql"
	"adarocket/controller/informer"
	"log"
	"time"
)

const timeout = 10

// make a part of InformerServer structs?
func AutoSave(server *informer.InformServer) {
	//postgresql.InitDatabase()
	db := postgresql.Postg

	go func() {
		for _ = range time.Tick(time.Minute * timeout) {
			for _, value := range server.NodeStatistics {
				// if already exist?
				if err := db.CreateNodeAuthData(*value.NodeAuthData); err != nil {
					log.Println(err)
				}

				if err := db.CreateServerBasicData(*value.Statistic.ServerBasicData, value.NodeAuthData.Uuid); err != nil {
					log.Println(err)
				}
				if err := db.CreateNodeStateData(*value.Statistic.NodeState, value.NodeAuthData.Uuid); err != nil {
					log.Println(err)
				}
				if err := db.CreateNodePerformanceData(*value.Statistic.NodePerformance, value.NodeAuthData.Uuid); err != nil {
					log.Println(err)
				}
				if err := db.CreateMemoryStateData(*value.Statistic.MemoryState, value.NodeAuthData.Uuid); err != nil {
					log.Println(err)
				}
				if err := db.CreateOnlineData(*value.Statistic.Online, value.NodeAuthData.Uuid); err != nil {
					log.Println(err)
				}
				if err := db.CreateStakeInfoData(*value.Statistic.StakeInfo, value.NodeAuthData.Uuid); err != nil {
					log.Println(err)
				}
				if err := db.CreateSecurityData(*value.Statistic.Security, value.NodeAuthData.Uuid); err != nil {
					log.Println(err)
				}
				if err := db.CreateUpdatesData(*value.Statistic.Updates, value.NodeAuthData.Uuid); err != nil {
					log.Println(err)
				}
				if err := db.CreateBlocksData(*value.Statistic.Blocks, value.NodeAuthData.Uuid); err != nil {
					log.Println(err)
				}
				if err := db.CreateEpochData(*value.Statistic.Epoch, value.NodeAuthData.Uuid); err != nil {
					log.Println(err)
				}
				if err := db.CreateNodeBasicData(*value.Statistic.NodeBasicData, value.NodeAuthData.Uuid); err != nil {
					log.Println(err)
				}
				if err := db.CreateCpuStateData(*value.Statistic.CpuState, value.NodeAuthData.Uuid); err != nil {
					log.Println(err)
				}
				if err := db.CreateKesData(*value.Statistic.KesData, value.NodeAuthData.Uuid); err != nil {
					log.Println(err)
				}
			}
		}
	}()
}
