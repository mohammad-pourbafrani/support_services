package handlers

import (
	"context"
	"fmt"
	pb "support_services_authentication/proto/api"
)

type AuthenticationHandler struct {
	pb.UnimplementedAuthenticationServiceServer
}

func NewAuthenticationHandler() *AuthenticationHandler {
	return &AuthenticationHandler{}
}
func (c *AuthenticationHandler) SignIn(ctx context.Context, request *pb.SigninRequest) (*pb.SigninResponse, error) {
	fmt.Println("call api sign in")
	phone := request.PhoneNumber
	pass := request.PassWord
	fmt.Printf("your phone:%v , your pass:%v \n", phone, pass)
	return &pb.SigninResponse{Data: "corrected", Message: "success"}, nil
}
