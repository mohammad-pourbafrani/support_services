package controllers

import (
	"context"
	pb "support_services_user_gw_http/proto/authentication/api"
)

type AuthenticationController struct {
	pb.UnimplementedAuthenticationServiceServer
}

func NewAuthenticationHandler() *AuthenticationController {
	return &AuthenticationController{}
}

func (c *AuthenticationController) SignIn(ctx context.Context, request *pb.SignInRequest) (*pb.SignInResponse, error) {
	return nil, nil
}

func (c *AuthenticationController) SignUp(ctx context.Context, request *pb.SignUpRequest) (*pb.Empty, error) {
	return nil, nil
}

func (c *AuthenticationController) Verify(ctx context.Context, request *pb.VerifyCode) (*pb.Token, error) {
	return nil, nil
}

func (c *AuthenticationController) ResetPassword(ctx context.Context, request *pb.ResetPasswordRequest) (*pb.Empty, error) {
	return nil, nil
}

func (c *AuthenticationController) ChangePassword(ctx context.Context, request *pb.ChangePasswordRequest) (*pb.Empty, error) {
	return nil, nil
}
