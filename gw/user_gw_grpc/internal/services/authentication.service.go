package services

import (
	"context"
	pb "support_services_user_gw_http/proto/authentication/api"
)

type (
	AuthenticationService interface {
		SignUp(ctx context.Context, request *pb.SignUpRequest) (*pb.Empty, error)
	}

	authenticationService struct {
		authenticationClient pb.AuthenticationServiceClient
	}
)

func NewAuthenticationService(authenticationClient pb.AuthenticationServiceClient) AuthenticationService {
	return &authenticationService{authenticationClient: authenticationClient}
}

// SignIn implements AuthenticationService.
func (s *authenticationService) SignUp(ctx context.Context, request *pb.SignUpRequest) (*pb.Empty, error) {
	return s.authenticationClient.SignUp(ctx, request)
}
