package grpc

import (
	"context"
	api "github.com/sunkuet02/go-grpc-rest/pkg/api/v1"
	"github.com/sunkuet02/go-grpc-rest/pkg/logger"
	service "github.com/sunkuet02/go-grpc-rest/pkg/service/v1"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
)

func RunServer(ctx context.Context, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	api.RegisterAuthenticationServiceServer(server, service.NewAuthenticationServiceServer())
	api.RegisterGreetingsServiceServer(server, service.NewGreetingsServiceServer())

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			logger.Log.Warn("Shutting down gRPC server");
			server.GracefulStop();
			<-ctx.Done()
		}
	}()

	logger.Log.Info("Starting gRPC server on port : " + port);
	return server.Serve(listen)
}
