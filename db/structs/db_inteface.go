package structs

import (
	"time"

	cardanoPb "github.com/adarocket/proto/proto-gen/cardano"
	commonPB "github.com/adarocket/proto/proto-gen/common"
)

type Database interface {
	CreateAllTables() error
	NodeInterface
	ServerDataInterface
	CardanoDataInterface
	Ping() error
}

type Node struct {
	NodeAuthData  commonPB.NodeAuthData
	NodeBasicData commonPB.NodeBasicData
	LastUpdate    time.Time
}

type ServerData struct {
	Uuid string
	commonPB.Updates
	commonPB.CPUState
	commonPB.Online
	commonPB.MemoryState
	commonPB.Security
	commonPB.ServerBasicData
	LastUpdate time.Time
}

type CardanoData struct {
	Uuid string
	cardanoPb.Epoch
	cardanoPb.KESData
	cardanoPb.Blocks
	cardanoPb.StakeInfo
	cardanoPb.NodePerformance
	cardanoPb.NodeState
	LastUpdate time.Time
}

type NodeInterface interface {
	GetNodesData() ([]Node, error)
	GetNodeData(uuid string) (Node, error)
	CreateNodeData(data Node) error
}

type ServerDataInterface interface {
	GetNodeServerData(uuid string) ([]ServerData, error)
	CreateNodeServerData(data ServerData) error
}

type CardanoDataInterface interface {
	GetCardanoData(uuid string) ([]CardanoData, error)
	CreateCardanoData(data CardanoData) error
}
