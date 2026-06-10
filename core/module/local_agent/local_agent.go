package local_agent

import (
	"context"
	"github.com/nuzur/nem/core/module/local_agent/types"
	"github.com/nuzur/nem/core/repository"
	coretypes "github.com/nuzur/nem/core/types"
	"github.com/nuzur/nem/monitoring"
	"sync"
)

type Module interface {
	FetchLocalAgentByUUID(ctx context.Context, req types.FetchLocalAgentByUUIDRequest, opts ...Option) (types.FetchLocalAgentByUUIDResponse, error)
	FetchLocalAgentByStatus(ctx context.Context, req types.FetchLocalAgentByStatusRequest, opts ...Option) (types.FetchLocalAgentByStatusResponse, error)

	List(ctx context.Context, req types.ListRequest, opts ...Option) (types.ListResponse, error)

	Upsert(ctx context.Context, req types.UpsertRequest, partial bool, opts ...Option) (types.UpsertResponse, error)
	Insert(ctx context.Context, req types.UpsertRequest, opts ...Option) (types.UpsertResponse, error)
	Update(ctx context.Context, req types.UpsertRequest, partial bool, opts ...Option) (types.UpsertResponse, error)
}

type module struct {
	mu         sync.Mutex
	repository *repository.Implementation
	monitoring *monitoring.Implementation
}

func New(params coretypes.ModuleParams) Module {
	return &module{
		repository: params.Repository,
		monitoring: params.Monitoring,
	}
}
