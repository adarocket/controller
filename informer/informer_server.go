package informer

import (
	"context"

	"adarocket/controller/auth"
	"adarocket/controller/config"
	"adarocket/controller/helpers"

	pb "github.com/adarocket/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// InformServer -
type InformServer struct {
	NodeStatistics map[string]*pb.SaveStatisticRequest
	loadedConfig   config.Config

	jwtManager *auth.JWTManager
}

// NewInformServer -
func NewInformServer(jwtManager *auth.JWTManager, loadedConfig config.Config) *InformServer {
	return &InformServer{
		NodeStatistics: make(map[string]*pb.SaveStatisticRequest),
		jwtManager:     jwtManager,
		loadedConfig:   loadedConfig,
	}
}

// SaveStatistic -
func (server *InformServer) SaveStatistic(ctx context.Context, request *pb.SaveStatisticRequest) (response *pb.SaveStatisticResponse, err error) {
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
		nodeStatistic.NodeAuthData = new(pb.NodeAuthData)
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

	server.NodeStatistics[request.NodeAuthData.Uuid] = nodeStatistic

	// request

	response = &pb.SaveStatisticResponse{
		Status: "Ok",
	}

	return response, nil
}

// GetStatistic -
func (server *InformServer) GetNodeList(ctx context.Context, request *pb.GetNodeListRequest) (response *pb.GetNodeListResponse, err error) {
	response = new(pb.GetNodeListResponse)
	for _, n := range server.loadedConfig.Nodes {
		nodeAuthData := new(pb.NodeAuthData)
		nodeAuthData.Ticker = n.Ticker
		nodeAuthData.Uuid = n.UUID
		response.NodeAuthData = append(response.NodeAuthData, nodeAuthData)
	}

	// for _, n := range server.NodeStatistics {
	// 	response.NodeAuthData = append(response.NodeAuthData, n.NodeAuthData)
	// }

	return response, nil
}

// GetStatistic -
func (server *InformServer) GetStatistic(ctx context.Context, request *pb.GetStatisticRequest) (response *pb.SaveStatisticRequest, err error) {
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
	response.NodeAuthData = new(pb.NodeAuthData)

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
