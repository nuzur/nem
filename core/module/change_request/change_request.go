package change_request

import (
	"context"
	"nem/core/module/change_request/types"
	"nem/core/repository"
	coretypes "nem/core/types"
	"nem/monitoring"
	"sync"

	"nem/core/events"
)

type Module interface {
	FetchChangeRequestByUUID(ctx context.Context, req types.FetchChangeRequestByUUIDRequest, opts ...Option) (types.FetchChangeRequestByUUIDResponse, error)
	FetchChangeRequestByVersion(ctx context.Context, req types.FetchChangeRequestByVersionRequest, opts ...Option) (types.FetchChangeRequestByVersionResponse, error)
	FetchChangeRequestByStatus(ctx context.Context, req types.FetchChangeRequestByStatusRequest, opts ...Option) (types.FetchChangeRequestByStatusResponse, error)
	FetchChangeRequestByVersionAndStatus(ctx context.Context, req types.FetchChangeRequestByVersionAndStatusRequest, opts ...Option) (types.FetchChangeRequestByVersionAndStatusResponse, error)

	List(ctx context.Context, req types.ListRequest, opts ...Option) (types.ListResponse, error)

	Upsert(ctx context.Context, req types.UpsertRequest, partial bool, opts ...Option) (types.UpsertResponse, error)
	Insert(ctx context.Context, req types.UpsertRequest, opts ...Option) (types.UpsertResponse, error)
	Update(ctx context.Context, req types.UpsertRequest, partial bool, opts ...Option) (types.UpsertResponse, error)

	Search(ctx context.Context, query string, limit int32, offset int32) (types.SearchChangeRequestResponse, error)
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
