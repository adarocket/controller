package structs

import (
	cardanoPb "github.com/adarocket/proto/proto-gen/cardano"
	commonPB "github.com/adarocket/proto/proto-gen/common"
)

type Database interface {
	CreateAllTables() error
	NodeAuthData
	NodeBasicData
	ServerBasicData
	Epoch
	KesData
	Blocks
	Updates
	Security
	StakeInfo
	Online
	MemoryState
	NodePerformance
	CpuState
	NodeState
}

type NodeAuthData interface {
	GetNodeAuthData() ([]commonPB.NodeAuthData, error)
	CreateNodeAuthData(data commonPB.NodeAuthData) error
	UpdateNodeAuthData(data commonPB.NodeAuthData) error
	DeleteNodeAuthData(data commonPB.NodeAuthData) error
}

type NodeBasicData interface {
	GetNodeBasicData() ([]commonPB.NodeBasicData, error)
	CreateNodeBasicData(data commonPB.NodeBasicData, uuid string) error
}

type ServerBasicData interface {
	GetServerBasicData() ([]commonPB.ServerBasicData, error)
	CreateServerBasicData(data commonPB.ServerBasicData, uuid string) error
}

type Epoch interface {
	GetEpochData() ([]cardanoPb.Epoch, error)
	CreateEpochData(data cardanoPb.Epoch, uuid string) error
}

type KesData interface {
	GetKesData() ([]cardanoPb.KESData, error)
	CreateKesData(data cardanoPb.KESData, uuid string) error
}

type Blocks interface {
	GetBlocksData() ([]cardanoPb.Blocks, error)
	CreateBlocksData(data cardanoPb.Blocks, uuid string) error
}

type Updates interface {
	GetUpdatesData() ([]commonPB.Updates, error)
	CreateUpdatesData(data commonPB.Updates, uuid string) error
}

type Security interface {
	GetSecurityData() ([]commonPB.Security, error)
	CreateSecurityData(data commonPB.Security, uuid string) error
}

type StakeInfo interface {
	GetStakeInfoData() ([]cardanoPb.StakeInfo, error)
	CreateStakeInfoData(data cardanoPb.StakeInfo, uuid string) error
}

type Online interface {
	GetOnlineData() ([]commonPB.Online, error)
	CreateOnlineData(data commonPB.Online, uuid string) error
}

type MemoryState interface {
	GetMemoryStateData() ([]commonPB.MemoryState, error)
	CreateMemoryStateData(data commonPB.MemoryState, uuid string) error
}

type CpuState interface {
	GetCpuStateData() ([]commonPB.CPUState, error)
	CreateCpuStateData(data commonPB.CPUState, uuid string) error
}

type NodeState interface {
	GetNodeStateData() ([]cardanoPb.NodeState, error)
	CreateNodeStateData(data cardanoPb.NodeState, uuid string) error
}

type NodePerformance interface {
	GetNodePerformanceData() ([]cardanoPb.NodePerformance, error)
	CreateNodePerformanceData(data cardanoPb.NodePerformance, uuid string) error
}
