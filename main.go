package main

import (
	"flag"
	"github.com/adarocket/controller/auth"
	"github.com/adarocket/controller/config"
	"github.com/adarocket/controller/informer"
	"github.com/adarocket/controller/user"
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
var add = flag.Bool("add", false, "add new node’s uuid to conf")
var remove = flag.Bool("remove", false, "remove exist node’s uuid from conf")
var status = flag.Bool("status", false, "show data for all nodes that are in the config")

func main() {
	flag.Parse()
	loadedConfig, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	/*db, err := postgresql.InitDatabase(loadedConfig)
	if err != nil {
		log.Println(err)
		return
	}*/

	if *add {
		if err = loadedConfig.CreateNodes(); err != nil {
			log.Println(err)
		}
	}
	if *remove {
		if err = loadedConfig.RemoveNodes(); err != nil {
			log.Println(err)
		}
	}
	if *status {
		/*nodes, err := db.GetNodesData()
		if err != nil {
			log.Println(err)
		} else {
			for _, node := range nodes {
				fmt.Println(node)
			}
		}*/
	}

	userStore := user.NewInMemoryUserStore()
	if err := seedUsers(userStore); err != nil {
		log.Fatal("cannot seed users: ", err)
	}

	listener, err := net.Listen("tcp", ":"+loadedConfig.ServerPort)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	// ----------------------------------------------------------------------

	jwtManager := auth.NewJWTManager(secretKey, tokenDuration)
	authServer := auth.NewAuthServer(userStore, jwtManager)

	commonServer := informer.NewCommonInformServer(jwtManager, loadedConfig)
	cardanoServer := informer.NewCardanoInformServer(jwtManager, loadedConfig)
	chiaServer := informer.NewChiaInformServer(jwtManager, loadedConfig)

	interceptor := auth.NewAuthInterceptor(jwtManager, accessiblePermissions())

	/*err = db.CreateAllTables()
	if err != nil {
		log.Println(err)
	}
	go save.AutoSave(cardanoServer, db, 5)*/

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

func createUser(userStore user.UserStore, username, password string, permissions []string) error {
	user, err := user.NewUser(username, password, permissions)
	if err != nil {
		log.Println(err)
		return err
	}
	return userStore.Save(user)
}

func seedUsers(userStore user.UserStore) error {
	if err := createUser(userStore, "admin1", "secret", []string{"basic", "server_technical", "node_technical", "node_financial"}); err != nil {
		log.Println(err)
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
