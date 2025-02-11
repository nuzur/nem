package extension_execution

import (
	"context"
	"nem/core/module/extension_execution/types"
	"nem/core/repository"
	coretypes "nem/core/types"
	"nem/monitoring"
	"sync"

	"nem/core/events"
)

type Module interface {
	FetchExtensionExecutionByUUID(ctx context.Context, req types.FetchExtensionExecutionByUUIDRequest, opts ...Option) (types.FetchExtensionExecutionByUUIDResponse, error)
	FetchExtensionExecutionByStatus(ctx context.Context, req types.FetchExtensionExecutionByStatusRequest, opts ...Option) (types.FetchExtensionExecutionByStatusResponse, error)

	List(ctx context.Context, req types.ListRequest, opts ...Option) (types.ListResponse, error)

	Upsert(ctx context.Context, req types.UpsertRequest, partial bool, opts ...Option) (types.UpsertResponse, error)
	Insert(ctx context.Context, req types.UpsertRequest, opts ...Option) (types.UpsertResponse, error)
	Update(ctx context.Context, req types.UpsertRequest, partial bool, opts ...Option) (types.UpsertResponse, error)
}

type module struct {
	mu         sync.Mutex
	repository *repository.Implementation
	monitoring *monitoring.Implementation

	events *events.Implementation
}

func New(params coretypes.ModuleParams) Module {
	return &module{
		repository: params.Repository,
		monitoring: params.Monitoring,

		events: params.Events,
	}
}
