package main

import (
	"log"
	"net"
	"time"

	"adarocket/controller/auth"
	"adarocket/controller/config"
	"adarocket/controller/informer"
	"adarocket/controller/user"

	pb "github.com/adarocket/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	secretKey     = "secret"
	tokenDuration = 15 * time.Minute
)

var loadedConfig config.Config

func main() {
	loadedConfig, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", ":"+loadedConfig.ServerPort)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	userStore := user.NewInMemoryUserStore()
	if err := seedUsers(userStore); err != nil {
		log.Fatal("cannot seed users: ", err)
	}

	jwtManager := auth.NewJWTManager(secretKey, tokenDuration)
	authServer := auth.NewAuthServer(userStore, jwtManager)
	informServer := informer.NewInformServer(jwtManager, loadedConfig)
	interceptor := auth.NewAuthInterceptor(jwtManager, accessiblePermissions())

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
	)

	pb.RegisterAuthServiceServer(grpcServer, authServer)
	pb.RegisterInformerServer(grpcServer, informServer)

	grpcServer.Serve(listener)
}

// ----------------------------------------------------------------

func createUser(userStore user.UserStore, username, password string, permissions []string) error {
	user, err := user.NewUser(username, password, permissions)
	if err != nil {
		return err
	}
	return userStore.Save(user)
}

func seedUsers(userStore user.UserStore) error {
	if err := createUser(userStore, "admin1", "secret", []string{"basic", "server_technical", "node_technical", "node_financial"}); err != nil {
		return err
	}

	return createUser(userStore, "user1", "secret", []string{"basic"})
}

func accessiblePermissions() map[string][]string {
	const informerServicePath = "/proto.Informer/"

	return map[string][]string{
		informerServicePath + "GetStatistic": {"basic", "server_technical", "node_technical", "node_financial"},
		informerServicePath + "GetNodeList":  {"basic", "server_technical", "node_technical", "node_financial"},
	}
}
