package v1

import (
	"context"
	"github.com/sunkuet02/go-grpc-rest/pkg/api/v1"
)

type authenticationService struct {
}

func NewAuthenticationServiceServer() v1.AuthenticationServiceServer {
	return &authenticationService{}
}

func (a *authenticationService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	return &v1.LoginResponse{
		AccessToken: "token",
		Success:true,
	}, nil
}
