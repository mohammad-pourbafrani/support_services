package services

import (
	"context"
	pb "support_services_user_gw_http/proto/authentication/api"
)

type (
	AuthenticationService interface {
		SignUp(ctx context.Context, request *pb.SignUpRequest) (*pb.Empty, error)
		VerifySignUp(ctx context.Context, request *pb.VerifyRequest) (*pb.Token, error)
		SignIn(ctx context.Context, request *pb.SignInRequest) (*pb.SignInResponse, error)
		ResetPass(ctx context.Context, request *pb.ResetPasswordRequest) (*pb.Empty, error)
		ChangePass(ctx context.Context, request *pb.ChangePasswordRequest) (*pb.Empty, error)
		RNewToken(ctx context.Context, request *pb.Token) (*pb.Token, error)
	}

	authenticationService struct {
		authenticationClient pb.AuthenticationServiceClient
	}
)

func NewAuthenticationService(authenticationClient pb.AuthenticationServiceClient) AuthenticationService {
	return &authenticationService{authenticationClient: authenticationClient}
}

// SignUp implements AuthenticationService.
func (s *authenticationService) SignUp(ctx context.Context, request *pb.SignUpRequest) (*pb.Empty, error) {
	return s.authenticationClient.SignUp(ctx, request)
}

// VerifyOtp implements AuthenticationService.
func (s *authenticationService) VerifySignUp(ctx context.Context, request *pb.VerifyRequest) (*pb.Token, error) {
	return s.authenticationClient.Verify(ctx, request)
}

func (s *authenticationService) SignIn(ctx context.Context, request *pb.SignInRequest) (*pb.SignInResponse, error) {
	return s.authenticationClient.SignIn(ctx, request)
}

func (s *authenticationService) ResetPass(ctx context.Context, request *pb.ResetPasswordRequest) (*pb.Empty, error) {
	return s.authenticationClient.ResetPassword(ctx, request)
}

func (s *authenticationService) ChangePass(ctx context.Context, request *pb.ChangePasswordRequest) (*pb.Empty, error) {
	return s.authenticationClient.ChangePassword(ctx, request)
}

func (s *authenticationService) RNewToken(ctx context.Context, request *pb.Token) (*pb.Token, error) {
	return s.authenticationClient.RefreshToken(ctx, request)
}
