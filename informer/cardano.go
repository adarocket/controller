package informer

import (
	"context"

	"github.com/adarocket/controller/auth"
	"github.com/adarocket/controller/config"
	"github.com/adarocket/controller/helpers"

	pb "github.com/adarocket/proto/proto-gen/cardano"
	commonPB "github.com/adarocket/proto/proto-gen/common"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// CardanoInformServer -
type CardanoInformServer struct {
	NodeStatistics map[string]*pb.SaveStatisticRequest
	loadedConfig   config.Config

	jwtManager *auth.JWTManager

	pb.UnimplementedCardanoServer
}

// NewCardanoInformServer -
func NewCardanoInformServer(jwtManager *auth.JWTManager, loadedConfig config.Config) *CardanoInformServer {
	return &CardanoInformServer{
		NodeStatistics: make(map[string]*pb.SaveStatisticRequest),
		jwtManager:     jwtManager,
		loadedConfig:   loadedConfig,
	}
}

// SaveStatistic -
func (server *CardanoInformServer) SaveStatistic(ctx context.Context, request *pb.SaveStatisticRequest) (response *pb.SaveStatisticResponse, err error) {
	// request.NodeAuthData.Ticker
	// request.NodeAuthData.Uuid
	// server.loadedConfig.Nodes

	if !helpers.Contains(server.loadedConfig.Nodes, request.NodeAuthData.Ticker, request.NodeAuthData.Uuid) {
		response = &pb.SaveStatisticResponse{
			Status: "Error",
		}
		return response, nil
	}

	nodeStatistic := server.NodeStatistics[request.NodeAuthData.Uuid]
	if nodeStatistic == nil {
		nodeStatistic = new(pb.SaveStatisticRequest)
		nodeStatistic.Statistic = new(pb.Statistic)
		nodeStatistic.NodeAuthData = new(commonPB.NodeAuthData)
	}

	// Ticker - string
	// Uuid - string
	// Status - string
	if request.NodeAuthData != nil {
		nodeStatistic.NodeAuthData = request.NodeAuthData
	}

	// Ticker - string
	// Type - string
	// Location - string
	// NodeVersion - string
	if request.Statistic.NodeBasicData != nil {
		nodeStatistic.Statistic.NodeBasicData = request.Statistic.NodeBasicData
	}

	// Ipv4 - string
	// Ipv6 - string
	// LinuxName - string
	// LinuxVersion - string
	if request.Statistic.ServerBasicData != nil {
		nodeStatistic.Statistic.ServerBasicData = request.Statistic.ServerBasicData
	}

	// EpochNumber - int64
	if request.Statistic.Epoch != nil {
		nodeStatistic.Statistic.Epoch = request.Statistic.Epoch
	}

	// KesCurrent - int64
	// KesRemaining - int64
	// KesExpDate - string
	if request.Statistic.KesData != nil {
		nodeStatistic.Statistic.KesData = request.Statistic.KesData
	}

	// BlockLeader - int64
	// BlockAdopted - int64
	// BlockInvalid - int64
	if request.Statistic.Blocks != nil {
		nodeStatistic.Statistic.Blocks = request.Statistic.Blocks
	}

	// InformerActual - string
	// InformerAvailable - string
	// UpdaterActual - string
	// UpdaterAvailable - string
	// PackagesAvailable - int64
	if request.Statistic.Updates != nil {
		nodeStatistic.Statistic.Updates = request.Statistic.Updates
	}

	// SshAttackAttempts - int64
	// SecurityPackagesAvailable - int64
	if request.Statistic.Security != nil {
		nodeStatistic.Statistic.Security = request.Statistic.Security
	}

	// LiveStake - int64
	// ActiveStake - int64
	// Pledge - int64
	if request.Statistic.StakeInfo != nil {
		nodeStatistic.Statistic.StakeInfo = request.Statistic.StakeInfo
	}

	// SinceStart - int64
	// Pings - int64
	// NodeActive - bool
	// NodeActivePings - int64
	// ServerActive - bool
	if request.Statistic.Online != nil {
		nodeStatistic.Statistic.Online = request.Statistic.Online
	}

	// Total - uint64
	// Used - uint64
	// Buffers - uint64
	// Cached - uint64
	// Free - uint64
	// Available - uint64
	// Active - uint64
	// Inactive - uint64
	// SwapTotal - uint64
	// SwapUsed - uint64
	// SwapCached - uint64
	// SwapFree - uint64
	// MemAvailableEnabled - bool
	if request.Statistic.MemoryState != nil {
		nodeStatistic.Statistic.MemoryState = request.Statistic.MemoryState
	}

	// CpuQty - int64
	// AverageWorkload - float32
	if request.Statistic.CpuState != nil {
		nodeStatistic.Statistic.CpuState = request.Statistic.CpuState
	}

	// TipDiff - int64
	// Density - float32
	if request.Statistic.NodeState != nil {
		nodeStatistic.Statistic.NodeState = request.Statistic.NodeState
	}

	// ProcessedTx - int64
	// PeersIn - int64
	// PeersOut - int64
	if request.Statistic.NodePerformance != nil {
		nodeStatistic.Statistic.NodePerformance = request.Statistic.NodePerformance
	}

	server.NodeStatistics[request.NodeAuthData.Uuid] = nodeStatistic

	// request

	response = &pb.SaveStatisticResponse{
		Status: "Ok",
	}

	return response, nil
}

// GetStatistic -
func (server *CardanoInformServer) GetStatistic(ctx context.Context, request *pb.GetStatisticRequest) (response *pb.SaveStatisticRequest, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return response, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return response, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	claims, err := server.jwtManager.Verify(values[0])
	if err != nil {
		return response, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	response = new(pb.SaveStatisticRequest)
	response.Statistic = new(pb.Statistic)
	response.NodeAuthData = new(commonPB.NodeAuthData)

	node := server.NodeStatistics[request.Uuid]
	if node == nil {
		return response, nil
	}

	response.NodeAuthData = node.NodeAuthData

	for _, permission := range claims.Permissions {
		switch permission {
		case "basic":
			response.Statistic.NodeBasicData = node.Statistic.NodeBasicData
			response.Statistic.ServerBasicData = node.Statistic.ServerBasicData
			response.Statistic.Online = node.Statistic.Online

		case "server_technical":
			response.Statistic.MemoryState = node.Statistic.MemoryState
			response.Statistic.CpuState = node.Statistic.CpuState
			response.Statistic.Updates = node.Statistic.Updates
			response.Statistic.Security = node.Statistic.Security

		case "node_technical":
			response.Statistic.Epoch = node.Statistic.Epoch
			response.Statistic.NodeState = node.Statistic.NodeState
			response.Statistic.NodePerformance = node.Statistic.NodePerformance
			response.Statistic.KesData = node.Statistic.KesData

		case "node_financial":
			response.Statistic.Blocks = node.Statistic.Blocks
			response.Statistic.StakeInfo = node.Statistic.StakeInfo
		}
	}

	return response, nil
}
