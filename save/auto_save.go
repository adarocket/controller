package save

import (
	"adarocket/controller/informer"
	"adarocket/controller/postgresql"
	"log"
	"time"
)

const timeout = 10

func AutoSave(server informer.InformServer) {
	postgresql.InitDatabase()
	db := postgresql.Postg

	go func() {
		for _ = range time.Tick(time.Minute * timeout) {
			for _, value := range server.NodeStatistics {
				// if already exist?
				if err := db.CreateNodeAuthData(*value.NodeAuthData); err != nil {
					log.Println(err)
				}

				if err := db.CreateServerBasicData(*value.Statistic.ServerBasicData); err != nil {
					log.Println(err)
				}
				if err := db.CreateNodeStateData(*value.Statistic.NodeState); err != nil {
					log.Println(err)
				}
				if err := db.CreateNodePerformanceData(*value.Statistic.NodePerformance); err != nil {
					log.Println(err)
				}
				if err := db.CreateMemoryStateData(*value.Statistic.MemoryState); err != nil {
					log.Println(err)
				}
				if err := db.CreateOnlineData(*value.Statistic.Online); err != nil {
					log.Println(err)
				}
				if err := db.CreateStakeInfoData(*value.Statistic.StakeInfo); err != nil {
					log.Println(err)
				}
				if err := db.CreateSecurityData(*value.Statistic.Security); err != nil {
					log.Println(err)
				}
				if err := db.CreateUpdatesData(*value.Statistic.Updates); err != nil {
					log.Println(err)
				}
				if err := db.CreateBlocksData(*value.Statistic.Blocks); err != nil {
					log.Println(err)
				}
				if err := db.CreateEpochData(*value.Statistic.Epoch); err != nil {
					log.Println(err)
				}
				if err := db.CreateNodeBasicData(*value.Statistic.NodeBasicData); err != nil {
					log.Println(err)
				}
				if err := db.CreateCpuStateData(*value.Statistic.CpuState); err != nil {
					log.Println(err)
				}
				if err := db.CreateKesData(*value.Statistic.KesData); err != nil {
					log.Println(err)
				}
			}
		}
	}()
}
