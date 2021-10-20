package informer

import (
	"context"
	"time"

	"github.com/adarocket/controller/auth"
	"github.com/adarocket/controller/config"
	pb "github.com/adarocket/proto/proto-gen/common"
)

const statusOK = "OK"
const statusError = "ERROR"
const statusWarning = "WARNING"

// CommonInformServer -
type CommonInformServer struct {
	loadedConfig config.Config
	pb.UnimplementedControllerServer

	NodeLastUpdate map[string]time.Time

	// NodeStatistics map[string]*pb.SaveStatisticRequest
	// jwtManager   *auth.JWTManager
}

// NewCommonInformServer -
func NewCommonInformServer(jwtManager *auth.JWTManager, loadedConfig config.Config) *CommonInformServer {
	return &CommonInformServer{
		loadedConfig:   loadedConfig,
		NodeLastUpdate: make(map[string]time.Time),

		// NodeStatistics: make(map[string]*pb.SaveStatisticRequest),
		// jwtManager:   jwtManager,
	}
}

// GetStatistic -
func (server *CommonInformServer) GetNodeList(ctx context.Context, request *pb.GetNodeListRequest) (response *pb.GetNodeListResponse, err error) {
	response = new(pb.GetNodeListResponse)
	tNow := time.Now()

	for _, n := range server.loadedConfig.Nodes {
		nodeAuthData := new(pb.NodeAuthData)
		nodeAuthData.Ticker = n.Ticker
		nodeAuthData.Uuid = n.UUID
		nodeAuthData.Blockchain = n.Blockchain

		nodeAuthData.Status = statusError

		lastUpdate, ok := server.NodeLastUpdate[n.UUID]
		if ok {
			elapsed := tNow.Sub(lastUpdate)
			if elapsed.Seconds() < 10 {
				nodeAuthData.Status = statusOK
			} else if elapsed.Seconds() < 60 {
				nodeAuthData.Status = statusWarning
			}
		}
		response.NodeAuthData = append(response.NodeAuthData, nodeAuthData)
	}

	// for _, n := range server.NodeStatistics {
	// 	response.NodeAuthData = append(response.NodeAuthData, n.NodeAuthData)
	// }

	return response, nil
}
