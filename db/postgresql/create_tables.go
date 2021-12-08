package postgresql

import "log"

const createNodesTableExec = `
	CREATE TABLE IF NOT EXISTS Nodes (
    Ticker varchar(40) not null default '',
    Uuid varchar(40) PRIMARY KEY,
    Status varchar(40) not null default '',
	Type varchar(40) not null default '',
    Location varchar(40) not null default '',
    NodeVersion varchar(40) not null default '',
	LastUpdate timestamp without time zone not null)
`

const createServerDataTableExec = `
	CREATE TABLE IF NOT EXISTS ServerData (
	Uuid varchar(40) REFERENCES Nodes(uuid), 
    Ipv4 varchar(40) not null default '',
    Ipv6 varchar(40) not null default '',
    LinuxName varchar(40) not null default '',
    LinuxVersion varchar(40) not null default '',
	InformerActual varchar(40) not null default '',
    InformerAvailable varchar(40) not null default '',
    UpdaterActual varchar(40) not null default '',
	UpdaterAvailable varchar(40) not null default '',
	PackagesAvailable bigint not null default 0,
	SshAttackAttempts bigint not null default 0,
    SecurityPackagesAvailable bigint not null default 0,
	SinceStart bigint not null default 0,
    Pings bigint not null default 0,
    NodeActive bool not null default false,
	NodeActivePings bigint not null default 0,
	ServerActive bool not null default false,
	Total bigint not null default 0,
    Used bigint not null default 0,
    Buffers bigint not null default 0,
	Cached bigint not null default 0,
	Free bigint not null default 0,
    Available bigint not null default 0,
    Active bigint not null default 0,
	Inactive bigint not null default 0,
	SwapTotal bigint not null default 0,
	SwapUsed bigint not null default 0,
    SwapCached bigint not null default 0,
    SwapFree bigint not null default 0,
	MemAvailableEnabled bool not null default false,
	CpuQty bigint not null default 0,
    AverageWorkload float8 not null default 0,
	LastUpdate timestamp without time zone not null,
	PRIMARY KEY(uuid, lastupdate))
`

const createCardanoDataTableExec = `
	CREATE TABLE IF NOT EXISTS CardanoData (
    Uuid varchar(40) REFERENCES Nodes(uuid), 
	EpochNumber bigint not null default 0,
	KesCurrent bigint not null default 0,
    KesRemaining bigint not null default 0,
    KesExpDate varchar(40) not null default '',
	BlockLeader bigint not null default 0,
    BlockAdopted bigint not null default 0,
    BlockInvalid bigint not null default 0,
	LiveStake bigint not null default 0,
    ActiveStake bigint not null default 0,
    Pledge bigint not null default 0,
	TipDiff bigint not null default 0,
    Density float8 not null default 0,
	ProcessedTx bigint not null default 0,
    PeersIn bigint not null default 0,
    PeersOut bigint not null default 0,
	LastUpdate timestamp without time zone not null,
	PRIMARY KEY(uuid, lastupdate))
`

func (p Postgresql) CreateAllTables() error {
	if _, err := p.dbConn.Exec(createNodesTableExec); err != nil {
		log.Println("createNodesTableExec", err)
		return err
	}

	if _, err := p.dbConn.Exec(createServerDataTableExec); err != nil {
		log.Println("createServerDataTableExec", err)
		return err
	}

	if _, err := p.dbConn.Exec(createCardanoDataTableExec); err != nil {
		log.Println("createCardanoDataTableExec", err)
		return err
	}

	return nil
}
