package postgresql

import "log"

const createNodeAuthTableExec = `
	CREATE TABLE IF NOT EXISTS NodeAuth (
    Ticker varchar(40) not null,
    Uuid varchar(40) PRIMARY KEY,
    Status varchar(40) not null)
`

const createNodeBasicDataTableExec = `
	CREATE TABLE IF NOT EXISTS NodeBasicData (
    Ticker varchar(40) PRIMARY KEY,
    Type varchar(40) not null,
    Location varchar(40) not null,
    NodeVersion varchar(40) not null)
`

const createServerBasicDataTableExec = `
	CREATE TABLE IF NOT EXISTS ServerBasicData (
    Ipv4 varchar(40) not null,
    Ipv6 varchar(40) not null,
    LinuxName varchar(40) not null,
    LinuxVersion varchar(40) not null)
`

const createEpochDataTableExec = `
	CREATE TABLE IF NOT EXISTS EpochData (
    EpochNumber bigint not null)
`

const createKesDataTableExec = `
	CREATE TABLE IF NOT EXISTS KesData (
    KesCurrent bigint not null,
    KesRemaining bigint not null,
    KesExpDate varchar(40) not null)
`

const createBlocksDataTableExec = `
	CREATE TABLE IF NOT EXISTS BlocksData (
    BlockLeader bigint not null,
    BlockAdopted bigint not null,
    BlockInvalid bigint not null)
`

const createUpdatesDataTableExec = `
	CREATE TABLE IF NOT EXISTS UpdatesData (
    InformerActual varchar(40) not null,
    InformerAvailable varchar(40) not null,
    UpdaterActual varchar(40) not null,
	UpdaterAvailable varchar(40) not null,
	PackagesAvailable bigint not null)
`

const createSecurityDataTableExec = `
	CREATE TABLE IF NOT EXISTS SecurityData (
    SshAttackAttempts bigint not null,
    SecurityPackagesAvailable bigint not null)
`

const createStakeDataTableExec = `
	CREATE TABLE IF NOT EXISTS StackData (
    LiveStake bigint not null,
    ActiveStake bigint not null,
    Pledge bigint not null)
`

const createOnlineDataTableExec = `
	CREATE TABLE IF NOT EXISTS OnlineData (
    SinceStart bigint not null,
    Pings bigint not null,
    NodeActive bool not null,
	NodeActivePings bigint not null,
	ServerActive bool not null)
`

const createMemoryStateDataTableExec = `
	CREATE TABLE IF NOT EXISTS MemoryStateData (
    Total bigint not null,
    Used bigint not null,
    Buffers bigint not null,
	Cached bigint not null,
	Free bigint not null,
    Available bigint not null,
    Active bigint not null,
	Inactive bigint not null,
	SwapTotal bool not null,
	SwapUsed bigint not null,
    SwapCached bigint not null,
    SwapFree bigint not null,
	MemAvailableEnabled bigint not null)
`

const createCpuStateTableExec = `
	CREATE TABLE IF NOT EXISTS CpuStateData (
    CpuQty bigint not null,
    AverageWorkload float8 not null)
`

const createNodeStateTableExec = `
	CREATE TABLE IF NOT EXISTS NodeStateData (
    TipDiff bigint not null,
    Density float8 not null)
`

const createNodePerformanceTableExec = `
	CREATE TABLE IF NOT EXISTS NodePerformanceData (
    ProcessedTx bigint not null,
    PeersIn bigint not null,
    PeersOut bigint not null)
`

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
