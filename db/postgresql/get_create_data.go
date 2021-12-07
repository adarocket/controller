package postgresql

import (
	"log"
	"time"

	cardanoPb "github.com/adarocket/proto/proto-gen/cardano"
	commonPB "github.com/adarocket/proto/proto-gen/common"
)

const getNodeAuthQuery = `
	SELECT Ticker, Uuid, Status 
	FROM NodeAuth
`

func (p Postgresql) GetNodeAuthData() ([]commonPB.NodeAuthData, error) {
	rows, err := p.dbConn.Query(getNodeAuthQuery)
	if err != nil {
		log.Fatal("GetNodesAuthData", err)
		return []commonPB.NodeAuthData{}, err
	}
	defer rows.Close()

	nodesAuthData := make([]commonPB.NodeAuthData, 0, 10)
	for rows.Next() {
		nodeAuthData := commonPB.NodeAuthData{}
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
	(Ticker, Uuid, Status, LastUpdate) 
	VALUES ($1, $2, $3, $4)
	ON CONFLICT (uuid) DO UPDATE 
  	SET status = excluded.status;
`

func (p Postgresql) CreateNodeAuthData(data commonPB.NodeAuthData) error {
	if _, err := p.dbConn.Exec(createNodeAuthExec,
		data.Ticker, data.Uuid, data.Status, time.Now()); err != nil {
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

func (p Postgresql) UpdateNodeAuthData(data commonPB.NodeAuthData) error {
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

func (p Postgresql) DeleteNodeAuthData(data commonPB.NodeAuthData) error {
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

func (p Postgresql) GetNodeBasicData() ([]commonPB.NodeBasicData, error) {
	rows, err := p.dbConn.Query(getNodeBasicDataQuery)
	if err != nil {
		log.Fatal("GetNodesAuthData", err)
		return []commonPB.NodeBasicData{}, err
	}
	defer rows.Close()

	nodesBasicData := make([]commonPB.NodeBasicData, 0, 10)
	for rows.Next() {
		nodeBasicData := commonPB.NodeBasicData{}
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
	(uuid, ticker, type, location, nodeversion, lastupdate) 
	VALUES ($1, $2, $3, $4, $5, $6)
	ON CONFLICT (uuid) DO UPDATE 
  	SET ticker = excluded.ticker,
    	type = excluded.type,
  		location = excluded.location,
  	    nodeversion = excluded.nodeversion;
`

func (p Postgresql) CreateNodeBasicData(data commonPB.NodeBasicData, uuid string) error {
	if _, err := p.dbConn.Exec(createNodeBasicDataExec,
		uuid, data.Ticker, data.Type, data.Location, data.NodeVersion, time.Now()); err != nil {
		log.Println("CreateNodeBasicData", err)
		return err
	}

	return nil
}

const getServerBasicData = `
	SELECT ipv4, ipv6, linuxname, linuxversion
	FROM serverbasicdata
`

func (p Postgresql) GetServerBasicData() ([]commonPB.ServerBasicData, error) {
	rows, err := p.dbConn.Query(getServerBasicData)
	if err != nil {
		log.Fatal("GetServerBasicData", err)
		return []commonPB.ServerBasicData{}, err
	}
	defer rows.Close()

	serverBasicDates := make([]commonPB.ServerBasicData, 0, 10)
	for rows.Next() {
		serverBasicData := commonPB.ServerBasicData{}
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
	(uuid, ipv4, ipv6, linuxname, linuxversion, lastupdate) 
	VALUES ($1, $2, $3, $4, $5, $6)
	ON CONFLICT (uuid) DO UPDATE 
  	SET ipv4 = excluded.ipv4, 
      	ipv6 = excluded.ipv6,
  	    linuxname = excluded.linuxname,
  	    linuxversion = excluded.linuxversion;
`

func (p Postgresql) CreateServerBasicData(data commonPB.ServerBasicData, uuid string) error {
	if _, err := p.dbConn.Exec(createServerBasicDataExec,
		uuid, data.Ipv4, data.Ipv6, data.LinuxName, data.LinuxVersion, time.Now()); err != nil {
		log.Println("CreateServerBasicData", err)
		return err
	}

	return nil
}

const getEpochDataQuery = `
	SELECT epochnumber
	FROM epochdata
`

func (p Postgresql) GetEpochData() ([]cardanoPb.Epoch, error) {
	rows, err := p.dbConn.Query(getEpochDataQuery)
	if err != nil {
		log.Fatal("GetEpochData", err)
		return []cardanoPb.Epoch{}, err
	}
	defer rows.Close()

	epochDates := make([]cardanoPb.Epoch, 0, 10)
	for rows.Next() {
		epochData := cardanoPb.Epoch{}
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
	(uuid, epochnumber, lastupdate)
	VALUES ($1, $2, $3)
	ON CONFLICT (uuid) DO UPDATE 
  	SET epochnumber = excluded.epochnumber;
`

func (p Postgresql) CreateEpochData(data cardanoPb.Epoch, uuid string) error {
	if _, err := p.dbConn.Exec(createEpochDataExec,
		uuid, data.EpochNumber, time.Now()); err != nil {
		log.Println("CreateEpochData", err)
		return err
	}

	return nil
}

const getKesDataQuery = `
	SELECT epochnumber
	FROM epochdata
`

func (p Postgresql) GetKesData() ([]cardanoPb.KESData, error) {
	rows, err := p.dbConn.Query(getKesDataQuery)
	if err != nil {
		log.Fatal("GetKesData", err)
		return []cardanoPb.KESData{}, err
	}
	defer rows.Close()

	kesDates := make([]cardanoPb.KESData, 0, 10)
	for rows.Next() {
		kesData := cardanoPb.KESData{}
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
	(uuid, kescurrent, kesremaining, kesexpdate, lastupdate) 
	VALUES ($1, $2, $3, $4, $5)
	ON CONFLICT (uuid) DO UPDATE 
  	SET kescurrent = excluded.kescurrent,
  	    kesremaining = excluded.kesremaining,
  	    kesexpdate = excluded.kesexpdate;  	    
`

func (p Postgresql) CreateKesData(data cardanoPb.KESData, uuid string) error {
	if _, err := p.dbConn.Exec(createKesDataExec,
		uuid, data.KesCurrent, data.KesRemaining, data.KesExpDate, time.Now()); err != nil {
		log.Println("CreateKesData", err)
		return err
	}

	return nil
}

const getBlocksDataQuery = `
	SELECT blockleader, blockadopted, blockinvalid
	FROM blocksdata
`

func (p Postgresql) GetBlocksData() ([]cardanoPb.Blocks, error) {
	rows, err := p.dbConn.Query(getBlocksDataQuery)
	if err != nil {
		log.Fatal("GetBlocksData", err)
		return []cardanoPb.Blocks{}, err
	}
	defer rows.Close()

	blockDates := make([]cardanoPb.Blocks, 0, 10)
	for rows.Next() {
		blockData := cardanoPb.Blocks{}
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
	(uuid, blockleader, blockadopted, blockinvalid, lastupdate)
	VALUES ($1, $2, $3, $4, $5)
	ON CONFLICT (uuid) DO UPDATE 
  	SET blockleader = excluded.blockleader,
  	    blockadopted = excluded.blockadopted,
  	    blockinvalid = excluded.blockinvalid;  
`

func (p Postgresql) CreateBlocksData(data cardanoPb.Blocks, uuid string) error {
	if _, err := p.dbConn.Exec(createBlocksDataExec,
		uuid, data.BlockLeader, data.BlockAdopted, data.BlockInvalid, time.Now()); err != nil {
		log.Println("CreateBlocksData", err)
		return err
	}

	return nil
}

const getUpdatesDataQuery = `
	SELECT INFORMERACTUAL, INFORMERAVAILABLE, UPDATERACTUAL, UPDATERAVAILABLE, PACKAGESAVAILABLE
	FROM updatesdata
`

func (p Postgresql) GetUpdatesData() ([]commonPB.Updates, error) {
	rows, err := p.dbConn.Query(getUpdatesDataQuery)
	if err != nil {
		log.Fatal("GetUpdatesData", err)
		return []commonPB.Updates{}, err
	}
	defer rows.Close()

	updatesDates := make([]commonPB.Updates, 0, 10)
	for rows.Next() {
		updateData := commonPB.Updates{}
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
	(uuid, informeractual, informeravailable, updateractual, updateravailable, packagesavailable, lastupdate)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	ON CONFLICT (uuid) DO UPDATE 
  	SET informeractual = excluded.informeractual,
  	    informeravailable = excluded.informeravailable,
  	    updateractual = excluded.updateractual, 
  	    updateravailable = excluded.updateravailable,
  	    packagesavailable = excluded.packagesavailable;
`

func (p Postgresql) CreateUpdatesData(data commonPB.Updates, uuid string) error {
	if _, err := p.dbConn.Exec(createUpdatesDataExec,
		uuid, data.InformerActual, data.InformerAvailable, data.UpdaterActual,
		data.UpdaterAvailable, data.PackagesAvailable, time.Now()); err != nil {
		log.Println("CreateUpdatesData", err)
		return err
	}

	return nil
}

const getSecurityDataQuery = `
	SELECT sshattackattempts, securitypackagesavailable
	FROM securitydata
`

func (p Postgresql) GetSecurityData() ([]commonPB.Security, error) {
	rows, err := p.dbConn.Query(getSecurityDataQuery)
	if err != nil {
		log.Fatal("GetSecurityData", err)
		return []commonPB.Security{}, err
	}
	defer rows.Close()

	securityDates := make([]commonPB.Security, 0, 10)
	for rows.Next() {
		securityData := commonPB.Security{}
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
	(uuid, sshattackattempts, securitypackagesavailable, lastupdate)
	VALUES ($1, $2, $3, $4)
	ON CONFLICT (uuid) DO UPDATE 
  	SET sshattackattempts = excluded.sshattackattempts,
  	    securitypackagesavailable = excluded.securitypackagesavailable;
`

func (p Postgresql) CreateSecurityData(data commonPB.Security, uuid string) error {
	if _, err := p.dbConn.Exec(createSecurityDataExec,
		uuid, data.SshAttackAttempts, data.SecurityPackagesAvailable, time.Now()); err != nil {
		log.Println("CreateSecurityData", err)
		return err
	}

	return nil
}

const getStakeInfoDataQuery = `
	SELECT livestake, activestake, pledge
	FROM stackdata
`

func (p Postgresql) GetStakeInfoData() ([]cardanoPb.StakeInfo, error) {
	rows, err := p.dbConn.Query(getStakeInfoDataQuery)
	if err != nil {
		log.Fatal("GetStakeInfoData", err)
		return []cardanoPb.StakeInfo{}, err
	}
	defer rows.Close()

	stakeInfoDates := make([]cardanoPb.StakeInfo, 0, 10)
	for rows.Next() {
		stakeInfoData := cardanoPb.StakeInfo{}
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
	(uuid, livestake, activestake, pledge, lastupdate)
	VALUES ($1, $2, $3, $4, $5)
	ON CONFLICT (uuid) DO UPDATE 
  	SET livestake = excluded.livestake,
  	    activestake = excluded.activestake,
  	    pledge = excluded.pledge;
`

func (p Postgresql) CreateStakeInfoData(data cardanoPb.StakeInfo, uuid string) error {
	if _, err := p.dbConn.Exec(createStakeInfoDataExec,
		uuid, data.LiveStake, data.ActiveStake, data.Pledge, time.Now()); err != nil {
		log.Println("CreateStakeInfoData", err)
		return err
	}

	return nil
}

const getOnlineDataQuery = `
	SELECT sincestart, pings, nodeactive, nodeactivepings, serveractive
	FROM onlinedata
`

func (p Postgresql) GetOnlineData() ([]commonPB.Online, error) {
	rows, err := p.dbConn.Query(getOnlineDataQuery)
	if err != nil {
		log.Fatal("GetOnlineData", err)
		return []commonPB.Online{}, err
	}
	defer rows.Close()

	onlineDates := make([]commonPB.Online, 0, 10)
	for rows.Next() {
		onlineData := commonPB.Online{}
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
	(uuid, sincestart, pings, nodeactive, nodeactivepings, serveractive, lastupdate)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	ON CONFLICT (uuid) DO UPDATE 
  	SET sincestart = excluded.sincestart,
  	    pings = excluded.pings,
  	    nodeactive = excluded.nodeactive,
  	    nodeactivepings = excluded.nodeactivepings,
  	    serveractive = excluded.serveractive;
`

func (p Postgresql) CreateOnlineData(data commonPB.Online, uuid string) error {
	if _, err := p.dbConn.Exec(createOnlineDataExec,
		uuid, data.SinceStart, data.Pings, data.NodeActive,
		data.NodeActivePings, data.ServerActive, time.Now()); err != nil {
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

func (p Postgresql) GetMemoryStateData() ([]commonPB.MemoryState, error) {
	rows, err := p.dbConn.Query(getMemoryStateDataQuery)
	if err != nil {
		log.Fatal("GetMemoryStateData", err)
		return []commonPB.MemoryState{}, err
	}
	defer rows.Close()

	memoryStateDates := make([]commonPB.MemoryState, 0, 10)
	for rows.Next() {
		memoryStateData := commonPB.MemoryState{}
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
	 swaptotal, swapused, swapcached, swapfree, memavailableenabled, lastupdate) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
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

func (p Postgresql) CreateMemoryStateData(data commonPB.MemoryState, uuid string) error {
	if _, err := p.dbConn.Exec(createMemoryStateDataExec,
		uuid, data.Total, data.Used, data.Buffers,
		data.Cached, data.Free, data.Available,
		data.Active, data.Inactive, data.SwapTotal,
		data.SwapUsed, data.SwapCached, data.SwapFree,
		data.MemAvailableEnabled, time.Now()); err != nil {
		log.Println("CreateMemoryStateData", err)
		return err
	}

	return nil
}

const getNodePerformanceDataQuery = `
	SELECT processedtx, peersin, peersout
	FROM nodeperformancedata
`

func (p Postgresql) GetNodePerformanceData() ([]cardanoPb.NodePerformance, error) {
	rows, err := p.dbConn.Query(getNodePerformanceDataQuery)
	if err != nil {
		log.Fatal("GetNodePerformanceData", err)
		return []cardanoPb.NodePerformance{}, err
	}
	defer rows.Close()

	nodePerformanceDates := make([]cardanoPb.NodePerformance, 0, 10)
	for rows.Next() {
		nodePerformanceData := cardanoPb.NodePerformance{}
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
	(uuid, processedtx, peersin, peersout, lastupdate)
	VALUES ($1, $2, $3, $4, $5)
	ON CONFLICT (uuid) DO UPDATE 
  	SET processedtx = excluded.processedtx,
  	    peersin = excluded.peersin,
  	    peersout = excluded.peersout;
`

func (p Postgresql) CreateNodePerformanceData(data cardanoPb.NodePerformance, uuid string) error {
	if _, err := p.dbConn.Exec(createNodePerformanceDataExec,
		uuid, data.ProcessedTx, data.PeersIn, data.PeersOut, time.Now()); err != nil {
		log.Println("CreateNodePerformanceData", err)
		return err
	}

	return nil
}

const getCpuStateDataQuery = `
	SELECT cpuqty, averageworkload
	FROM cpustatedata
`

func (p Postgresql) GetCpuStateData() ([]commonPB.CPUState, error) {
	rows, err := p.dbConn.Query(getCpuStateDataQuery)
	if err != nil {
		log.Fatal("GetCpuStateData", err)
		return []commonPB.CPUState{}, err
	}
	defer rows.Close()

	cpuStateDates := make([]commonPB.CPUState, 0, 10)
	for rows.Next() {
		cpuStateData := commonPB.CPUState{}
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
	(uuid, cpuqty, averageworkload, lastupdate)
	VALUES ($1, $2, $3, $4)
	ON CONFLICT (uuid) DO UPDATE 
  	SET cpuqty = excluded.cpuqty,
  	    averageworkload = excluded.averageworkload;
`

func (p Postgresql) CreateCpuStateData(data commonPB.CPUState, uuid string) error {
	if _, err := p.dbConn.Exec(createCpuStateDataExec,
		uuid, data.CpuQty, data.AverageWorkload, time.Now()); err != nil {
		log.Println("CreateCpuStateData", err)
		return err
	}

	return nil
}

const getNodesStateDataQuery = `
	SELECT tipdiff, density
	FROM nodestatedata
`

func (p Postgresql) GetNodeStateData() ([]cardanoPb.NodeState, error) {
	rows, err := p.dbConn.Query(getNodesStateDataQuery)
	if err != nil {
		log.Fatal("GetNodeStateData", err)
		return []cardanoPb.NodeState{}, err
	}
	defer rows.Close()

	nodeStateDates := make([]cardanoPb.NodeState, 0, 10)
	for rows.Next() {
		nodeStateData := cardanoPb.NodeState{}
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
	(uuid, tipdiff, density, lastupdate) 
	VALUES ($1, $2, $3, $4)
	ON CONFLICT (uuid) DO UPDATE 
  	SET tipdiff = excluded.tipdiff,
  	    density = excluded.density;
`

func (p Postgresql) CreateNodeStateData(data cardanoPb.NodeState, uuid string) error {
	if _, err := p.dbConn.Exec(createNodeStateDataExec,
		uuid, data.TipDiff, data.Density, time.Now()); err != nil {
		log.Println("CreateNodeStateData", err)
		return err
	}

	return nil
}
