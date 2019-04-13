package cmd

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
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

func init() {
	viper.SetConfigFile(`config/config.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}

}

func RunServer() error {
	ctx := context.Background()

	var cfg Config

	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connectionString := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", connectionString)
	defer db.Close()
	if err != nil {
		return fmt.Errorf("Invalid database connection : %s\n", err)
	}

	cfg.GRPCPort = viper.GetString(`server.grpc-port`)
	cfg.HTTPPort = viper.GetString(`server.http-port`)

	if len(cfg.GRPCPort) == 0 {
		return fmt.Errorf("Invalid TCP port for gRPC based server: '%s'", cfg.GRPCPort)
	}

	if len(cfg.HTTPPort) == 0 {
		return fmt.Errorf("Invalid TCP port for HTTP based server: '%s'", cfg.HTTPPort)
	}

	cfg.LogLevel = viper.GetInt(`log.level`)
	cfg.LogTimeFormat = viper.GetString(`log.time-format`)
	if err := logger.Init(cfg.LogLevel, cfg.LogTimeFormat); err != nil {
		return fmt.Errorf("Failed to initialize logger: %v\n", err)
	}

	go func() {
		_ = rest.RunServer(ctx, cfg.GRPCPort, cfg.HTTPPort)
	}()

	return grpc.RunServer(ctx, cfg.GRPCPort);
}
