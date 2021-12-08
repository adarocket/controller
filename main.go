package main

import (
	auth2 "github.com/adarocket/controller/repository/auth"
	"github.com/adarocket/controller/repository/config"
	informer2 "github.com/adarocket/controller/repository/informer"
	user2 "github.com/adarocket/controller/repository/user"
	"log"
	"net"
	"time"

	authPB "github.com/adarocket/proto/proto-gen/auth"
	cardanoPB "github.com/adarocket/proto/proto-gen/cardano"
	chiaPB "github.com/adarocket/proto/proto-gen/chia"
	commonPB "github.com/adarocket/proto/proto-gen/common"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	secretKey     = "secret"
	tokenDuration = 15 * time.Minute
)

// var loadedConfig config.Config

func main() {
	loadedConfig, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	userStore := user2.NewInMemoryUserStore()
	if err := seedUsers(userStore); err != nil {
		log.Fatal("cannot seed users: ", err)
	}

	listener, err := net.Listen("tcp", ":"+loadedConfig.ServerPort)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	// ----------------------------------------------------------------------

	jwtManager := auth2.NewJWTManager(secretKey, tokenDuration)
	authServer := auth2.NewAuthServer(userStore, jwtManager)

	commonServer := informer2.NewCommonInformServer(jwtManager, loadedConfig)
	cardanoServer := informer2.NewCardanoInformServer(jwtManager, loadedConfig)
	chiaServer := informer2.NewChiaInformServer(jwtManager, loadedConfig)

	interceptor := auth2.NewAuthInterceptor(jwtManager, accessiblePermissions())

	/*db, err := postgresql.InitDatabase(loadedConfig)
	if err != nil {
		log.Println(err)
		return
	}
	err = db.CreateAllTables()
	if err != nil {
		log.Println(err)
	}
	go save.AutoSave(cardanoServer, db)*/

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
	)

	authPB.RegisterAuthServiceServer(grpcServer, authServer)
	commonPB.RegisterControllerServer(grpcServer, commonServer)
	cardanoPB.RegisterCardanoServer(grpcServer, cardanoServer)
	chiaPB.RegisterChiaServer(grpcServer, chiaServer)

	// ----------------------------------------------------------------------

	grpcServer.Serve(listener)
}

// ----------------------------------------------------------------

func createUser(userStore user2.UserStore, username, password string, permissions []string) error {
	user, err := user2.NewUser(username, password, permissions)
	if err != nil {
		return err
	}
	return userStore.Save(user)
}

func seedUsers(userStore user2.UserStore) error {
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
