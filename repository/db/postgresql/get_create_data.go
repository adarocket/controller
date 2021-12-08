package postgresql

import (
	"github.com/adarocket/controller/repository/db/structs"
	"log"
	"time"
)

const getNodesDataQuery = `
	SELECT ticker, uuid, status, type, 
	       location, nodeversion, lastupdate 
	FROM Nodes
`

func (p Postgresql) GetNodesData() ([]structs.Node, error) {
	rows, err := p.dbConn.Query(getNodesDataQuery)
	if err != nil {
		log.Fatal("GetNodesAuthData", err)
		return []structs.Node{}, err
	}
	defer rows.Close()

	nodesData := make([]structs.Node, 0, 10)
	for rows.Next() {
		data := structs.Node{}
		if err := rows.Scan(&data.NodeAuthData.Ticker, &data.Uuid, &data.Status,
			&data.Type, &data.Location, &data.NodeVersion, time.Now()); err != nil {
			log.Println("NodesAuth: parse err", err)
			continue
		}

		nodesData = append(nodesData, data)
	}

	return nodesData, nil
}

const createNodeExec = `
	INSERT INTO Nodes
	(TICKER, UUID, STATUS, TYPE, LOCATION, NODEVERSION, LASTUPDATE) 
	VALUES ($1, $2, $3, $4, $5, $6, $7)
`

func (p Postgresql) CreateNodeData(data structs.Node) error {
	if _, err := p.dbConn.Exec(createNodeExec,
		data.NodeAuthData.Ticker, data.Uuid, data.Status,
		data.Type, data.Location, data.NodeVersion, time.Now()); err != nil {
		log.Println("CreateNode", err)
		return err
	}

	return nil
}

const getNodeServerDataQuery = `
	SELECT uuid, ipv4, ipv6, linuxname, linuxversion, informeractual, informeravailable,
	       updateractual, updateravailable, packagesavailable, sshattackattempts,
	       securitypackagesavailable, sincestart, pings, nodeactive, nodeactivepings,
	       serveractive, total, used, buffers, cached, free, available, active, inactive,
	       swaptotal, swapused, swapcached, swapfree, memavailableenabled,
	       cpuqty, averageworkload
	FROM serverdata
`

func (p Postgresql) GetNodeServerData() ([]structs.ServerData, error) {
	rows, err := p.dbConn.Query(getNodeServerDataQuery)
	if err != nil {
		log.Fatal("GetNodesServerData", err)
		return []structs.ServerData{}, err
	}
	defer rows.Close()

	serverData := make([]structs.ServerData, 0, 10)
	for rows.Next() {
		data := structs.ServerData{}
		if err := rows.Scan(&data.Uuid, &data.Ipv4, data.Ipv6, &data.LinuxName, &data.LinuxVersion,
			&data.InformerActual, &data.InformerAvailable, &data.UpdaterActual,
			&data.UpdaterAvailable, &data.PackagesAvailable, &data.SshAttackAttempts,
			&data.SecurityPackagesAvailable, &data.SinceStart, &data.Pings,
			&data.NodeActive, &data.NodeActivePings, &data.ServerActive,
			&data.Total, &data.Used, &data.Buffers, &data.Cached, &data.Free,
			&data.Available, &data.Active, &data.Inactive, &data.SwapTotal, &data.SwapUsed,
			&data.SwapCached, &data.SwapFree, &data.MemAvailableEnabled, &data.CpuQty,
			&data.AverageWorkload); err != nil {
			log.Println("NodesAuth: parse err", err)
			continue
		}

		serverData = append(serverData, data)
	}

	return serverData, nil
}

const createNodeServerDataExec = `
	INSERT INTO serverdata 
	(uuid, ipv4, ipv6, linuxname, linuxversion, informeractual,
	 informeravailable, updateractual, updateravailable, 
	 packagesavailable, sshattackattempts, securitypackagesavailable,
	 sincestart, pings, nodeactive, nodeactivepings, serveractive,
	 total, used, buffers, cached, free, available, active, 
	 inactive, swaptotal, swapused, swapcached, swapfree,
	 memavailableenabled, cpuqty, averageworkload, lastupdate) 
	VALUES ($1, $2, $3, $4, $5, $6,
	        $7, $8, $9, $10, $11, $12,
	        $13, $14, $15, $16, $17, $18,
	        $19, $20, $21, $22, $23, $24,
	        $25, $26, $27, $28, $29, $30,
	        $31, $32, $33)
`

func (p Postgresql) CreateNodeServerData(data structs.ServerData) error {
	if _, err := p.dbConn.Exec(createNodeServerDataExec,
		data.Uuid, data.Ipv4, data.Ipv6, data.LinuxName, data.LinuxVersion,
		data.InformerActual, data.InformerAvailable, data.UpdaterActual,
		data.UpdaterAvailable, data.PackagesAvailable, data.SshAttackAttempts,
		data.SecurityPackagesAvailable, data.SinceStart, data.Pings,
		data.NodeActive, data.NodeActivePings, data.ServerActive,
		data.Total, data.Used, data.Buffers, data.Cached, data.Free,
		data.Available, data.Active, data.Inactive, data.SwapTotal, data.SwapUsed,
		data.SwapCached, data.SwapFree, data.MemAvailableEnabled, data.CpuQty,
		data.AverageWorkload, time.Now()); err != nil {
		log.Println("CreateNodeServerData", err)
		return err
	}

	return nil
}

const getCardanoData = `
	SELECT uuid, epochnumber, kescurrent, kesremaining,
	       kesexpdate, blockleader, blockadopted, blockinvalid,
	       livestake, activestake, pledge, tipdiff, density, processedtx,
	       peersin, peersout, lastupdate
	FROM cardanodata
`

func (p Postgresql) GetCardanoData() ([]structs.CardanoData, error) {
	rows, err := p.dbConn.Query(getCardanoData)
	if err != nil {
		log.Fatal("GetServerBasicData", err)
		return []structs.CardanoData{}, err
	}
	defer rows.Close()

	cardanoData := make([]structs.CardanoData, 0, 10)
	for rows.Next() {
		data := structs.CardanoData{}
		if err := rows.Scan(&data.Uuid, &data.EpochNumber, &data.KesCurrent, &data.KesRemaining,
			&data.KesExpDate, &data.BlockLeader, &data.BlockAdopted,
			&data.BlockInvalid, &data.LiveStake, &data.ActiveStake,
			&data.Pledge, &data.TipDiff, &data.Density,
			&data.ProcessedTx, &data.PeersIn, &data.PeersOut); err != nil {
			log.Println("serverBasicData: parse err", err)
			continue
		}

		cardanoData = append(cardanoData, data)
	}

	return cardanoData, nil
}

const createCardanoDataExec = `
	INSERT INTO cardanodata
	(uuid, epochnumber, kescurrent, kesremaining,
	 kesexpdate, blockleader, blockadopted, blockinvalid,
	 livestake, activestake, pledge, tipdiff, density,
	 processedtx, peersin, peersout, lastupdate) 
	VALUES ($1, $2, $3, $4, $5, $6,
	        $7, $8, $9, $10, $11, $12,
	        $13, $14, $15, $16, $17)
`

func (p Postgresql) CreateCardanoData(data structs.CardanoData) error {
	if _, err := p.dbConn.Exec(createCardanoDataExec,
		data.Uuid, data.EpochNumber, data.KesCurrent, data.KesRemaining,
		data.KesExpDate, data.BlockLeader, data.BlockAdopted,
		data.BlockInvalid, data.LiveStake, data.ActiveStake,
		data.Pledge, data.TipDiff, data.Density,
		data.ProcessedTx, data.PeersIn, data.PeersOut, time.Now()); err != nil {
		log.Println("CreateServerBasicData", err)
		return err
	}

	return nil
}
