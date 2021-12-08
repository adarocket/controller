package structs

import (
	cardanoPb "github.com/adarocket/proto/proto-gen/cardano"
	commonPB "github.com/adarocket/proto/proto-gen/common"
)

type Database interface {
	CreateAllTables() error
	NodeInterface
	ServerDataInterface
	CaradanoDataInterface
	Ping() error
}

type Nodes struct {
	commonPB.NodeAuthData
	commonPB.NodeBasicData
}

type ServerData struct {
	Uuid string
	commonPB.Updates
	commonPB.CPUState
	commonPB.Online
	commonPB.MemoryState
	commonPB.Security
	commonPB.ServerBasicData
}

type CardanoData struct {
	Uuid string
	cardanoPb.Epoch
	cardanoPb.KESData
	cardanoPb.Blocks
	cardanoPb.StakeInfo
	cardanoPb.NodePerformance
	cardanoPb.NodeState
}

type NodeInterface interface {
	GetNodesData() ([]Nodes, error)
	CreateNodeData(data Nodes) error
}

type ServerDataInterface interface {
	GetNodeServerData() ([]ServerData, error)
	CreateNodeServerData(data ServerData) error
}

type CaradanoDataInterface interface {
	GetCardanoData() ([]CardanoData, error)
	CreateCardanoData(data CardanoData) error
}
