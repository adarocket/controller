package informer

import (
	"context"
	"time"

	"github.com/adarocket/controller/auth"
	"github.com/adarocket/controller/config"
	"github.com/adarocket/controller/helpers"

	cardanoPB "github.com/adarocket/proto/proto-gen/cardano"
	commonPB "github.com/adarocket/proto/proto-gen/common"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// CardanoInformServer -
type CardanoInformServer struct {
	NodeStatistics map[string]*cardanoPB.SaveStatisticRequest
	loadedConfig   config.Config

	jwtManager         *auth.JWTManager
	commonInformServer *CommonInformServer

	cardanoPB.UnimplementedCardanoServer
}

// NewCardanoInformServer -
func NewCardanoInformServer(jwtManager *auth.JWTManager, loadedConfig config.Config, commonInformServer *CommonInformServer) *CardanoInformServer {
	return &CardanoInformServer{
		NodeStatistics:     make(map[string]*cardanoPB.SaveStatisticRequest),
		jwtManager:         jwtManager,
		loadedConfig:       loadedConfig,
		commonInformServer: commonInformServer,
	}
}

// SaveStatistic -
func (server *CardanoInformServer) SaveStatistic(ctx context.Context, request *cardanoPB.SaveStatisticRequest) (response *cardanoPB.SaveStatisticResponse, err error) {
	// request.NodeAuthData.Ticker
	// request.NodeAuthData.Uuid
	// server.loadedConfig.Nodes

	if !helpers.Contains(server.loadedConfig.Nodes, request.NodeAuthData.Ticker, request.NodeAuthData.Uuid) {
		response = &cardanoPB.SaveStatisticResponse{
			Status: "Error",
		}
		return response, nil
	}

	nodeStatistic := server.NodeStatistics[request.NodeAuthData.Uuid]
	if nodeStatistic == nil {
		nodeStatistic = new(cardanoPB.SaveStatisticRequest)
		nodeStatistic.Statistic = new(cardanoPB.Statistic)
		nodeStatistic.NodeAuthData = new(commonPB.NodeAuthData)
	}

	if request.NodeAuthData != nil {
		nodeStatistic.NodeAuthData = request.NodeAuthData
	}

	if request.Statistic.NodeBasicData != nil {
		nodeStatistic.Statistic.NodeBasicData = request.Statistic.NodeBasicData
	}

	if request.Statistic.ServerBasicData != nil {
		nodeStatistic.Statistic.ServerBasicData = request.Statistic.ServerBasicData
	}

	if request.Statistic.Epoch != nil {
		nodeStatistic.Statistic.Epoch = request.Statistic.Epoch
	}

	if request.Statistic.KesData != nil {
		nodeStatistic.Statistic.KesData = request.Statistic.KesData
	}

	if request.Statistic.Blocks != nil {
		nodeStatistic.Statistic.Blocks = request.Statistic.Blocks
	}

	if request.Statistic.Updates != nil {
		nodeStatistic.Statistic.Updates = request.Statistic.Updates
	}

	if request.Statistic.Security != nil {
		nodeStatistic.Statistic.Security = request.Statistic.Security
	}

	if request.Statistic.StakeInfo != nil {
		nodeStatistic.Statistic.StakeInfo = request.Statistic.StakeInfo
	}

	if request.Statistic.Online != nil {
		nodeStatistic.Statistic.Online = request.Statistic.Online
	}

	if request.Statistic.MemoryState != nil {
		nodeStatistic.Statistic.MemoryState = request.Statistic.MemoryState
	}

	if request.Statistic.CpuState != nil {
		nodeStatistic.Statistic.CpuState = request.Statistic.CpuState
	}

	if request.Statistic.NodeState != nil {
		nodeStatistic.Statistic.NodeState = request.Statistic.NodeState
	}

	if request.Statistic.NodePerformance != nil {
		nodeStatistic.Statistic.NodePerformance = request.Statistic.NodePerformance
	}

	server.commonInformServer.NodeLastUpdate[request.NodeAuthData.Uuid] = time.Now()

	server.NodeStatistics[request.NodeAuthData.Uuid] = nodeStatistic

	// request

	response = &cardanoPB.SaveStatisticResponse{
		Status: "Ok",
	}

	return response, nil
}

// GetStatistic -
func (server *CardanoInformServer) GetStatistic(ctx context.Context, request *cardanoPB.GetStatisticRequest) (response *cardanoPB.SaveStatisticRequest, err error) {
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

	response = new(cardanoPB.SaveStatisticRequest)
	response.Statistic = new(cardanoPB.Statistic)
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
