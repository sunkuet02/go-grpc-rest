package rest

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sunkuet02/go-grpc-rest/pkg/api/v1"
	"github.com/sunkuet02/go-grpc-rest/pkg/protocol/rest/middleware"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func RunServer(ctx context.Context, grpcPort, httpPort string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := v1.RegisterAuthenticationServiceHandlerFromEndpoint(ctx, mux, "localhost:"+grpcPort, opts); err != nil {
		log.Fatalf("Failed to start HTTP gateway: %v", err)
	}

	if err := v1.RegisterGreetingsServiceHandlerFromEndpoint(ctx, mux, "localhost:"+grpcPort, opts); err != nil {
		log.Fatalf("Failed to start HTTP gateway: %v", err)
	}

	server := &http.Server{
		Addr:    ":" + httpPort,
		Handler: middleware.AddRequestID(middleware.AddLogger(mux)),
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {

		}
		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		_ = server.Shutdown(ctx)
	}()

	log.Println("Starting REST/HTTP server in port " + httpPort)
	return server.ListenAndServe()
}
