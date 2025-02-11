package project

import (
	"context"
	"nem/core/module/project/types"
	"nem/core/repository"
	coretypes "nem/core/types"
	"nem/monitoring"
	"sync"

	"nem/core/events"
)

type Module interface {
	FetchProjectByUUID(ctx context.Context, req types.FetchProjectByUUIDRequest, opts ...Option) (types.FetchProjectByUUIDResponse, error)
	FetchProjectByVersion(ctx context.Context, req types.FetchProjectByVersionRequest, opts ...Option) (types.FetchProjectByVersionResponse, error)
	FetchProjectByName(ctx context.Context, req types.FetchProjectByNameRequest, opts ...Option) (types.FetchProjectByNameResponse, error)
	FetchProjectByVersionAndName(ctx context.Context, req types.FetchProjectByVersionAndNameRequest, opts ...Option) (types.FetchProjectByVersionAndNameResponse, error)
	FetchProjectByStatus(ctx context.Context, req types.FetchProjectByStatusRequest, opts ...Option) (types.FetchProjectByStatusResponse, error)
	FetchProjectByVersionAndStatus(ctx context.Context, req types.FetchProjectByVersionAndStatusRequest, opts ...Option) (types.FetchProjectByVersionAndStatusResponse, error)
	FetchProjectByNameAndStatus(ctx context.Context, req types.FetchProjectByNameAndStatusRequest, opts ...Option) (types.FetchProjectByNameAndStatusResponse, error)
	FetchProjectByVersionAndNameAndStatus(ctx context.Context, req types.FetchProjectByVersionAndNameAndStatusRequest, opts ...Option) (types.FetchProjectByVersionAndNameAndStatusResponse, error)

	List(ctx context.Context, req types.ListRequest, opts ...Option) (types.ListResponse, error)

	Upsert(ctx context.Context, req types.UpsertRequest, partial bool, opts ...Option) (types.UpsertResponse, error)
	Insert(ctx context.Context, req types.UpsertRequest, opts ...Option) (types.UpsertResponse, error)
	Update(ctx context.Context, req types.UpsertRequest, partial bool, opts ...Option) (types.UpsertResponse, error)

	Search(ctx context.Context, query string, limit int32, offset int32) (types.SearchProjectResponse, error)
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
