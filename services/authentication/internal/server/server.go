package server

import (
	"fmt"
	"log"
	"net"
	// "support_services_authentication/internal/database"
	"support_services_authentication/internal/handlers"
	pb "support_services_authentication/proto/api"

	"google.golang.org/grpc"
)

func RunServer() {

	var (
		authenticationHandler handlers.AuthenticationHandler = *handlers.NewAuthenticationHandler()
	)

	//connect to PostgreSql database 
	// db , dbErr := database.ConnectToPostgres("",5432 ,"support_services_db", "podtgres","m.pourbafrani")
	// if dbErr !=nil{
	// 	log.Fatalf("failed to connect  posgresql: %v", dbErr)
	// }

	//create net listenr for listen to grpc connection
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen on port 50051: %v", err)
	}
	fmt.Println("create listener")


	//create grpc and serve on listener
	grpcServer := grpc.NewServer()
	pb.RegisterAuthenticationServiceServer(grpcServer, &authenticationHandler)
	fmt.Println("create grpc")
	fmt.Printf("pre run server %v", listener.Addr().String())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	fmt.Println("after run server")

}
