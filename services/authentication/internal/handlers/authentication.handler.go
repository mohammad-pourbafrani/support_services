package handlers

import (
	"context"
	// "fmt"
	"support_services_authentication/internal/models"
	"support_services_authentication/internal/services"

	// "support_services_authentication/internal/models"
	pb "support_services_authentication/proto/api"
	// "google.golang.org/grpc/codes"
	// "google.golang.org/grpc/status"
)

type AuthenticationHandler struct {
	pb.UnimplementedAuthenticationServiceServer
	authenticationService services.AuthenticationService
	tokenService          services.TokenService
}

func NewAuthenticationHandler(authenticationService services.AuthenticationService, tokenService services.TokenService) *AuthenticationHandler {
	return &AuthenticationHandler{
		authenticationService: authenticationService,
		tokenService:          tokenService,
	}
}
func (c *AuthenticationHandler) SignIn(ctx context.Context, request *pb.SignInRequest) (*pb.SignInResponse, error) {
	// fmt.Println("call api sign in")
	// phone := request.PhoneNumber
	// pass := request.Password
	// fmt.Printf("your phone:%v , your pass:%v \n", phone, pass)
	// val := c.authenticationService.AddUser(&models.UserDto{Phone_number: phone, UserRole: "client", UserStatus: "enable"})
	// if val != nil {
	// 	return nil, nil
	// } else {
	// 	return nil, status.Error(codes.Unknown, "some thing error")
	// }
	return nil, nil
}

func (c *AuthenticationHandler) SignUp(ctx context.Context, request *pb.SignUpRequest) (*pb.Empty, error) {
	_, err := c.authenticationService.RegisterUser(&models.UserDto{PhoneNumber: request.GetPhoneNumber(), UserRole: "client", Verify: false})
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	return &pb.Empty{}, nil
}

func (c *AuthenticationHandler) Verify(ctx context.Context, request *pb.VerifyCode) (*pb.Token, error) {
	user, err := c.authenticationService.VerifyRegisterUser(&models.VerifyCodeDto{Code: request.Code, PhoneNumber: request.PhoneNumber})
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	token, tokenErr := c.tokenService.AddToken(user)
	if tokenErr != nil {
		return nil, tokenErr.ErrorToGRPCStatus()
	}
	return &pb.Token{
		AccessToken:    token.AccessToken,
		RefreshToken:   token.RefreshToken,
		AccessExpTime:  token.AccessExpireTime.Unix(),
		RefreshExpTime: token.AccessExpireTime.Unix(),
	}, nil
}

func (c *AuthenticationHandler) ResetPassword(ctx context.Context, request *pb.ResetPasswordRequest) (*pb.Empty, error) {
	return nil, nil
}

func (c *AuthenticationHandler) ChangePassword(ctx context.Context, request *pb.ChangePasswordRequest) (*pb.Empty, error) {
	return nil, nil
}

func (c *AuthenticationHandler) RefreshToken(ctx context.Context, request *pb.Token) (*pb.Token, error) {
	return nil, nil
}
