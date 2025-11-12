package ai_usage

import (
	"context"
	"github.com/nuzur/nem/core/module/ai_usage/types"
	"github.com/nuzur/nem/core/repository"
	coretypes "github.com/nuzur/nem/core/types"
	"github.com/nuzur/nem/monitoring"
	"sync"

	"github.com/nuzur/nem/core/events"
)

type Module interface {
	FetchAiUsageByUUID(ctx context.Context, req types.FetchAiUsageByUUIDRequest, opts ...Option) (types.FetchAiUsageByUUIDResponse, error)
	FetchAiUsageByContext(ctx context.Context, req types.FetchAiUsageByContextRequest, opts ...Option) (types.FetchAiUsageByContextResponse, error)
	FetchAiUsageByProvider(ctx context.Context, req types.FetchAiUsageByProviderRequest, opts ...Option) (types.FetchAiUsageByProviderResponse, error)
	FetchAiUsageByContextAndProvider(ctx context.Context, req types.FetchAiUsageByContextAndProviderRequest, opts ...Option) (types.FetchAiUsageByContextAndProviderResponse, error)
	FetchAiUsageByStatus(ctx context.Context, req types.FetchAiUsageByStatusRequest, opts ...Option) (types.FetchAiUsageByStatusResponse, error)
	FetchAiUsageByContextAndStatus(ctx context.Context, req types.FetchAiUsageByContextAndStatusRequest, opts ...Option) (types.FetchAiUsageByContextAndStatusResponse, error)
	FetchAiUsageByProviderAndStatus(ctx context.Context, req types.FetchAiUsageByProviderAndStatusRequest, opts ...Option) (types.FetchAiUsageByProviderAndStatusResponse, error)
	FetchAiUsageByContextAndProviderAndStatus(ctx context.Context, req types.FetchAiUsageByContextAndProviderAndStatusRequest, opts ...Option) (types.FetchAiUsageByContextAndProviderAndStatusResponse, error)

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
