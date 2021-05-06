package postgresql

import "log"

const createNodeAuthTableExec = `CREATE TABLE IF NOT EXISTS NodeAuth (
    Ticker text PRIMARY KEY,
    Uuid text PRIMARY KEY,
    Status text not null)`

const createNodeBasicDataTableExec = `CREATE TABLE IF NOT EXISTS NodeBasicData (
    Ticker text PRIMARY KEY,
    Type text not null,
    Location text not null,
    NodeVersion text not null)`

const createServerBasicDataTableExec = `CREATE TABLE IF NOT EXISTS ServerBasicData (
    Ipv4 text not null,
    Ipv6 text not null,
    LinuxName text not null,
    LinuxVersion text not null)`

const createEpochDataTableExec = `CREATE TABLE IF NOT EXISTS EpochData (
    EpochNumber int8 not null)`

const createKesDataTableExec = `CREATE TABLE IF NOT EXISTS KesData (
    KesCurrent int8 not null,
    KesRemaining int8 not null,
    KesExpDate text not null)`

const createBlocksDataTableExec = `CREATE TABLE IF NOT EXISTS BlocksData (
    BlockLeader int8 not null,
    BlockAdopted int8 not null,
    BlockInvalid int8 not null)`

const createUpdatesDataTableExec = `CREATE TABLE IF NOT EXISTS UpdatesData (
    InformerActual text not null,
    InformerAvailable text not null,
    UpdaterActual text not null,
	UpdaterAvailable text not null,
	PackagesAvailable int8 not null)`

const createSecurityDataTableExec = `CREATE TABLE IF NOT EXISTS SecurityData (
    SshAttackAttempts int8 not null,
    SecurityPackagesAvailable int8 not null)`

const createStakeDataTableExec = `CREATE TABLE IF NOT EXISTS StackData (
    LiveStake int8 not null,
    ActiveStake int8 not null,
    Pledge int8 not null)`

const createOnlineDataTableExec = `CREATE TABLE IF NOT EXISTS OnlineData (
    SinceStart int8 not null,
    Pings int8 not null,
    NodeActive bool not null,
	NodeActivePings int8 not null,
	ServerActive bool not null)`

const createMemoryStateDataTableExec = `CREATE TABLE IF NOT EXISTS MemoryStateData (
    Total int8 not null,
    Used int8 not null,
    Buffers int8 not null,
	Cached int8 not null,
	Free int8 not null,
    Available int8 not null,
    Active int8 not null,
	Inactive int8 not null,
	SwapTotal bool not null,
	SwapUsed int8 not null,
    SwapCached int8 not null,
    SwapFree int8 not null,
	MemAvailableEnabled int8 not null)`

const createCpuStateTableExec = `CREATE TABLE IF NOT EXISTS CpuStateData (
    CpuQty int8 not null,
    AverageWorkload float4 not null)`

const createNodeStateTableExec = `CREATE TABLE IF NOT EXISTS NodeStateData (
    TipDiff int8 not null,
    Density float4 not null)`

const createNodePerformanceTableExec = `CREATE TABLE IF NOT EXISTS NodePerformanceData (
    ProcessedTx int8 not null,
    PeersIn int8 not null,
    PeersOut int8 not null)`

func (p postgresql) CreateAllTables() error {
	if _, err := p.dbConn.Exec(createNodeAuthTableExec); err != nil {
		log.Println("CreateNodeAuthTable", err)
		return err
	}

	if _, err := p.dbConn.Exec(createNodeBasicDataTableExec); err != nil {
		log.Println("createNodeBasicDataTableExec", err)
		return err
	}

	if _, err := p.dbConn.Exec(createServerBasicDataTableExec); err != nil {
		log.Println("createNodeBasicDataTableExec", err)
		return err
	}

	if _, err := p.dbConn.Exec(createEpochDataTableExec); err != nil {
		log.Println("createNodeBasicDataTableExec", err)
		return err
	}

	if _, err := p.dbConn.Exec(createKesDataTableExec); err != nil {
		log.Println("createNodeBasicDataTableExec", err)
		return err
	}

	if _, err := p.dbConn.Exec(createBlocksDataTableExec); err != nil {
		log.Println("createNodeBasicDataTableExec", err)
		return err
	}

	if _, err := p.dbConn.Exec(createUpdatesDataTableExec); err != nil {
		log.Println("createNodeBasicDataTableExec", err)
		return err
	}

	if _, err := p.dbConn.Exec(createSecurityDataTableExec); err != nil {
		log.Println("createNodeBasicDataTableExec", err)
		return err
	}

	if _, err := p.dbConn.Exec(createStakeDataTableExec); err != nil {
		log.Println("createNodeBasicDataTableExec", err)
		return err
	}

	if _, err := p.dbConn.Exec(createOnlineDataTableExec); err != nil {
		log.Println("createNodeBasicDataTableExec", err)
		return err
	}

	if _, err := p.dbConn.Exec(createMemoryStateDataTableExec); err != nil {
		log.Println("createNodeBasicDataTableExec", err)
		return err
	}

	if _, err := p.dbConn.Exec(createCpuStateTableExec); err != nil {
		log.Println("createNodeBasicDataTableExec", err)
		return err
	}

	if _, err := p.dbConn.Exec(createNodeStateTableExec); err != nil {
		log.Println("createNodeBasicDataTableExec", err)
		return err
	}

	if _, err := p.dbConn.Exec(createNodePerformanceTableExec); err != nil {
		log.Println("createNodeBasicDataTableExec", err)
		return err
	}

	return nil
}
