package main

import (
	nemconfig "github.com/nuzur/nem/config"
	"github.com/nuzur/nem/core"
	"go.uber.org/fx"
	"go.uber.org/zap"

	pbserver "github.com/nuzur/nem/idl/server"

	auth "github.com/nuzur/nem/auth/keycloak"

	"github.com/nuzur/nem/monitoring"
)

func main() {
	fx.New(
		fx.Provide(
			zap.NewProduction,
			nemconfig.New,
			core.New,
			monitoring.New,
			auth.New,
		),

		fx.Invoke(httpServer),
		fx.Invoke(pbserver.New),
	).Run()
}
