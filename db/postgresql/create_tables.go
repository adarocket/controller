package postgresql

import "log"

const createNodesTableExec = `
    CREATE TABLE IF NOT EXISTS nodes (
    Ticker             varchar(40) not null default '',
    Uuid               varchar(40) PRIMARY KEY,
    Status             varchar(40) not null default '',
    Type               varchar(40) not null default '',
    Location           varchar(40) not null default '',
    Node_Version       varchar(40) not null default '',
    Blockchain         varchar(40) not null default '',
    Last_Update        timestamp   without time zone not null)
`

const createServerDataTableExec = `
    CREATE TABLE IF NOT EXISTS node_server_data (
    Uuid                             varchar(40) REFERENCES Nodes(uuid), 
    Ipv4                             varchar(40) not null default '',
    Ipv6                             varchar(40) not null default '',
    Linux_Name                       varchar(40) not null default '',
    Linux_Version                    varchar(40) not null default '',
    Informer_Actual                  varchar(40) not null default '',
    Informer_Available               varchar(40) not null default '',
    Updater_Actual                   varchar(40) not null default '',
    Updater_Available                varchar(40) not null default '',
    Packages_Available               bigint      not null default 0,
    SshAttack_Attempts               bigint      not null default 0,
    Security_Packages_Available      bigint      not null default 0,
    Since_Start                      bigint      not null default 0,
    Pings                            bigint      not null default 0,
    Node_Active                      bool        not null default false,
    Node_Active_Pings                bigint      not null default 0,
    Server_Active                    bool        not null default false,
    Total                            bigint      not null default 0,
    Used                             bigint      not null default 0,
    Buffers                          bigint      not null default 0,
    Cached                           bigint      not null default 0,
    Free                             bigint      not null default 0,
    Available                        bigint      not null default 0,
    Active                           bigint      not null default 0,
    Inactive                         bigint      not null default 0,
    Swap_Total                       bigint      not null default 0,
    Swap_Used                        bigint      not null default 0,
    Swap_Cached                      bigint      not null default 0,
    Swap_Free                        bigint      not null default 0,
    Mem_Available_Enabled            bool        not null default false,
    Cpu_Qty                          bigint      not null default 0,
    Average_Workload                 float8      not null default 0,
    Last_Update                      timestamp   without time zone not null,
    PRIMARY KEY(uuid, last_update))
`

const createCardanoDataTableExec = `
    CREATE TABLE IF NOT EXISTS cardano_data (
    Uuid             varchar(40)    REFERENCES Nodes(uuid), 
    Epoch_Number     bigint         not null default 0,
    Kes_Current      bigint         not null default 0,
    Kes_Remaining    bigint         not null default 0,
    Kes_Exp_Date     varchar(40)    not null default '',
    Block_Leader     bigint         not null default 0,
    Block_Adopted    bigint         not null default 0,
    Block_Invalid    bigint         not null default 0,
    Live_Stake       bigint         not null default 0,
    Active_Stake     bigint         not null default 0,
    Pledge           bigint         not null default 0,
    Tip_Diff         bigint         not null default 0,
    Density          float8         not null default 0,
    Processed_Tx     bigint         not null default 0,
    Peers_In         bigint         not null default 0,
    Peers_Out        bigint         not null default 0,
    Last_Update      timestamp      without time zone not null,
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
