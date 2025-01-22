package handlers

import (
	"context"
	"fmt"
	"support_services_authentication/internal/services"

	"support_services_authentication/internal/models"
	pb "support_services_authentication/proto/api"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthenticationHandler struct {
	pb.UnimplementedAuthenticationServiceServer
	authenticationService services.AuthenticationService
}

func NewAuthenticationHandler(authenticationService services.AuthenticationService) *AuthenticationHandler {
	return &AuthenticationHandler{
		authenticationService: authenticationService,
	}
}
func (c *AuthenticationHandler) SignIn(ctx context.Context, request *pb.SigninRequest) (*pb.SigninResponse, error) {
	fmt.Println("call api sign in")
	phone := request.PhoneNumber
	pass := request.PassWord
	fmt.Printf("your phone:%v , your pass:%v \n", phone, pass)
	val := c.authenticationService.AddUser(&models.UserDto{Phone_number: phone, UserRole: "client", UserStatus: "enable"})
	if val != nil {
		return &pb.SigninResponse{Data: &pb.UserInfo{UserId: fmt.Sprint(val.UserId), PhoneNumber: val.Phone_number, UserRole: val.UserRole, UserStatus: val.UserStatus}, Message: "success"}, nil
	} else {
		return nil, status.Error(codes.Unknown, "some thing error")
	}
}
