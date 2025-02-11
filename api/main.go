package main

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	nemconfig "nem/config"
	"nem/core"

	pbserver "nem/idl/server"

	auth "nem/auth/keycloak"

	"nem/monitoring"

	"nem/core/events"

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
