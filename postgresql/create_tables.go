package postgresql

import "log"

const createNodeAuthTableExec = `
	CREATE TABLE IF NOT EXISTS NodeAuth (
    Ticker varchar(40) not null default '',
    Uuid varchar(40) PRIMARY KEY,
    Status varchar(40) not null default '')
`

const createNodeBasicDataTableExec = `
	CREATE TABLE IF NOT EXISTS NodeBasicData (
	Uuid varchar(40) PRIMARY KEY REFERENCES NodeAuth(uuid),    
    Ticker varchar(40) not null default '',
    Type varchar(40) not null default '',
    Location varchar(40) not null default '',
    NodeVersion varchar(40) not null default '')
`

const createServerBasicDataTableExec = `
	CREATE TABLE IF NOT EXISTS ServerBasicData (
	Uuid varchar(40) PRIMARY KEY REFERENCES NodeAuth(uuid), 
    Ipv4 varchar(40) not null default '',
    Ipv6 varchar(40) not null default '',
    LinuxName varchar(40) not null default '',
    LinuxVersion varchar(40) not null default '')
`

const createEpochDataTableExec = `
	CREATE TABLE IF NOT EXISTS EpochData (
    Uuid varchar(40) PRIMARY KEY REFERENCES NodeAuth(uuid), 
	EpochNumber bigint not null default 0)
`

const createKesDataTableExec = `
	CREATE TABLE IF NOT EXISTS KesData (
    KesCurrent bigint not null default 0,
    KesRemaining bigint not null default 0,
    KesExpDate varchar(40) not null default '')
`

const createBlocksDataTableExec = `
	CREATE TABLE IF NOT EXISTS BlocksData (
	Uuid varchar(40) PRIMARY KEY REFERENCES NodeAuth(uuid), 
    BlockLeader bigint not null default 0,
    BlockAdopted bigint not null default 0,
    BlockInvalid bigint not null default 0)
`

const createUpdatesDataTableExec = `
	CREATE TABLE IF NOT EXISTS UpdatesData (
	Uuid varchar(40) PRIMARY KEY REFERENCES NodeAuth(uuid), 
    InformerActual varchar(40) not null default '',
    InformerAvailable varchar(40) not null default '',
    UpdaterActual varchar(40) not null default '',
	UpdaterAvailable varchar(40) not null default '',
	PackagesAvailable bigint not null default 0)
`

const createSecurityDataTableExec = `
	CREATE TABLE IF NOT EXISTS SecurityData (
	Uuid varchar(40) PRIMARY KEY REFERENCES NodeAuth(uuid), 
    SshAttackAttempts bigint not null default 0,
    SecurityPackagesAvailable bigint not null default 0)
`

const createStakeDataTableExec = `
	CREATE TABLE IF NOT EXISTS StackData (
	Uuid varchar(40) PRIMARY KEY REFERENCES NodeAuth(uuid), 
    LiveStake bigint not null default 0,
    ActiveStake bigint not null default 0,
    Pledge bigint not null default 0)
`

const createOnlineDataTableExec = `
	CREATE TABLE IF NOT EXISTS OnlineData (
	Uuid varchar(40) PRIMARY KEY REFERENCES NodeAuth(uuid), 
    SinceStart bigint not null default 0,
    Pings bigint not null default 0,
    NodeActive bool not null default false,
	NodeActivePings bigint not null default 0,
	ServerActive bool not null default false)
`

const createMemoryStateDataTableExec = `
	CREATE TABLE IF NOT EXISTS MemoryStateData (
	Uuid varchar(40) PRIMARY KEY REFERENCES NodeAuth(uuid), 
    Total bigint not null default 0,
    Used bigint not null default 0,
    Buffers bigint not null default 0,
	Cached bigint not null default 0,
	Free bigint not null default 0,
    Available bigint not null default 0,
    Active bigint not null default 0,
	Inactive bigint not null default 0,
	SwapTotal bool not null default false,
	SwapUsed bigint not null default 0,
    SwapCached bigint not null default 0,
    SwapFree bigint not null default 0,
	MemAvailableEnabled bigint not null default 0)
`

const createCpuStateTableExec = `
	CREATE TABLE IF NOT EXISTS CpuStateData (
	Uuid varchar(40) PRIMARY KEY REFERENCES NodeAuth(uuid), 
    CpuQty bigint not null default 0,
    AverageWorkload float8 not null default 0)
`

const createNodeStateTableExec = `
	CREATE TABLE IF NOT EXISTS NodeStateData (
	Uuid varchar(40) PRIMARY KEY REFERENCES NodeAuth(uuid), 
    TipDiff bigint not null default 0,
    Density float8 not null default 0)
`

const createNodePerformanceTableExec = `
	CREATE TABLE IF NOT EXISTS NodePerformanceData (
	Uuid varchar(40) PRIMARY KEY REFERENCES NodeAuth(uuid), 
    ProcessedTx bigint not null default 0,
    PeersIn bigint not null default 0,
    PeersOut bigint not null default 0)
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
