package team

import (
	"context"
	"nem/core/module/team/types"
	"nem/core/repository"
	coretypes "nem/core/types"
	"nem/monitoring"
	"sync"

	"nem/core/events"
)

type Module interface {
	FetchTeamByUUID(ctx context.Context, req types.FetchTeamByUUIDRequest, opts ...Option) (types.FetchTeamByUUIDResponse, error)
	FetchTeamByVersion(ctx context.Context, req types.FetchTeamByVersionRequest, opts ...Option) (types.FetchTeamByVersionResponse, error)
	FetchTeamByStatus(ctx context.Context, req types.FetchTeamByStatusRequest, opts ...Option) (types.FetchTeamByStatusResponse, error)
	FetchTeamByVersionAndStatus(ctx context.Context, req types.FetchTeamByVersionAndStatusRequest, opts ...Option) (types.FetchTeamByVersionAndStatusResponse, error)

	List(ctx context.Context, req types.ListRequest, opts ...Option) (types.ListResponse, error)

	Upsert(ctx context.Context, req types.UpsertRequest, partial bool, opts ...Option) (types.UpsertResponse, error)
	Insert(ctx context.Context, req types.UpsertRequest, opts ...Option) (types.UpsertResponse, error)
	Update(ctx context.Context, req types.UpsertRequest, partial bool, opts ...Option) (types.UpsertResponse, error)

	Search(ctx context.Context, query string, limit int32, offset int32) (types.SearchTeamResponse, error)
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
