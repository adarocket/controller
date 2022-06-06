package postgresql

import (
	"log"
	"time"

	"github.com/adarocket/controller/db/structs"
)

const getNodesDataQuery = `
	SELECT ticker, uuid, status, name, type, 
	       location, node_version, blockchain, last_update 
	FROM nodes
`

func (p postgresql) GetNodesData() ([]structs.Node, error) {
	rows, err := p.dbConn.Query(getNodesDataQuery)
	if err != nil {
		log.Println("GetNodesAuthData", err)
		return []structs.Node{}, err
	}
	defer rows.Close()

	nodesData := make([]structs.Node, 0, 10)
	for rows.Next() {
		data := structs.Node{}
		err := rows.Scan(
			&data.NodeAuthData.Ticker,
			&data.NodeAuthData.Uuid,
			&data.NodeAuthData.Status,
			&data.NodeAuthData.Name,
			&data.NodeBasicData.Type,
			&data.NodeBasicData.Location,
			&data.NodeBasicData.NodeVersion,
			&data.NodeAuthData.Blockchain,
			&data.LastUpdate)
		if err != nil {
			log.Println("NodesAuth: parse err", err)
			continue
		}

		data.NodeAuthData.Type = data.NodeBasicData.Type

		nodesData = append(nodesData, data)
	}

	return nodesData, nil
}

const getNodeDataQuery = `
	SELECT ticker, uuid, status, name, type, 
	       location, node_version, blockchain, last_update 
	FROM nodes
	WHERE uuid = $1
`

func (p postgresql) GetNodeData(uuid string) (structs.Node, error) {
	rows, err := p.dbConn.Query(getNodeDataQuery, uuid)
	if err != nil {
		log.Println("GetNodesAuthData", err)
		return structs.Node{}, err
	}
	defer rows.Close()

	data := structs.Node{}
	for rows.Next() {
		err := rows.Scan(&data.NodeAuthData.Ticker,
			&data.NodeAuthData.Uuid,
			&data.NodeAuthData.Status,
			&data.NodeAuthData.Name,
			&data.NodeBasicData.Type,
			&data.NodeBasicData.Location,
			&data.NodeBasicData.NodeVersion,
			&data.NodeAuthData.Blockchain,
			&data.LastUpdate)
		if err != nil {
			log.Println("NodesAuth: parse err", err)
			continue
		}
		data.NodeAuthData.Type = data.NodeBasicData.Type
	}

	return data, nil
}

const createNodeExec = `
	INSERT INTO nodes (
		ticker, 
		uuid, 
		status, 
		name,
		type, 
		location,
		node_version, 
		blockchain, 
		last_update) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	ON CONFLICT (uuid) DO UPDATE 
  	SET ticker 			= excluded.ticker,
  	    status 			= excluded.status,
  	    name 			= excluded.name,
  	    type 			= excluded.type,
  	    location 		= excluded.location,
  	    node_version 	= excluded.node_version,
  	    blockchain 		= excluded.blockchain,
  	    last_update 	= excluded.last_update;
`

func (p postgresql) CreateNodeData(data structs.Node) error {
	_, err := p.dbConn.Exec(createNodeExec,
		data.NodeAuthData.Ticker,
		data.NodeAuthData.Uuid,
		data.NodeAuthData.Status,
		data.NodeAuthData.Name,
		data.NodeBasicData.Type,
		data.NodeBasicData.Location,
		data.NodeBasicData.NodeVersion,
		data.NodeAuthData.Blockchain,
		time.Now())
	if err != nil {
		log.Println("CreateNode", err)
		return err
	}

	return nil
}

const getNodeServerDataQuery = `
	SELECT uuid, ipv4, ipv6, linux_name, linux_version,
	       informer_actual, informer_available, updater_actual,
	       updater_available, packages_available, sshattack_attempts,
	       security_packages_available, since_start, pings, node_active,
	       node_active_pings, server_active, total, used, buffers, cached,
	       free, available, active, inactive, swap_total, swap_used, swap_cached,
	       swap_free, mem_available_enabled, cpu_qty, average_workload, last_update
	FROM node_server_data
	WHERE uuid = $1
    ORDER BY "last_update" DESC
    LIMIT 100
`

func (p postgresql) GetNodeServerData(uuid string) ([]structs.ServerData, error) {
	rows, err := p.dbConn.Query(getNodeServerDataQuery, uuid)
	if err != nil {
		log.Println("GetNodesServerData", err)
		return []structs.ServerData{}, err
	}
	defer rows.Close()

	serverData := make([]structs.ServerData, 0, 50)
	for rows.Next() {
		data := structs.ServerData{}
		err := rows.Scan(&data.Uuid, &data.Ipv4, data.Ipv6, &data.LinuxName, &data.LinuxVersion,
			&data.InformerActual, &data.InformerAvailable, &data.UpdaterActual,
			&data.UpdaterAvailable, &data.PackagesAvailable, &data.SshAttackAttempts,
			&data.SecurityPackagesAvailable, &data.SinceStart, &data.Pings,
			&data.NodeActive, &data.NodeActivePings, &data.ServerActive,
			&data.Total, &data.Used, &data.Buffers, &data.Cached, &data.Free,
			&data.Available, &data.Active, &data.Inactive, &data.SwapTotal, &data.SwapUsed,
			&data.SwapCached, &data.SwapFree, &data.MemAvailableEnabled, &data.CpuQty,
			&data.AverageWorkload, &data.LastUpdate)
		if err != nil {
			log.Println("NodesAuth: parse err", err)
			continue
		}

		serverData = append(serverData, data)
	}

	return serverData, nil
}

const createNodeServerDataExec = `
	INSERT INTO node_server_data 
	(uuid, ipv4, ipv6, linux_name, linux_version,
	 informer_actual, informer_available, updater_actual,
	 updater_available, packages_available, sshattack_attempts,
	 security_packages_available, since_start, pings, node_active,
	 node_active_pings, server_active, total, used, buffers, cached,
	 free, available, active, inactive, swap_total, swap_used, swap_cached,
	 swap_free, mem_available_enabled, cpu_qty, average_workload, last_update) 
	VALUES ($1, $2, $3, $4, $5, $6,
	        $7, $8, $9, $10, $11, $12,
	        $13, $14, $15, $16, $17, $18,
	        $19, $20, $21, $22, $23, $24,
	        $25, $26, $27, $28, $29, $30,
	        $31, $32, $33)
`

func (p postgresql) CreateNodeServerData(data structs.ServerData) error {
	_, err := p.dbConn.Exec(createNodeServerDataExec,
		data.Uuid, data.Ipv4, data.Ipv6, data.LinuxName, data.LinuxVersion,
		data.InformerActual, data.InformerAvailable, data.UpdaterActual,
		data.UpdaterAvailable, data.PackagesAvailable, data.SshAttackAttempts,
		data.SecurityPackagesAvailable, data.SinceStart, data.Pings,
		data.NodeActive, data.NodeActivePings, data.ServerActive,
		data.Total, data.Used, data.Buffers, data.Cached, data.Free,
		data.Available, data.Active, data.Inactive, data.SwapTotal, data.SwapUsed,
		data.SwapCached, data.SwapFree, data.MemAvailableEnabled, data.CpuQty,
		data.AverageWorkload, time.Now())
	if err != nil {
		log.Println("CreateNodeServerData", err)
		return err
	}

	return nil
}

const getCardanoData = `
	SELECT uuid, epoch_number, kes_current, kes_remaining,
	       kes_exp_date, block_leader, block_adopted, block_invalid,
	       live_stake, active_stake, pledge, tip_diff, density, processed_tx,
	       peers_in, peers_out, last_update
	FROM cardano_data
	WHERE uuid = $1
    ORDER BY "last_update" DESC
    LIMIT 100
`

func (p postgresql) GetCardanoData(uuid string) ([]structs.CardanoData, error) {
	rows, err := p.dbConn.Query(getCardanoData, uuid)
	if err != nil {
		log.Println("GetServerBasicData", err)
		return []structs.CardanoData{}, err
	}
	defer rows.Close()

	cardanoData := make([]structs.CardanoData, 0, 50)
	for rows.Next() {
		data := structs.CardanoData{}
		err := rows.Scan(&data.Uuid, &data.EpochNumber, &data.KesCurrent, &data.KesRemaining,
			&data.KesExpDate, &data.BlockLeader, &data.BlockAdopted,
			&data.BlockInvalid, &data.LiveStake, &data.ActiveStake,
			&data.Pledge, &data.TipDiff, &data.Density,
			&data.ProcessedTx, &data.PeersIn, &data.PeersOut, &data.LastUpdate)
		if err != nil {
			log.Println("serverBasicData: parse err", err)
			continue
		}

		cardanoData = append(cardanoData, data)
	}

	return cardanoData, nil
}

const createCardanoDataExec = `
	INSERT INTO cardano_data
	(uuid, epoch_number, kes_current, kes_remaining,
	 kes_exp_date, block_leader, block_adopted, block_invalid,
	 live_stake, active_stake, pledge, tip_diff, density, processed_tx,
	 peers_in, peers_out, last_update) 
	VALUES ($1, $2, $3, $4, $5, $6,
	        $7, $8, $9, $10, $11, $12,
	        $13, $14, $15, $16, $17)
`

func (p postgresql) CreateCardanoData(data structs.CardanoData) error {
	_, err := p.dbConn.Exec(createCardanoDataExec,
		data.Uuid, data.EpochNumber, data.KesCurrent, data.KesRemaining,
		data.KesExpDate, data.BlockLeader, data.BlockAdopted,
		data.BlockInvalid, data.LiveStake, data.ActiveStake,
		data.Pledge, data.TipDiff, data.Density,
		data.ProcessedTx, data.PeersIn, data.PeersOut, time.Now())
	if err != nil {
		log.Println("CreateServerBasicData", err)
		return err
	}

	return nil
}
