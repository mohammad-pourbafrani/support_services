package handlers

import (
	"context"
	// "fmt"
	"support_services_authentication/internal/models"
	"support_services_authentication/internal/services"
	"support_services_authentication/internal/types"

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

	var loginDto *models.LogInDto = &models.LogInDto{
		PhoneNumber:  request.PhoneNumber,
		Password:     "",
		SignInMethod: models.SignInMethod(request.SignInMethod)}

	if request.SignInMethod == pb.SignInMethod_PASSWORD {
		if request.Password == nil {
			return nil, types.NewBadRequestError("password is empty , please send password").ErrorToGRPCStatus()
		}
		loginDto.Password = *request.Password
	}

	user, err := c.authenticationService.LogIn(loginDto)

	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}

	if request.SignInMethod == pb.SignInMethod_PASSWORD {
		token, tokenErr := c.tokenService.AddToken(user)
		if tokenErr != nil {
			return nil, tokenErr.ErrorToGRPCStatus()
		}
		return &pb.SignInResponse{
			Token: &pb.Token{
				AccessToken:    token.AccessToken,
				RefreshToken:   token.RefreshToken,
				AccessExpTime:  token.AccessExpireTime.Unix(),
				RefreshExpTime: token.AccessExpireTime.Unix(),
			},
		}, nil
	}

	return nil, nil
}

func (c *AuthenticationHandler) SignUp(ctx context.Context, request *pb.SignUpRequest) (*pb.Empty, error) {
	err := c.authenticationService.RegisterUser(&models.UserDto{PhoneNumber: request.GetPhoneNumber(), Password: request.GetPassword(), UserRole: "client", Verify: false})
	if err != nil {
		return nil, err.ErrorToGRPCStatus()
	}
	return &pb.Empty{}, nil
}

func (c *AuthenticationHandler) Verify(ctx context.Context, request *pb.VerifyRequest) (*pb.Token, error) {

	if request.VerifyMethod == pb.VerifyMethod_SIGNUP {

		user, err := c.authenticationService.VerifyRegisterUser(&models.VerifyCodeDto{Code: request.VerifyCode.Code, PhoneNumber: request.VerifyCode.PhoneNumber})
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
	} else if request.VerifyMethod == pb.VerifyMethod_SIGNIN {
		user, err := c.authenticationService.LogInWithVerifyCode(&models.VerifyCodeDto{Code: request.VerifyCode.Code, PhoneNumber: request.VerifyCode.PhoneNumber})
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
	} else {
		return nil, types.NewBadRequestError("verify method not correct").ErrorToGRPCStatus()
	}

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
