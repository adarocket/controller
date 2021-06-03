package postgresql

import (
	"log"

	pb "github.com/adarocket/proto/proto"
)

const getNodeAuthQuery = `
	SELECT Ticker, Uuid, Status 
	FROM NodeAuth
`

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
	VALUES ($1, $2, $3)
	ON CONFLICT (uuid) DO UPDATE 
  	SET status = excluded.status;
`

func (p postgresql) CreateNodeAuthData(data pb.NodeAuthData) error {
	if _, err := p.dbConn.Exec(createNodeAuthExec,
		data.Ticker, data.Uuid, data.Status); err != nil {
		log.Println("CreateNodeAuth", err)
		return err
	}

	return nil
}

const updateNodeAuthData = `
	UPDATE nodeauth
	SET ticker = $1, status = $2
	WHERE uuid = $3
`

func (p postgresql) UpdateNodeAuthData(data pb.NodeAuthData) error {
	_, err := p.dbConn.Exec(updateNodeAuthData,
		data.Ticker, data.Status)
	if err != nil {
		log.Println("UpdateToken", err)
		return err
	}

	return nil
}

const deleteNodeAuthExec = `
	DELETE FROM nodeauth
	WHERE uuid = $1
`

func (p postgresql) DeleteNodeAuthData(data pb.NodeAuthData) error {
	if _, err := p.dbConn.Exec(deleteNodeAuthExec, data.Ticker); err != nil {
		log.Println("DeleteNodeAuth", err)
		return err
	}

	return nil
}

const getNodeBasicDataQuery = `
	SELECT ticker, type, location, nodeversion 
	FROM NodeBasicData
`

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

const createNodeBasicDataExec = `
	INSERT INTO nodebasicdata 
	(uuid, ticker, type, location, nodeversion) 
	VALUES ($1, $2, $3, $4, $5)
	ON CONFLICT (uuid) DO UPDATE 
  	SET ticker = excluded.ticker,
    	type = excluded.type,
  		location = excluded.location,
  	    nodeversion = excluded.nodeversion;
`

func (p postgresql) CreateNodeBasicData(data pb.NodeBasicData, uuid string) error {
	if _, err := p.dbConn.Exec(createNodeBasicDataExec,
		uuid, data.Ticker, data.Type, data.Location, data.NodeVersion); err != nil {
		log.Println("CreateNodeBasicData", err)
		return err
	}

	return nil
}

const getServerBasicData = `
	SELECT ipv4, ipv6, linuxname, linuxversion
	FROM serverbasicdata
`

func (p postgresql) GetServerBasicData() ([]pb.ServerBasicData, error) {
	rows, err := p.dbConn.Query(getServerBasicData)
	if err != nil {
		log.Fatal("GetServerBasicData", err)
		return []pb.ServerBasicData{}, err
	}
	defer rows.Close()

	serverBasicDates := make([]pb.ServerBasicData, 0, 10)
	for rows.Next() {
		serverBasicData := pb.ServerBasicData{}
		if err := rows.Scan(&serverBasicData.Ipv4, &serverBasicData.Ipv6,
			&serverBasicData.LinuxName, &serverBasicData.LinuxVersion); err != nil {
			log.Println("serverBasicData: parse err", err)
			continue
		}
		// можно ли так делать?
		serverBasicDates = append(serverBasicDates, serverBasicData)
	}

	return serverBasicDates, nil
}

const createServerBasicDataExec = `
	INSERT INTO serverbasicdata 
	(uuid, ipv4, ipv6, linuxname, linuxversion) 
	VALUES ($1, $2, $3, $4, $5)
	ON CONFLICT (uuid) DO UPDATE 
  	SET ipv4 = excluded.ipv4, 
      	ipv6 = excluded.ipv6,
  	    linuxname = excluded.linuxname,
  	    linuxversion = excluded.linuxversion;
`

func (p postgresql) CreateServerBasicData(data pb.ServerBasicData, uuid string) error {
	if _, err := p.dbConn.Exec(createServerBasicDataExec,
		uuid, data.Ipv4, data.Ipv6, data.LinuxName, data.LinuxVersion); err != nil {
		log.Println("CreateServerBasicData", err)
		return err
	}

	return nil
}

const getEpochDataQuery = `
	SELECT epochnumber
	FROM epochdata
`

func (p postgresql) GetEpochData() ([]pb.Epoch, error) {
	rows, err := p.dbConn.Query(getEpochDataQuery)
	if err != nil {
		log.Fatal("GetEpochData", err)
		return []pb.Epoch{}, err
	}
	defer rows.Close()

	epochDates := make([]pb.Epoch, 0, 10)
	for rows.Next() {
		epochData := pb.Epoch{}
		if err := rows.Scan(&epochData.EpochNumber); err != nil {
			log.Println("epochData: parse err", err)
			continue
		}
		// можно ли так делать?
		epochDates = append(epochDates, epochData)
	}

	return epochDates, nil
}

const createEpochDataExec = `
	INSERT INTO epochdata
	(uuid, epochnumber)
	VALUES ($1, $2)
	ON CONFLICT (uuid) DO UPDATE 
  	SET epochnumber = excluded.epochnumber;
`

func (p postgresql) CreateEpochData(data pb.Epoch, uuid string) error {
	if _, err := p.dbConn.Exec(createEpochDataExec,
		uuid, data.EpochNumber); err != nil {
		log.Println("CreateEpochData", err)
		return err
	}

	return nil
}

const getKesDataQuery = `
	SELECT epochnumber
	FROM epochdata
`

func (p postgresql) GetKesData() ([]pb.KESData, error) {
	rows, err := p.dbConn.Query(getKesDataQuery)
	if err != nil {
		log.Fatal("GetKesData", err)
		return []pb.KESData{}, err
	}
	defer rows.Close()

	kesDates := make([]pb.KESData, 0, 10)
	for rows.Next() {
		kesData := pb.KESData{}
		if err := rows.Scan(&kesData.KesCurrent,
			&kesData.KesRemaining, &kesData.KesExpDate); err != nil {
			log.Println("kesData: parse err", err)
			continue
		}
		// можно ли так делать?
		kesDates = append(kesDates, kesData)
	}

	return kesDates, nil
}

const createKesDataExec = `
	INSERT INTO kesdata
	(uuid, kescurrent, kesremaining, kesexpdate) 
	VALUES ($1, $2, $3, $4)
	ON CONFLICT (uuid) DO UPDATE 
  	SET kescurrent = excluded.kescurrent,
  	    kesremaining = excluded.kesremaining,
  	    kesexpdate = excluded.kesexpdate;  	    
`

func (p postgresql) CreateKesData(data pb.KESData, uuid string) error {
	if _, err := p.dbConn.Exec(createKesDataExec,
		uuid, data.KesCurrent, data.KesRemaining, data.KesExpDate); err != nil {
		log.Println("CreateKesData", err)
		return err
	}

	return nil
}

const getBlocksDataQuery = `
	SELECT blockleader, blockadopted, blockinvalid
	FROM blocksdata
`

func (p postgresql) GetBlocksData() ([]pb.Blocks, error) {
	rows, err := p.dbConn.Query(getBlocksDataQuery)
	if err != nil {
		log.Fatal("GetBlocksData", err)
		return []pb.Blocks{}, err
	}
	defer rows.Close()

	blockDates := make([]pb.Blocks, 0, 10)
	for rows.Next() {
		blockData := pb.Blocks{}
		if err := rows.Scan(&blockData.BlockLeader,
			&blockData.BlockAdopted, &blockData.BlockInvalid); err != nil {
			log.Println("blockData: parse err", err)
			continue
		}
		// можно ли так делать?
		blockDates = append(blockDates, blockData)
	}

	return blockDates, nil
}

const createBlocksDataExec = `
	INSERT INTO blocksdata
	(uuid, blockleader, blockadopted, blockinvalid)
	VALUES ($1, $2, $3, $4)
	ON CONFLICT (uuid) DO UPDATE 
  	SET blockleader = excluded.blockleader,
  	    blockadopted = excluded.blockadopted,
  	    blockinvalid = excluded.blockinvalid;  
`

func (p postgresql) CreateBlocksData(data pb.Blocks, uuid string) error {
	if _, err := p.dbConn.Exec(createBlocksDataExec,
		uuid, data.BlockLeader, data.BlockAdopted, data.BlockInvalid); err != nil {
		log.Println("CreateBlocksData", err)
		return err
	}

	return nil
}

const getUpdatesDataQuery = `
	SELECT INFORMERACTUAL, INFORMERAVAILABLE, UPDATERACTUAL, UPDATERAVAILABLE, PACKAGESAVAILABLE
	FROM updatesdata
`

func (p postgresql) GetUpdatesData() ([]pb.Updates, error) {
	rows, err := p.dbConn.Query(getUpdatesDataQuery)
	if err != nil {
		log.Fatal("GetUpdatesData", err)
		return []pb.Updates{}, err
	}
	defer rows.Close()

	updatesDates := make([]pb.Updates, 0, 10)
	for rows.Next() {
		updateData := pb.Updates{}
		if err := rows.Scan(&updateData.InformerActual,
			&updateData.InformerAvailable, &updateData.UpdaterActual,
			&updateData.UpdaterAvailable, &updateData.PackagesAvailable); err != nil {
			log.Println("updateData: parse err", err)
			continue
		}
		// можно ли так делать?
		updatesDates = append(updatesDates, updateData)
	}

	return updatesDates, nil
}

const createUpdatesDataExec = `
	INSERT INTO updatesdata
	(uuid, informeractual, informeravailable, updateractual, updateravailable, packagesavailable)
	VALUES ($1, $2, $3, $4, $5, $6)
	ON CONFLICT (uuid) DO UPDATE 
  	SET informeractual = excluded.informeractual,
  	    informeravailable = excluded.informeravailable,
  	    updateractual = excluded.updateractual, 
  	    updateravailable = excluded.updateravailable,
  	    packagesavailable = excluded.packagesavailable;
`

func (p postgresql) CreateUpdatesData(data pb.Updates, uuid string) error {
	if _, err := p.dbConn.Exec(createUpdatesDataExec,
		uuid, data.InformerActual, data.InformerAvailable, data.UpdaterActual,
		data.UpdaterAvailable, data.PackagesAvailable); err != nil {
		log.Println("CreateUpdatesData", err)
		return err
	}

	return nil
}

const getSecurityDataQuery = `
	SELECT sshattackattempts, securitypackagesavailable
	FROM securitydata
`

func (p postgresql) GetSecurityData() ([]pb.Security, error) {
	rows, err := p.dbConn.Query(getSecurityDataQuery)
	if err != nil {
		log.Fatal("GetSecurityData", err)
		return []pb.Security{}, err
	}
	defer rows.Close()

	securityDates := make([]pb.Security, 0, 10)
	for rows.Next() {
		securityData := pb.Security{}
		if err := rows.Scan(&securityData.SshAttackAttempts,
			&securityData.SecurityPackagesAvailable); err != nil {
			log.Println("securityData: parse err", err)
			continue
		}
		// можно ли так делать?
		securityDates = append(securityDates, securityData)
	}

	return securityDates, nil
}

const createSecurityDataExec = `
	INSERT INTO securitydata
	(uuid, sshattackattempts, securitypackagesavailable)
	VALUES ($1, $2, $3)
	ON CONFLICT (uuid) DO UPDATE 
  	SET sshattackattempts = excluded.sshattackattempts,
  	    securitypackagesavailable = excluded.securitypackagesavailable;
`

func (p postgresql) CreateSecurityData(data pb.Security, uuid string) error {
	if _, err := p.dbConn.Exec(createSecurityDataExec,
		uuid, data.SshAttackAttempts, data.SecurityPackagesAvailable); err != nil {
		log.Println("CreateSecurityData", err)
		return err
	}

	return nil
}

const getStakeInfoDataQuery = `
	SELECT livestake, activestake, pledge
	FROM stackdata
`

func (p postgresql) GetStakeInfoData() ([]pb.StakeInfo, error) {
	rows, err := p.dbConn.Query(getStakeInfoDataQuery)
	if err != nil {
		log.Fatal("GetStakeInfoData", err)
		return []pb.StakeInfo{}, err
	}
	defer rows.Close()

	stakeInfoDates := make([]pb.StakeInfo, 0, 10)
	for rows.Next() {
		stakeInfoData := pb.StakeInfo{}
		if err := rows.Scan(&stakeInfoData.LiveStake,
			&stakeInfoData.ActiveStake, &stakeInfoData.Pledge); err != nil {
			log.Println("stakeInfoData: parse err", err)
			continue
		}
		// можно ли так делать?
		stakeInfoDates = append(stakeInfoDates, stakeInfoData)
	}

	return stakeInfoDates, nil
}

const createStakeInfoDataExec = `
	INSERT INTO stackdata
	(uuid, livestake, activestake, pledge)
	VALUES ($1, $2, $3, $4)
	ON CONFLICT (uuid) DO UPDATE 
  	SET livestake = excluded.livestake,
  	    activestake = excluded.activestake,
  	    pledge = excluded.pledge;
`

func (p postgresql) CreateStakeInfoData(data pb.StakeInfo, uuid string) error {
	if _, err := p.dbConn.Exec(createStakeInfoDataExec,
		uuid, data.LiveStake, data.ActiveStake, data.Pledge); err != nil {
		log.Println("CreateStakeInfoData", err)
		return err
	}

	return nil
}

const getOnlineDataQuery = `
	SELECT sincestart, pings, nodeactive, nodeactivepings, serveractive
	FROM onlinedata
`

func (p postgresql) GetOnlineData() ([]pb.Online, error) {
	rows, err := p.dbConn.Query(getOnlineDataQuery)
	if err != nil {
		log.Fatal("GetOnlineData", err)
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
			log.Println("onlineData: parse err", err)
			continue
		}
		// можно ли так делать?
		onlineDates = append(onlineDates, onlineData)
	}

	return onlineDates, nil
}

const createOnlineDataExec = `
	INSERT INTO onlinedata
	(uuid, sincestart, pings, nodeactive, nodeactivepings, serveractive)
	VALUES ($1, $2, $3, $4, $5, $6)
	ON CONFLICT (uuid) DO UPDATE 
  	SET sincestart = excluded.sincestart,
  	    pings = excluded.pings,
  	    nodeactive = excluded.nodeactive,
  	    nodeactivepings = excluded.nodeactivepings,
  	    serveractive = excluded.serveractive;
`

func (p postgresql) CreateOnlineData(data pb.Online, uuid string) error {
	if _, err := p.dbConn.Exec(createOnlineDataExec,
		uuid, data.SinceStart, data.Pings, data.NodeActive,
		data.NodeActivePings, data.ServerActive); err != nil {
		log.Println("CreateOnlineData", err)
		return err
	}

	return nil
}

const getMemoryStateDataQuery = `
	SELECT total, used, buffers, cached, free, available, active, inactive, 
	       swaptotal, swapused, swapcached, swapfree, memavailableenabled
	FROM memorystatedata
`

func (p postgresql) GetMemoryStateData() ([]pb.MemoryState, error) {
	rows, err := p.dbConn.Query(getMemoryStateDataQuery)
	if err != nil {
		log.Fatal("GetMemoryStateData", err)
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
			log.Println("memoryStateData: parse err", err)
			continue
		}
		// можно ли так делать?
		memoryStateDates = append(memoryStateDates, memoryStateData)
	}

	return memoryStateDates, nil
}

const createMemoryStateDataExec = `
	INSERT INTO memorystatedata
	(uuid, total, used, buffers, cached, free, available, active, inactive,
	 swaptotal, swapused, swapcached, swapfree, memavailableenabled) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
	ON CONFLICT (uuid) DO UPDATE 
  	SET total = excluded.total,
  	    used = excluded.used,
  	    buffers = excluded.buffers,
  	    cached = excluded.cached,
  	    free = excluded.free,
  	    available = excluded.available,
  	    active = excluded.inactive,
  	    swaptotal = excluded.swaptotal,
  	    swapused = excluded.swapused,
  	    swapcached = excluded.swapcached,
  	    swapfree = excluded.swapfree,
  	    memavailableenabled = excluded.memavailableenabled;
`

func (p postgresql) CreateMemoryStateData(data pb.MemoryState, uuid string) error {
	if _, err := p.dbConn.Exec(createMemoryStateDataExec,
		uuid, data.Total, data.Used, data.Buffers,
		data.Cached, data.Free, data.Available,
		data.Active, data.Inactive, data.SwapTotal,
		data.SwapUsed, data.SwapCached, data.SwapFree, data.MemAvailableEnabled); err != nil {
		log.Println("CreateMemoryStateData", err)
		return err
	}

	return nil
}

const getNodePerformanceDataQuery = `
	SELECT processedtx, peersin, peersout
	FROM nodeperformancedata
`

func (p postgresql) GetNodePerformanceData() ([]pb.NodePerformance, error) {
	rows, err := p.dbConn.Query(getNodePerformanceDataQuery)
	if err != nil {
		log.Fatal("GetNodePerformanceData", err)
		return []pb.NodePerformance{}, err
	}
	defer rows.Close()

	nodePerformanceDates := make([]pb.NodePerformance, 0, 10)
	for rows.Next() {
		nodePerformanceData := pb.NodePerformance{}
		if err := rows.Scan(&nodePerformanceData.ProcessedTx,
			&nodePerformanceData.PeersIn, &nodePerformanceData.PeersOut); err != nil {
			log.Println("nodePerformanceData: parse err", err)
			continue
		}
		// можно ли так делать?
		nodePerformanceDates = append(nodePerformanceDates, nodePerformanceData)
	}

	return nodePerformanceDates, nil
}

const createNodePerformanceDataExec = `
	INSERT INTO nodeperformancedata
	(uuid, processedtx, peersin, peersout)
	VALUES ($1, $2, $3, $4)
	ON CONFLICT (uuid) DO UPDATE 
  	SET processedtx = excluded.processedtx,
  	    peersin = excluded.peersin,
  	    peersout = excluded.peersout;
`

func (p postgresql) CreateNodePerformanceData(data pb.NodePerformance, uuid string) error {
	if _, err := p.dbConn.Exec(createNodePerformanceDataExec,
		uuid, data.ProcessedTx, data.PeersIn, data.PeersOut); err != nil {
		log.Println("CreateNodePerformanceData", err)
		return err
	}

	return nil
}

const getCpuStateDataQuery = `
	SELECT cpuqty, averageworkload
	FROM cpustatedata
`

func (p postgresql) GetCpuStateData() ([]pb.CPUState, error) {
	rows, err := p.dbConn.Query(getCpuStateDataQuery)
	if err != nil {
		log.Fatal("GetCpuStateData", err)
		return []pb.CPUState{}, err
	}
	defer rows.Close()

	cpuStateDates := make([]pb.CPUState, 0, 10)
	for rows.Next() {
		cpuStateData := pb.CPUState{}
		if err := rows.Scan(&cpuStateData.CpuQty,
			&cpuStateData.AverageWorkload); err != nil {
			log.Println("cpuStateDates: parse err", err)
			continue
		}
		// можно ли так делать?
		cpuStateDates = append(cpuStateDates, cpuStateData)
	}

	return cpuStateDates, nil
}

const createCpuStateDataExec = `
	INSERT INTO cpustatedata
	(uuid, cpuqty, averageworkload)
	VALUES ($1, $2, $3)
	ON CONFLICT (uuid) DO UPDATE 
  	SET cpuqty = excluded.cpuqty,
  	    averageworkload = excluded.averageworkload;
`

func (p postgresql) CreateCpuStateData(data pb.CPUState, uuid string) error {
	if _, err := p.dbConn.Exec(createCpuStateDataExec,
		uuid, data.CpuQty, data.AverageWorkload); err != nil {
		log.Println("CreateCpuStateData", err)
		return err
	}

	return nil
}

const getNodesStateDataQuery = `
	SELECT tipdiff, density
	FROM nodestatedata
`

func (p postgresql) GetNodeStateData() ([]pb.NodeState, error) {
	rows, err := p.dbConn.Query(getNodesStateDataQuery)
	if err != nil {
		log.Fatal("GetNodeStateData", err)
		return []pb.NodeState{}, err
	}
	defer rows.Close()

	nodeStateDates := make([]pb.NodeState, 0, 10)
	for rows.Next() {
		nodeStateData := pb.NodeState{}
		if err := rows.Scan(&nodeStateData.TipDiff,
			&nodeStateData.Density); err != nil {
			log.Println("nodeStateDates: parse err", err)
			continue
		}
		// можно ли так делать?
		nodeStateDates = append(nodeStateDates, nodeStateData)
	}

	return nodeStateDates, nil
}

const createNodeStateDataExec = `
	INSERT INTO nodestatedata
	(uuid, tipdiff, density) 
	VALUES ($1, $2, $3)
	ON CONFLICT (uuid) DO UPDATE 
  	SET tipdiff = excluded.tipdiff,
  	    density = excluded.density;
`

func (p postgresql) CreateNodeStateData(data pb.NodeState, uuid string) error {
	if _, err := p.dbConn.Exec(createNodeStateDataExec,
		uuid, data.TipDiff, data.Density); err != nil {
		log.Println("CreateNodeStateData", err)
		return err
	}

	return nil
}

const getChiaNodeFarmingDataQuery = `
	SELECT farmingstatus, totalchiafarmed,
	       usertransactionfees, blockrewards,
	       lastheightfarmed, plotcount,
	       totalsizeofplots, estimatednetworkspace,
	       expectedtimetowin
	FROM chianodefarmingdata
`

func (p postgresql) GetChiaNodeFarmingData() ([]pb.ChiaNodeFarming, error) {
	rows, err := p.dbConn.Query(getChiaNodeFarmingDataQuery)
	if err != nil {
		log.Fatal("GetChiaNodeFarmingData", err)
		return []pb.ChiaNodeFarming{}, err
	}
	defer rows.Close()

	chiaNodeFarmingDates := make([]pb.ChiaNodeFarming, 0, 10)
	for rows.Next() {
		chiaNodeFarmingData := pb.ChiaNodeFarming{}
		if err := rows.Scan(&chiaNodeFarmingData.FarmingStatus,
			&chiaNodeFarmingData.TotalSizeOfPlots, &chiaNodeFarmingData.UserTransactionFees,
			&chiaNodeFarmingData.BlockRewards, &chiaNodeFarmingData.LastHeightFarmed,
			&chiaNodeFarmingData.PlotCount, &chiaNodeFarmingData.TotalSizeOfPlots,
			&chiaNodeFarmingData.TotalSizeOfPlots, &chiaNodeFarmingData.EstimatedNetworkSpace,
			&chiaNodeFarmingData.ExpectedTimeToWin); err != nil {
			log.Println("chiaNodeFarmingData: parse err", err)
			continue
		}
		// можно ли так делать?
		chiaNodeFarmingDates = append(chiaNodeFarmingDates, chiaNodeFarmingData)
	}

	return chiaNodeFarmingDates, nil
}

const createChiaNodeFarmingDataExec = `
	INSERT INTO chianodefarmingdata
	(uuid, farmingstatus, totalchiafarmed, usertransactionfees,
	 blockrewards, lastheightfarmed,
	 plotcount, totalsizeofplots,
	 estimatednetworkspace, expectedtimetowin)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	ON CONFLICT (uuid) DO UPDATE 
  	SET farmingstatus = excluded.farmingstatus,
  	    totalchiafarmed = excluded.totalchiafarmed,
  	    usertransactionfees = excluded.usertransactionfees,
  	    blockrewards = excluded.blockrewards,
  	    lastheightfarmed = excluded.blockrewards,
  	    plotcount = excluded.plotcount,
  	    totalsizeofplots = excluded.totalsizeofplots,
  	    estimatednetworkspace = excluded.estimatednetworkspace,
  	    expectedtimetowin = excluded.expectedtimetowin;
`

func (p postgresql) CreateChiaNodeFarmingData(data pb.ChiaNodeFarming, uuid string) error {
	if _, err := p.dbConn.Exec(createChiaNodeFarmingDataExec,
		uuid); err != nil {
		log.Println("CreateChiaNodeFarmingData", err)
		return err
	}

	return nil
}
