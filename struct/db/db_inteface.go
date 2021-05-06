package db

import pb "github.com/adarocket/proto"

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
	GetNodeAuthData() ([]pb.NodeAuthData, error)
	CreateNodeAuthData(data pb.NodeAuthData) error
	UpdateNodeAuthData(data pb.NodeAuthData) error
	DeleteNodeAuthData(data pb.NodeAuthData) error
}

type NodeBasicData interface {
	GetNodeBasicData() ([]pb.NodeBasicData, error)
	CreateNodeBasicData(data pb.NodeBasicData) error
}

type ServerBasicData interface {
	GetServerBasicData() ([]pb.ServerBasicData, error)
	CreateServerBasicData(data pb.ServerBasicData) error
}

type Epoch interface {
	GetEpochData() ([]pb.Epoch, error)
	CreateEpochData(data pb.Epoch) error
}

type KesData interface {
	GetKesData() ([]pb.KESData, error)
	CreateKesData(data pb.KESData) error
}

type Blocks interface {
	GetBlocksData() ([]pb.Blocks, error)
	CreateBlocksData(data pb.Blocks) error
}

type Updates interface {
	GetUpdatesData() ([]pb.Updates, error)
	CreateUpdatesData(data pb.Updates) error
}

type Security interface {
	GetSecurityData() ([]pb.Security, error)
	CreateSecurityData(data pb.Security) error
}

type StakeInfo interface {
	GetStakeInfoData() ([]pb.StakeInfo, error)
	CreateStakeInfoData(data pb.StakeInfo) error
}

type Online interface {
	GetOnlineData() ([]pb.Online, error)
	CreateOnlineData(data pb.Online) error
}

type MemoryState interface {
	GetMemoryStateData() ([]pb.MemoryState, error)
	CreateMemoryStateData(data pb.MemoryState) error
}

type CpuState interface {
	GetCpuStateData() ([]pb.CPUState, error)
	CreateCpuStateData(data pb.CPUState) error
}

type NodeState interface {
	GetNodeStateData() ([]pb.NodeState, error)
	CreateNodeStateData(data pb.NodeState) error
}

type NodePerformance interface {
	GetNodePerformanceData() ([]pb.NodePerformance, error)
	CreateNodePerformanceData(data pb.NodePerformance) error
}
