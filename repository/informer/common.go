package informer

import (
	"context"
	"log"

	"github.com/adarocket/controller/repository/auth"
	"github.com/adarocket/controller/repository/config"
	"github.com/adarocket/controller/repository/repnodes"

	pb "github.com/adarocket/proto/proto-gen/common"
)

// CommonInformServer -
type CommonInformServer struct {
	loadedConfig config.Config
	pb.UnimplementedControllerServer
	repoNodes repnodes.RepoNodes
	// NodeStatistics map[string]*pb.SaveStatisticRequest
	// jwtManager   *auth.JWTManager

}

// NewCommonInformServer -
func NewCommonInformServer(jwtManager *auth.JWTManager, loadedConfig config.Config) *CommonInformServer {
	return &CommonInformServer{
		loadedConfig: loadedConfig,
		repoNodes:    repnodes.InitController(loadedConfig),
		// NodeStatistics: make(map[string]*pb.SaveStatisticRequest),
		// jwtManager:   jwtManager,
	}
}

// GetStatistic -
func (server *CommonInformServer) GetNodeList(ctx context.Context, request *pb.GetNodeListRequest) (response *pb.GetNodeListResponse, err error) {
	response = new(pb.GetNodeListResponse)
	for _, n := range server.loadedConfig.Nodes {
		node, err := server.repoNodes.GetNodeData(n.UUID)
		if err != nil {
			log.Println(err)
			continue
		}

		nodeAuthData := new(pb.NodeAuthData)
		nodeAuthData.Uuid = n.UUID
		nodeAuthData.Ticker = node.NodeAuthData.Ticker
		nodeAuthData.Blockchain = node.NodeAuthData.Blockchain

		nodeAuthData.Type = node.NodeAuthData.Type
		nodeAuthData.Name = node.NodeAuthData.Name

		response.NodeAuthData = append(response.NodeAuthData, nodeAuthData)
	}

	// for _, n := range server.NodeStatistics {
	// 	response.NodeAuthData = append(response.NodeAuthData, n.NodeAuthData)
	// }

	return response, nil
}
