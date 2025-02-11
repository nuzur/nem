package types

import (
	"nem/core/repository"

	"nem/core/events"

	"nem/monitoring"
)

type ModuleParams struct {
	Repository *repository.Implementation
	Monitoring *monitoring.Implementation

	Events *events.Implementation
}
