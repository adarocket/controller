package informer

import (
	"context"

	"github.com/adarocket/controller/auth"
	"github.com/adarocket/controller/config"
	"github.com/adarocket/controller/helpers"

	chiaPB "github.com/adarocket/proto/proto-gen/chia"
	commonPB "github.com/adarocket/proto/proto-gen/common"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// ChiaInformServer -
type ChiaInformServer struct {
	NodeStatistics map[string]*chiaPB.SaveStatisticRequest
	loadedConfig   config.Config

	jwtManager *auth.JWTManager

	chiaPB.UnimplementedChiaServer
}

// NewChiaInformServer -
func NewChiaInformServer(jwtManager *auth.JWTManager, loadedConfig config.Config) *ChiaInformServer {
	return &ChiaInformServer{
		NodeStatistics: make(map[string]*chiaPB.SaveStatisticRequest),
		jwtManager:     jwtManager,
		loadedConfig:   loadedConfig,
	}
}

// SaveStatistic -
func (server *ChiaInformServer) SaveStatistic(ctx context.Context, request *chiaPB.SaveStatisticRequest) (response *chiaPB.SaveStatisticResponse, err error) {
	if !helpers.Contains(server.loadedConfig.Nodes, request.NodeAuthData.Ticker, request.NodeAuthData.Uuid) {
		response = &chiaPB.SaveStatisticResponse{
			Status: "Error",
		}
		return response, nil
	}

	nodeStatistic := server.NodeStatistics[request.NodeAuthData.Uuid]
	if nodeStatistic == nil {
		nodeStatistic = new(chiaPB.SaveStatisticRequest)
		nodeStatistic.Statistic = new(chiaPB.Statistic)
		nodeStatistic.NodeAuthData = new(commonPB.NodeAuthData)
	}

	// -------------- Common Start --------------
	if request.NodeAuthData != nil {
		nodeStatistic.NodeAuthData = request.NodeAuthData
	}

	if request.Statistic.NodeBasicData != nil {
		nodeStatistic.Statistic.NodeBasicData = request.Statistic.NodeBasicData
	}

	if request.Statistic.ServerBasicData != nil {
		nodeStatistic.Statistic.ServerBasicData = request.Statistic.ServerBasicData
	}

	if request.Statistic.Updates != nil {
		nodeStatistic.Statistic.Updates = request.Statistic.Updates
	}

	if request.Statistic.Security != nil {
		nodeStatistic.Statistic.Security = request.Statistic.Security
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
	// -------------- Common End --------------

	if request.Statistic.ChiaNodeFarming != nil {
		nodeStatistic.Statistic.ChiaNodeFarming = request.Statistic.ChiaNodeFarming
	}

	server.NodeStatistics[request.NodeAuthData.Uuid] = nodeStatistic

	// request

	response = &chiaPB.SaveStatisticResponse{
		Status: "Ok",
	}

	return response, nil
}

// GetStatistic -
func (server *ChiaInformServer) GetStatistic(ctx context.Context, request *chiaPB.GetStatisticRequest) (response *chiaPB.SaveStatisticRequest, err error) {
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

	response = new(chiaPB.SaveStatisticRequest)
	response.Statistic = new(chiaPB.Statistic)
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
			response.Statistic.ChiaNodeFarming = node.Statistic.ChiaNodeFarming

		case "node_financial":
		}
	}

	return response, nil
}
