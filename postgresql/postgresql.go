package postgresql

import (
	"database/sql"
	pb "github.com/adarocket/proto"
	"log"
)

// PostgreSQL ...
var _ postgresql

// InitDatabase ...
func InitDatabase() {
	/*// todo везде где может возникать ошибка, нужно ее возвращать
	config, readerErr := reader.ReadJSON("config/dbConfig.json")
	if readerErr != nil {
		log.Fatal(readerErr)
	}
	connStr := fmt.Sprintf(`user=%s password=%s dbname=%s sslmode=%s`,
		config.User, config.Password, config.Dbname, config.Sslmode)
	// connStr := "user = postgres password=postgresql dbname=crypto sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	PostgreSQL.dbConn = db*/

}

type postgresql struct {
	dbConn *sql.DB
}

const getNodeAuthQuery = `
	SELECT Ticker, Uuid, Status 
	FROM NodeAuth`

func (p postgresql) GetNodeAuthData() ([]pb.NodeAuthData, error) {
	rows, err := p.dbConn.Query(getNodeAuthQuery)
	if err != nil {
		log.Fatal("GetNodesAuthData", err)
		return []pb.NodeAuthData{}, err
	}
	defer rows.Close()

	nodesAuthData := make([]pb.NodeAuthData, 0, 10)
	for rows.Next() {
		nodeAuthData := pb.NodeAuthData{}
		if err := rows.Scan(&nodeAuthData.Ticker, &nodeAuthData.Uuid,
			&nodeAuthData.Status); err != nil {
			log.Println("NodesAuth: parse err", err)
			continue
		}
		// можно ли так делать?
		nodesAuthData = append(nodesAuthData, nodeAuthData)
	}

	return nodesAuthData, nil
}

const createNodeAuthExec = `
	INSERT INTO NodeAuth 
	(Ticker, Uuid, Status) 
	VALUES (?, ?, ?)`

func (p postgresql) CreateNodeAuthData(data pb.NodeAuthData) error {
	if _, err := p.dbConn.Exec(createNodeAuthExec,
		data.Ticker, data.Uuid, data.Status); err != nil {
		log.Println("CreateNodeAuth", err)
		return err
	}

	return nil
}

func (p postgresql) UpdateNodeAuthData(data pb.NodeAuthData) error {
	panic("implement me")
}

const deleteNodeAuthExec = `
	DELETE FROM NodeAuth
	WHERE Ticker = ?`

func (p postgresql) DeleteNodeAuthData(data pb.NodeAuthData) error {
	if _, err := p.dbConn.Exec(deleteNodeAuthExec, data.Ticker); err != nil {
		log.Println("DeleteNodeAuth", err)
		return err
	}

	return nil
}

const getNodeBasicDataQuery = ``

func (p postgresql) GetNodeBasicData() ([]pb.NodeBasicData, error) {
	rows, err := p.dbConn.Query(getNodeBasicDataQuery)
	if err != nil {
		log.Fatal("GetNodesAuthData", err)
		return []pb.NodeBasicData{}, err
	}
	defer rows.Close()

	nodesBasicData := make([]pb.NodeBasicData, 0, 10)
	for rows.Next() {
		nodeBasicData := pb.NodeBasicData{}
		if err := rows.Scan(&nodeBasicData.Ticker, &nodeBasicData.Type,
			&nodeBasicData.Location, &nodeBasicData.NodeVersion); err != nil {
			log.Println("NodesAuth: parse err", err)
			continue
		}
		// можно ли так делать?
		nodesBasicData = append(nodesBasicData, nodeBasicData)
	}

	return nodesBasicData, nil
}

const createNodeBasicDataExec = ``

func (p postgresql) CreateNodeBasicData(data pb.NodeBasicData) error {
	if _, err := p.dbConn.Exec(createNodeBasicDataExec,
		data.Ticker, data.Type, data.Location, data.NodeVersion); err != nil {
		log.Println("CreateNodeAuth", err)
		return err
	}

	return nil
}

const getServerBasicData = ``

func (p postgresql) GetServerBasicData() ([]pb.ServerBasicData, error) {
	rows, err := p.dbConn.Query(getServerBasicData)
	if err != nil {
		log.Fatal("GetNodesAuthData", err)
		return []pb.ServerBasicData{}, err
	}
	defer rows.Close()

	serverBasicDates := make([]pb.ServerBasicData, 0, 10)
	for rows.Next() {
		serverBasicData := pb.ServerBasicData{}
		if err := rows.Scan(&serverBasicData.Ipv4, &serverBasicData.Ipv6,
			&serverBasicData.LinuxName, &serverBasicData.LinuxVersion); err != nil {
			log.Println("NodesAuth: parse err", err)
			continue
		}
		// можно ли так делать?
		serverBasicDates = append(serverBasicDates, serverBasicData)
	}

	return serverBasicDates, nil
}

const createServerBasicDataExec = ``

func (p postgresql) CreateServerBasicData(data pb.ServerBasicData) error {
	if _, err := p.dbConn.Exec(createServerBasicDataExec,
		data.Ipv4, data.Ipv6, data.LinuxName, data.LinuxVersion); err != nil {
		log.Println("CreateNodeAuth", err)
		return err
	}

	return nil
}

const getEpochDataQuery = ``

func (p postgresql) GetEpochData() ([]pb.Epoch, error) {
	rows, err := p.dbConn.Query(getEpochDataQuery)
	if err != nil {
		log.Fatal("GetNodesAuthData", err)
		return []pb.Epoch{}, err
	}
	defer rows.Close()

	epochDates := make([]pb.Epoch, 0, 10)
	for rows.Next() {
		epochData := pb.Epoch{}
		if err := rows.Scan(&epochData.EpochNumber); err != nil {
			log.Println("NodesAuth: parse err", err)
			continue
		}
		// можно ли так делать?
		epochDates = append(epochDates, epochData)
	}

	return epochDates, nil
}

const createEpochDataExec = ``

func (p postgresql) CreateEpochData(data pb.Epoch) error {
	if _, err := p.dbConn.Exec(createEpochDataExec,
		data.EpochNumber); err != nil {
		log.Println("CreateNodeAuth", err)
		return err
	}

	return nil
}

const getKesDataQuery = ``

func (p postgresql) GetKesData() ([]pb.KESData, error) {
	rows, err := p.dbConn.Query(getKesDataQuery)
	if err != nil {
		log.Fatal("GetNodesAuthData", err)
		return []pb.KESData{}, err
	}
	defer rows.Close()

	kesDates := make([]pb.KESData, 0, 10)
	for rows.Next() {
		kesData := pb.KESData{}
		if err := rows.Scan(&kesData.KesCurrent,
			&kesData.KesRemaining, &kesData.KesExpDate); err != nil {
			log.Println("NodesAuth: parse err", err)
			continue
		}
		// можно ли так делать?
		kesDates = append(kesDates, kesData)
	}

	return kesDates, nil
}

const createKesDataExec = ``

func (p postgresql) CreateKesData(data pb.KESData) error {
	if _, err := p.dbConn.Exec(createKesDataExec,
		data.KesCurrent, data.KesRemaining, data.KesExpDate); err != nil {
		log.Println("CreateNodeAuth", err)
		return err
	}

	return nil
}

const getBlocksDataQuery = ``

func (p postgresql) GetBlocksData() ([]pb.Blocks, error) {
	rows, err := p.dbConn.Query(getBlocksDataQuery)
	if err != nil {
		log.Fatal("GetNodesAuthData", err)
		return []pb.Blocks{}, err
	}
	defer rows.Close()

	blockDates := make([]pb.Blocks, 0, 10)
	for rows.Next() {
		blockData := pb.Blocks{}
		if err := rows.Scan(&blockData.BlockLeader,
			&blockData.BlockAdopted, &blockData.BlockInvalid); err != nil {
			log.Println("NodesAuth: parse err", err)
			continue
		}
		// можно ли так делать?
		blockDates = append(blockDates, blockData)
	}

	return blockDates, nil
}

const createBlocksDataExec = ``

func (p postgresql) CreateBlocksData(data pb.Blocks) error {
	if _, err := p.dbConn.Exec(createBlocksDataExec,
		data.BlockLeader, data.BlockAdopted, data.BlockInvalid); err != nil {
		log.Println("CreateNodeAuth", err)
		return err
	}

	return nil
}

const getUpdatesDataQuery = ``

func (p postgresql) GetUpdatesData() ([]pb.Updates, error) {
	rows, err := p.dbConn.Query(getUpdatesDataQuery)
	if err != nil {
		log.Fatal("GetNodesAuthData", err)
		return []pb.Updates{}, err
	}
	defer rows.Close()

	updatesDates := make([]pb.Updates, 0, 10)
	for rows.Next() {
		updateData := pb.Updates{}
		if err := rows.Scan(&updateData.InformerActual,
			&updateData.InformerAvailable, &updateData.UpdaterActual,
			&updateData.UpdaterAvailable, &updateData.PackagesAvailable); err != nil {
			log.Println("NodesAuth: parse err", err)
			continue
		}
		// можно ли так делать?
		updatesDates = append(updatesDates, updateData)
	}

	return updatesDates, nil
}

const createUpdatesDataExec = ``

func (p postgresql) CreateUpdatesData(data pb.Updates) error {
	if _, err := p.dbConn.Exec(createUpdatesDataExec,
		data.InformerActual, data.InformerAvailable, data.UpdaterActual,
		data.UpdaterAvailable, data.PackagesAvailable); err != nil {
		log.Println("CreateNodeAuth", err)
		return err
	}

	return nil
}

const getSecurityDataQuery = ``

func (p postgresql) GetSecurityData() ([]pb.Security, error) {
	rows, err := p.dbConn.Query(getSecurityDataQuery)
	if err != nil {
		log.Fatal("GetNodesAuthData", err)
		return []pb.Security{}, err
	}
	defer rows.Close()

	securityDates := make([]pb.Security, 0, 10)
	for rows.Next() {
		securityData := pb.Security{}
		if err := rows.Scan(&securityData.SshAttackAttempts,
			&securityData.SecurityPackagesAvailable); err != nil {
			log.Println("NodesAuth: parse err", err)
			continue
		}
		// можно ли так делать?
		securityDates = append(securityDates, securityData)
	}

	return securityDates, nil
}

const createSecurityDataExec = ``

func (p postgresql) CreateSecurityData(data pb.Security) error {
	if _, err := p.dbConn.Exec(createSecurityDataExec,
		data.SshAttackAttempts, data.SecurityPackagesAvailable); err != nil {
		log.Println("CreateNodeAuth", err)
		return err
	}

	return nil
}

const getStakeInfoDataQuery = ``

func (p postgresql) GetStakeInfoData() ([]pb.StakeInfo, error) {
	rows, err := p.dbConn.Query(getStakeInfoDataQuery)
	if err != nil {
		log.Fatal("GetNodesAuthData", err)
		return []pb.StakeInfo{}, err
	}
	defer rows.Close()

	stakeInfoDates := make([]pb.StakeInfo, 0, 10)
	for rows.Next() {
		stakeInfoData := pb.StakeInfo{}
		if err := rows.Scan(&stakeInfoData.LiveStake,
			&stakeInfoData.ActiveStake, &stakeInfoData.Pledge); err != nil {
			log.Println("NodesAuth: parse err", err)
			continue
		}
		// можно ли так делать?
		stakeInfoDates = append(stakeInfoDates, stakeInfoData)
	}

	return stakeInfoDates, nil
}

const createStakeInfoDataExec = ``

func (p postgresql) CreateStakeInfoData(data pb.StakeInfo) error {
	if _, err := p.dbConn.Exec(createStakeInfoDataExec,
		data.LiveStake, data.ActiveStake, data.Pledge); err != nil {
		log.Println("CreateNodeAuth", err)
		return err
	}

	return nil
}

const getOnlineDataQuery = ``

func (p postgresql) GetOnlineData() ([]pb.Online, error) {
	rows, err := p.dbConn.Query(getOnlineDataQuery)
	if err != nil {
		log.Fatal("GetNodesAuthData", err)
		return []pb.Online{}, err
	}
	defer rows.Close()

	onlineDates := make([]pb.Online, 0, 10)
	for rows.Next() {
		onlineData := pb.Online{}
		if err := rows.Scan(&onlineData.SinceStart,
			&onlineData.Pings, &onlineData.NodeActive,
			&onlineData.NodeActive, &onlineData.NodeActivePings,
			&onlineData.ServerActive); err != nil {
			log.Println("NodesAuth: parse err", err)
			continue
		}
		// можно ли так делать?
		onlineDates = append(onlineDates, onlineData)
	}

	return onlineDates, nil
}

const createOnlineDataExec = ``

func (p postgresql) CreateOnlineData(data pb.Online) error {
	if _, err := p.dbConn.Exec(createOnlineDataExec,
		data.SinceStart, data.Pings, data.NodeActive,
		data.NodeActivePings, data.ServerActive); err != nil {
		log.Println("CreateNodeAuth", err)
		return err
	}

	return nil
}

const getMemoryStateDataQuery = ``

func (p postgresql) GetMemoryStateData() ([]pb.MemoryState, error) {
	rows, err := p.dbConn.Query(getMemoryStateDataQuery)
	if err != nil {
		log.Fatal("GetNodesAuthData", err)
		return []pb.MemoryState{}, err
	}
	defer rows.Close()

	memoryStateDates := make([]pb.MemoryState, 0, 10)
	for rows.Next() {
		memoryStateData := pb.MemoryState{}
		if err := rows.Scan(&memoryStateData.Total,
			&memoryStateData.Used, &memoryStateData.Buffers,
			&memoryStateData.Cached, &memoryStateData.Cached,
			&memoryStateData.Free, &memoryStateData.Available,
			&memoryStateData.Available, &memoryStateData.Active,
			&memoryStateData.Inactive, &memoryStateData.SwapTotal,
			&memoryStateData.SwapUsed, &memoryStateData.SwapCached,
			&memoryStateData.SwapFree, &memoryStateData.MemAvailableEnabled); err != nil {
			log.Println("NodesAuth: parse err", err)
			continue
		}
		// можно ли так делать?
		memoryStateDates = append(memoryStateDates, memoryStateData)
	}

	return memoryStateDates, nil
}

const createMemoryStateDataExec = ``

func (p postgresql) CreateMemoryStateData(data pb.MemoryState) error {
	if _, err := p.dbConn.Exec(createMemoryStateDataExec,
		data.Total, data.Used, data.Buffers,
		data.Cached, data.Free, data.Available,
		data.Active, data.Inactive, data.SwapTotal,
		data.SwapUsed, data.SwapCached, data.SwapFree, data.MemAvailableEnabled); err != nil {
		log.Println("CreateNodeAuth", err)
		return err
	}

	return nil
}

const getNodePerformanceDataQuery = ``

func (p postgresql) GetNodePerformanceData() ([]pb.NodePerformance, error) {
	rows, err := p.dbConn.Query(getNodePerformanceDataQuery)
	if err != nil {
		log.Fatal("GetNodesAuthData", err)
		return []pb.NodePerformance{}, err
	}
	defer rows.Close()

	nodePerformanceDates := make([]pb.NodePerformance, 0, 10)
	for rows.Next() {
		nodePerformanceData := pb.NodePerformance{}
		if err := rows.Scan(&nodePerformanceData.ProcessedTx,
			&nodePerformanceData.PeersIn, &nodePerformanceData.PeersOut); err != nil {
			log.Println("NodesAuth: parse err", err)
			continue
		}
		// можно ли так делать?
		nodePerformanceDates = append(nodePerformanceDates, nodePerformanceData)
	}

	return nodePerformanceDates, nil
}

const createNodePerformanceDataExec = ``

func (p postgresql) CreateNodePerformanceData(data pb.NodePerformance) error {
	if _, err := p.dbConn.Exec(createNodePerformanceDataExec,
		data.ProcessedTx, data.PeersIn, data.PeersOut); err != nil {
		log.Println("CreateNodeAuth", err)
		return err
	}

	return nil
}

const getCpuStateDataQuery = ``

func (p postgresql) GetCpuStateData() ([]pb.CPUState, error) {
	rows, err := p.dbConn.Query(getCpuStateDataQuery)
	if err != nil {
		log.Fatal("GetNodesAuthData", err)
		return []pb.CPUState{}, err
	}
	defer rows.Close()

	cpuStateDates := make([]pb.CPUState, 0, 10)
	for rows.Next() {
		cpuStateData := pb.CPUState{}
		if err := rows.Scan(&cpuStateData.CpuQty,
			&cpuStateData.AverageWorkload); err != nil {
			log.Println("NodesAuth: parse err", err)
			continue
		}
		// можно ли так делать?
		cpuStateDates = append(cpuStateDates, cpuStateData)
	}

	return cpuStateDates, nil
}

const createCpuStateDataExec = ``

func (p postgresql) CreateCpuStateData(data pb.CPUState) error {
	if _, err := p.dbConn.Exec(createCpuStateDataExec,
		data.CpuQty, data.AverageWorkload); err != nil {
		log.Println("CreateNodeAuth", err)
		return err
	}

	return nil
}

const getNodesStateDataQuery = ``

func (p postgresql) GetNodeStateData() ([]pb.NodeState, error) {
	rows, err := p.dbConn.Query(getNodesStateDataQuery)
	if err != nil {
		log.Fatal("GetNodesAuthData", err)
		return []pb.NodeState{}, err
	}
	defer rows.Close()

	nodeStateDates := make([]pb.NodeState, 0, 10)
	for rows.Next() {
		nodeStateData := pb.NodeState{}
		if err := rows.Scan(&nodeStateData.TipDiff,
			&nodeStateData.Density); err != nil {
			log.Println("NodesAuth: parse err", err)
			continue
		}
		// можно ли так делать?
		nodeStateDates = append(nodeStateDates, nodeStateData)
	}

	return nodeStateDates, nil
}

const createNodeStateData = ``

func (p postgresql) CreateNodeStateData(data pb.NodeState) error {
	if _, err := p.dbConn.Exec(createNodeStateData,
		data.TipDiff, data.Density); err != nil {
		log.Println("CreateNodeAuth", err)
		return err
	}

	return nil
}
