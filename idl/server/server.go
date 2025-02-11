package server

import (
	"context"
	"fmt"
	"github.com/nuzur/nem/core"
	pb "github.com/nuzur/nem/idl/gen"
	"github.com/nuzur/nem/monitoring"
	"net"

	"go.uber.org/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	"github.com/nuzur/nem/auth"
)

type server struct {
	pb.UnimplementedNemServer
	core       *core.Implementation
	monitoring *monitoring.Implementation
}

type Params struct {
	fx.In
	Logger     *zap.Logger
	Lifecycle  fx.Lifecycle
	Core       *core.Implementation
	Config     config.Provider
	Monitoring *monitoring.Implementation

	Auth auth.Interface
}

func NewServer(params Params) pb.NemServer {
	return &server{
		core:       params.Core,
		monitoring: params.Monitoring,
	}
}

func New(params Params) *grpc.Server {
	log := params.Logger
	s := grpc.NewServer(

		grpc.UnaryInterceptor(AuthUnaryServerInterceptor(params.Auth, params.Monitoring)),
	)
	params.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// GRPC port from config
			grpcPort := params.Config.Get("ports.grpc").String()

			// proto server
			lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
			if err != nil {
				log.Error("failed to listen: %v", zap.Error(err))
				return err
			}

			// register grpc servers
			grpc_health_v1.RegisterHealthServer(s, health.NewServer())
			pb.RegisterNemServer(s, NewServer(params))
			reflection.Register(s)

			log.Info("GRPC Server listening at %v", zap.Any("addr", lis.Addr()))

			go s.Serve(lis)

			return nil

		},
		OnStop: func(ctx context.Context) error {
			s.Stop()
			return nil
		},
	})

	return s
}
