package server

import (
	"appconfigs"
	"appstates"
	"log"
	"net"
	"support_services_user_gw_http/internal/controllers"
	"support_services_user_gw_http/internal/middleware"
	pb "support_services_user_gw_http/proto/authentication/api"

	"google.golang.org/grpc"
)

func RunServer() {

	var (
		listenAddress               = appconfigs.String("listen-address", "server listen address")
		authenticationServerAddress = appconfigs.String("authentication-server-address", "authentication server address")
	)

	if err := appconfigs.Parse(); err != nil {
		appstates.PanicMissingEnvParams(err.Error())
	} else {
		appstates.DoneServerLaunch("master server " + *listenAddress + " and authentication server " + *authenticationServerAddress + " is run")
	}

	var (
		authenticationController controllers.AuthenticationController = *controllers.NewAuthenticationHandler()
	)

	//create net listenr for listen to grpc connection
	listener, err := net.Listen("tcp", *listenAddress)
	if err != nil {
		appstates.PanicServerSocketFailure(err.Error())
	}

	//create grpc and serve on listener
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(middleware.UnaryInterceptor))
	pb.RegisterAuthenticationServiceServer(grpcServer, &authenticationController)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
