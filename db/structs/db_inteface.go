package structs

import pb "github.com/adarocket/proto/proto"

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
	ChiaNodeFarming
}

type NodeAuthData interface {
	GetNodeAuthData() ([]pb.NodeAuthData, error)
	CreateNodeAuthData(data *pb.NodeAuthData) error
	UpdateNodeAuthData(data *pb.NodeAuthData) error
	DeleteNodeAuthData(data *pb.NodeAuthData) error
}

type NodeBasicData interface {
	GetNodeBasicData() ([]pb.NodeBasicData, error)
	CreateNodeBasicData(data *pb.NodeBasicData, uuid string) error
}

type ServerBasicData interface {
	GetServerBasicData() ([]pb.ServerBasicData, error)
	CreateServerBasicData(data *pb.ServerBasicData, uuid string) error
}

type Epoch interface {
	GetEpochData() ([]pb.Epoch, error)
	CreateEpochData(data *pb.Epoch, uuid string) error
}

type KesData interface {
	GetKesData() ([]pb.KESData, error)
	CreateKesData(data *pb.KESData, uuid string) error
}

type Blocks interface {
	GetBlocksData() ([]pb.Blocks, error)
	CreateBlocksData(data *pb.Blocks, uuid string) error
}

type Updates interface {
	GetUpdatesData() ([]pb.Updates, error)
	CreateUpdatesData(data *pb.Updates, uuid string) error
}

type Security interface {
	GetSecurityData() ([]pb.Security, error)
	CreateSecurityData(data *pb.Security, uuid string) error
}

type StakeInfo interface {
	GetStakeInfoData() ([]pb.StakeInfo, error)
	CreateStakeInfoData(data *pb.StakeInfo, uuid string) error
}

type Online interface {
	GetOnlineData() ([]pb.Online, error)
	CreateOnlineData(data *pb.Online, uuid string) error
}

type MemoryState interface {
	GetMemoryStateData() ([]pb.MemoryState, error)
	CreateMemoryStateData(data *pb.MemoryState, uuid string) error
}

type CpuState interface {
	GetCpuStateData() ([]pb.CPUState, error)
	CreateCpuStateData(data *pb.CPUState, uuid string) error
}

type NodeState interface {
	GetNodeStateData() ([]pb.NodeState, error)
	CreateNodeStateData(data *pb.NodeState, uuid string) error
}

type NodePerformance interface {
	GetNodePerformanceData() ([]pb.NodePerformance, error)
	CreateNodePerformanceData(data *pb.NodePerformance, uuid string) error
}

type ChiaNodeFarming interface {
	GetChiaNodeFarmingData() ([]pb.ChiaNodeFarming, error)
	CreateChiaNodeFarmingData(data *pb.ChiaNodeFarming, uuid string) error
}
