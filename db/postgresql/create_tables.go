package postgresql

import "log"

const createNodesTableExec = `
    CREATE TABLE IF NOT EXISTS nodes (
    ticker             varchar(40) not null default '',
    uuid               varchar(40) PRIMARY KEY,
    status             varchar(40) not null default '',
    type               varchar(40) not null default '',
    location           varchar(40) not null default '',
    node_Version       varchar(40) not null default '',
    blockchain         varchar(40) not null default '',
    last_Update        timestamp   without time zone not null)
`

const createServerDataTableExec = `
    CREATE TABLE IF NOT EXISTS node_server_data (
    uuid                             varchar(40) REFERENCES Nodes(uuid), 
    ipv4                             varchar(40) not null default '',
    ipv6                             varchar(40) not null default '',
    linux_Name                       varchar(40) not null default '',
    linux_Version                    varchar(40) not null default '',
    informer_Actual                  varchar(40) not null default '',
    informer_Available               varchar(40) not null default '',
    updater_Actual                   varchar(40) not null default '',
    updater_Available                varchar(40) not null default '',
    packages_Available               bigint      not null default 0,
    sshAttack_Attempts               bigint      not null default 0,
    security_Packages_Available      bigint      not null default 0,
    since_Start                      bigint      not null default 0,
    pings                            bigint      not null default 0,
    node_Active                      bool        not null default false,
    node_Active_Pings                bigint      not null default 0,
    server_Active                    bool        not null default false,
    total                            bigint      not null default 0,
    used                             bigint      not null default 0,
    buffers                          bigint      not null default 0,
    cached                           bigint      not null default 0,
    free                             bigint      not null default 0,
    available                        bigint      not null default 0,
    active                           bigint      not null default 0,
    inactive                         bigint      not null default 0,
    swap_Total                       bigint      not null default 0,
    swap_Used                        bigint      not null default 0,
    swap_Cached                      bigint      not null default 0,
    swap_Free                        bigint      not null default 0,
    mem_Available_Enabled            bool        not null default false,
    cpu_Qty                          bigint      not null default 0,
    average_Workload                 float8      not null default 0,
    last_Update                      timestamp   without time zone not null,
    PRIMARY KEY(uuid, last_update))
`

const createCardanoDataTableExec = `
    CREATE TABLE IF NOT EXISTS cardano_data (
    uuid             varchar(40)    REFERENCES Nodes(uuid), 
    epoch_Number     bigint         not null default 0,
    kes_Current      bigint         not null default 0,
    kes_Remaining    bigint         not null default 0,
    kes_Exp_Date     varchar(40)    not null default '',
    block_Leader     bigint         not null default 0,
    block_Adopted    bigint         not null default 0,
    block_Invalid    bigint         not null default 0,
    live_Stake       bigint         not null default 0,
    active_Stake     bigint         not null default 0,
    pledge           bigint         not null default 0,
    tip_Diff         bigint         not null default 0,
    density          float8         not null default 0,
    processed_Tx     bigint         not null default 0,
    peers_In         bigint         not null default 0,
    peers_Out        bigint         not null default 0,
    last_Update      timestamp      without time zone not null,
    PRIMARY KEY(uuid, last_update))
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
