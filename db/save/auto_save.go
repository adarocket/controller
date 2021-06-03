package save

import (
	"adarocket/controller/db/postgresql"
	"adarocket/controller/informer"
	"log"
	"time"
)

const timeout = 10

func AutoSave(server *informer.InformServer) {
	db := postgresql.Postg

	go func() {
		for _ = range time.Tick(time.Minute * timeout) {
			for _, value := range server.NodeStatistics {
				if err := db.CreateNodeAuthData(*value.NodeInfo.NodeAuthData); err != nil {
					log.Println(err)
				}

				if err := db.CreateServerBasicData(*value.NodeInfo.Statistic.ServerBasicData, value.NodeInfo.NodeAuthData.Uuid); err != nil {
					log.Println(err)
				}
				if err := db.CreateNodeStateData(*value.NodeInfo.Statistic.NodeState, value.NodeInfo.NodeAuthData.Uuid); err != nil {
					log.Println(err)
				}
				if err := db.CreateNodePerformanceData(*value.NodeInfo.Statistic.NodePerformance, value.NodeInfo.NodeAuthData.Uuid); err != nil {
					log.Println(err)
				}
				if err := db.CreateMemoryStateData(*value.NodeInfo.Statistic.MemoryState, value.NodeInfo.NodeAuthData.Uuid); err != nil {
					log.Println(err)
				}
				if err := db.CreateOnlineData(*value.NodeInfo.Statistic.Online, value.NodeInfo.NodeAuthData.Uuid); err != nil {
					log.Println(err)
				}
				if err := db.CreateStakeInfoData(*value.NodeInfo.Statistic.StakeInfo, value.NodeInfo.NodeAuthData.Uuid); err != nil {
					log.Println(err)
				}
				if err := db.CreateSecurityData(*value.NodeInfo.Statistic.Security, value.NodeInfo.NodeAuthData.Uuid); err != nil {
					log.Println(err)
				}
				if err := db.CreateUpdatesData(*value.NodeInfo.Statistic.Updates, value.NodeInfo.NodeAuthData.Uuid); err != nil {
					log.Println(err)
				}
				if err := db.CreateBlocksData(*value.NodeInfo.Statistic.Blocks, value.NodeInfo.NodeAuthData.Uuid); err != nil {
					log.Println(err)
				}
				if err := db.CreateEpochData(*value.NodeInfo.Statistic.Epoch, value.NodeInfo.NodeAuthData.Uuid); err != nil {
					log.Println(err)
				}
				if err := db.CreateNodeBasicData(*value.NodeInfo.Statistic.NodeBasicData, value.NodeInfo.NodeAuthData.Uuid); err != nil {
					log.Println(err)
				}
				if err := db.CreateCpuStateData(*value.NodeInfo.Statistic.CpuState, value.NodeInfo.NodeAuthData.Uuid); err != nil {
					log.Println(err)
				}
				if err := db.CreateKesData(*value.NodeInfo.Statistic.KesData, value.NodeInfo.NodeAuthData.Uuid); err != nil {
					log.Println(err)
				}
			}
		}
	}()
}
