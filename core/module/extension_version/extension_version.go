package extension_version

import (
	"context"
	"nem/core/module/extension_version/types"
	"nem/core/repository"
	coretypes "nem/core/types"
	"nem/monitoring"
	"sync"

	"nem/core/events"
)

type Module interface {
	FetchExtensionVersionByUUID(ctx context.Context, req types.FetchExtensionVersionByUUIDRequest, opts ...Option) (types.FetchExtensionVersionByUUIDResponse, error)
	FetchExtensionVersionByVersion(ctx context.Context, req types.FetchExtensionVersionByVersionRequest, opts ...Option) (types.FetchExtensionVersionByVersionResponse, error)
	FetchExtensionVersionByStatus(ctx context.Context, req types.FetchExtensionVersionByStatusRequest, opts ...Option) (types.FetchExtensionVersionByStatusResponse, error)
	FetchExtensionVersionByVersionAndStatus(ctx context.Context, req types.FetchExtensionVersionByVersionAndStatusRequest, opts ...Option) (types.FetchExtensionVersionByVersionAndStatusResponse, error)

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
