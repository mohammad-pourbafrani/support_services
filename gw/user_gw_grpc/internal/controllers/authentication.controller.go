package controllers

import (
	"context"
	"support_services_user_gw_http/internal/services"
	pb "support_services_user_gw_http/proto/authentication/api"
)

type AuthenticationController struct {
	pb.UnimplementedAuthenticationServiceServer
	authenticationService services.AuthenticationService
}

func NewAuthenticationHandler(authenticationService services.AuthenticationService) *AuthenticationController {
	return &AuthenticationController{authenticationService: authenticationService}
}

func (c *AuthenticationController) SignIn(ctx context.Context, request *pb.SignInRequest) (*pb.SignInResponse, error) {
	return nil, nil
}

func (c *AuthenticationController) SignUp(ctx context.Context, request *pb.SignUpRequest) (*pb.Empty, error) {
	return c.authenticationService.SignUp(ctx, request)
}

func (c *AuthenticationController) Verify(ctx context.Context, request *pb.VerifyCode) (*pb.Token, error) {
	return c.authenticationService.VerifySignUp(ctx, request)
}

func (c *AuthenticationController) ResetPassword(ctx context.Context, request *pb.ResetPasswordRequest) (*pb.Empty, error) {
	return nil, nil
}

func (c *AuthenticationController) ChangePassword(ctx context.Context, request *pb.ChangePasswordRequest) (*pb.Empty, error) {
	return nil, nil
}

func (c *AuthenticationController) RefreshToken(ctx context.Context, request *pb.Token) (*pb.Token, error) {
	return nil, nil
}
