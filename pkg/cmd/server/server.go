package cmd

import (
	"context"
	"flag"
	"fmt"
	"github.com/sunkuet02/go-grpc-rest/pkg/logger"
	"github.com/sunkuet02/go-grpc-rest/pkg/protocol/grpc"
	"github.com/sunkuet02/go-grpc-rest/pkg/protocol/rest"
)

type Config struct {

	GRPCPort string

	HTTPPort string

	LogLevel int

	LogTimeFormat string
}

func RunServer() error {
	ctx := context.Background()

	var cfg Config

	flag.StringVar(&cfg.GRPCPort, "grpc-port", "", "gRPC port to bind")
	flag.StringVar(&cfg.HTTPPort, "http-port", "", "HTTP port to bind")
	flag.IntVar(&cfg.LogLevel, "log-level", 0, "Global log level")
	flag.StringVar(&cfg.LogTimeFormat, "log-time-format", "", "Print time format for logger")
	flag.Parse()

	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("Invalid TCP port for gRPC based server: '%s'", cfg.GRPCPort)
	}

	if len(cfg.HTTPPort) == 0 {
		return fmt.Errorf("Invalid TCP port for HTTP based server: '%s'", cfg.HTTPPort)
	}

	if err := logger.Init(cfg.LogLevel, cfg.LogTimeFormat); err != nil {
		return fmt.Errorf("Failed to initialize logger: %v\n", err)
	}

	go func() {
		_ = rest.RunServer(ctx, cfg.GRPCPort, cfg.HTTPPort)
	}()

	return grpc.RunServer(ctx, cfg.GRPCPort);
}
