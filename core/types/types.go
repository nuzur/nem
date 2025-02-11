package types

import (
	"github.com/nuzur/nem/core/repository"

	"github.com/nuzur/nem/core/events"

	"github.com/nuzur/nem/monitoring"
)

type ModuleParams struct {
	Repository *repository.Implementation
	Monitoring *monitoring.Implementation

	Events *events.Implementation
}
