package main

import (
	nemconfig "github.com/nuzur/nem/config"
	"github.com/nuzur/nem/core"
	"go.uber.org/fx"
	"go.uber.org/zap"

	pbserver "github.com/nuzur/nem/idl/server"

	auth "github.com/nuzur/nem/auth/keycloak"

	"github.com/nuzur/nem/monitoring"

	"github.com/nuzur/nem/core/events"

	saramafx "github.com/mklfarha/sarama-fx"
)

func main() {
	fx.New(
		fx.Provide(
			zap.NewProduction,
			nemconfig.New,
			core.New,
			monitoring.New,
			auth.New,
			events.New,
		),
		saramafx.Module,
		fx.Invoke(httpServer),
		fx.Invoke(pbserver.New),
	).Run()
}
