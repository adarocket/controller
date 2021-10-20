package informer

import (
	"context"

	"github.com/adarocket/controller/auth"
	"github.com/adarocket/controller/config"
	pb "github.com/adarocket/proto/proto-gen/common"
)

// CommonInformServer -
type CommonInformServer struct {
	loadedConfig config.Config
	pb.UnimplementedControllerServer

	// NodeStatistics map[string]*pb.SaveStatisticRequest
	// jwtManager   *auth.JWTManager

}

// NewCommonInformServer -
func NewCommonInformServer(jwtManager *auth.JWTManager, loadedConfig config.Config) *CommonInformServer {
	return &CommonInformServer{
		loadedConfig: loadedConfig,

		// NodeStatistics: make(map[string]*pb.SaveStatisticRequest),
		// jwtManager:   jwtManager,
	}
}

// GetStatistic -
func (server *CommonInformServer) GetNodeList(ctx context.Context, request *pb.GetNodeListRequest) (response *pb.GetNodeListResponse, err error) {
	response = new(pb.GetNodeListResponse)
	for _, n := range server.loadedConfig.Nodes {
		nodeAuthData := new(pb.NodeAuthData)
		nodeAuthData.Ticker = n.Ticker
		nodeAuthData.Uuid = n.UUID
		nodeAuthData.Blockchain = n.Blockchain
		nodeAuthData.Status = n.Blockchain

		response.NodeAuthData = append(response.NodeAuthData, nodeAuthData)
	}

	// for _, n := range server.NodeStatistics {
	// 	response.NodeAuthData = append(response.NodeAuthData, n.NodeAuthData)
	// }

	return response, nil
}
