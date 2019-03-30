package v1

import (
	"context"
	"github.com/sunkuet02/go-grpc-rest/pkg/api/v1"
)

type greetingsService struct {
}

func NewGreetingsServiceServer() v1.GreetingsServiceServer {
	return &greetingsService{}
}

func (g *greetingsService) SayHello(ctx context.Context, req *v1.SayHelloRequest) (*v1.SayHelloResponse, error) {
	return &v1.SayHelloResponse{
		Message: "Hello " + req.Name,
	}, nil
}
