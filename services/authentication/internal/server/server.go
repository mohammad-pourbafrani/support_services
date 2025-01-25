package server

import (
	"appconfigs"
	"appstates"
	"fmt"
	"log"
	"net"
	"support_services_authentication/internal/database"
	"support_services_authentication/internal/handlers"
	"support_services_authentication/internal/middleware"
	"support_services_authentication/internal/redis"
	"support_services_authentication/internal/repository"
	"support_services_authentication/internal/services"
	pb "support_services_authentication/proto/api"

	"google.golang.org/grpc"
)

func RunServer() {

	var (
		listenAddress = appconfigs.String("listen-address", "server listen address")
		redisHost     = appconfigs.String("redis-host", "server redise host")
		redisPort     = appconfigs.String("redis-port", "server redise port")
		postgresHost  = appconfigs.String("postgres-host", "server postgres host")
		postgresPort  = appconfigs.Int("postgres-port", "server postgres port")
		dbName        = appconfigs.String("db-name", "data base name")
		dbUserName    = appconfigs.String("db-user-name", "user name database")
		dbPassword    = appconfigs.String("db-password", "password databse")
	)

	if err := appconfigs.Parse(); err != nil {
		appstates.PanicMissingEnvParams(err.Error())
	} else {
		appstates.DoneServerLaunch("authentication server " + *listenAddress + " is run")
	}

	//connect to PostgreSql database
	// db, dbErr := database.ConnectToPostgres("localhost", 5432, "support_services_db", "m.pourbafrani", "m.pourbafrani")
	db, dbErr := database.ConnectToPostgres(*postgresHost, *postgresPort, *dbName, *dbUserName, *dbPassword)
	if dbErr != nil {
		log.Fatalf("failed to connect  posgresql: %v", dbErr)
	}

	redisAddress := fmt.Sprintf("%v:%v", *redisHost, *redisPort)
	rdDb, rdDberr := redis.ConnectToRedis(redisAddress)
	if rdDberr != nil {
		appstates.PanicDBConnectionFailed(rdDberr.Error()) // Log an error if the database connection fails.
	}

	var (
		repository            repository.AuthenticationRepository = repository.NewAuthenticationRepository(db, rdDb)
		authenticationService services.AuthenticationService      = services.NewAuthenticationService(repository)
		authenticationHandler handlers.AuthenticationHandler      = *handlers.NewAuthenticationHandler(authenticationService)
	)

	//create net listenr for listen to grpc connection
	listener, err := net.Listen("tcp", *listenAddress)
	if err != nil {
		log.Fatalf("failed to listen on port 50051: %v", err)
	}
	fmt.Println("create listener")

	//create grpc and serve on listener
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(middleware.UnaryInterceptor))
	pb.RegisterAuthenticationServiceServer(grpcServer, &authenticationHandler)
	fmt.Println("create grpc")
	fmt.Printf("pre run server %v", listener.Addr().String())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	fmt.Println("after run server")

}
