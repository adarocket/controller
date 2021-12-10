package auth

import (
	"context"
	"github.com/adarocket/controller/repository/user"
	"log"

	pb "github.com/adarocket/proto/proto-gen/auth"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AuthServer -
type AuthServer struct {
	userStore  user.UserStore
	jwtManager *JWTManager

	pb.UnimplementedAuthServiceServer
}

// NewAuthServer -
func NewAuthServer(userStore user.UserStore, jwtManager *JWTManager) *AuthServer {
	return &AuthServer{
		userStore:  userStore,
		jwtManager: jwtManager,
	}
}

// Login -
func (server *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Println("Starting login...")
	user, err := server.userStore.Find(req.GetUsername())
	if err != nil {
		log.Println("cannot find user")
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}

	if user == nil || !user.IsCorrectPassword(req.GetPassword()) {
		log.Println("incorrect username/password")
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}

	token, err := server.jwtManager.Generate(user)
	if err != nil {
		log.Println("cannot generate access token")
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	res := &pb.LoginResponse{AccessToken: token}
	return res, nil
}
