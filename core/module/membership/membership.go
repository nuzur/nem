package membership

import (
	"context"
	"github.com/nuzur/nem/core/module/membership/types"
	"github.com/nuzur/nem/core/repository"
	coretypes "github.com/nuzur/nem/core/types"
	"github.com/nuzur/nem/monitoring"
	"sync"

	"github.com/nuzur/nem/core/events"
)

type Module interface {
	FetchMembershipByUUID(ctx context.Context, req types.FetchMembershipByUUIDRequest, opts ...Option) (types.FetchMembershipByUUIDResponse, error)
	FetchMembershipByType(ctx context.Context, req types.FetchMembershipByTypeRequest, opts ...Option) (types.FetchMembershipByTypeResponse, error)
	FetchMembershipByStatus(ctx context.Context, req types.FetchMembershipByStatusRequest, opts ...Option) (types.FetchMembershipByStatusResponse, error)
	FetchMembershipByTypeAndStatus(ctx context.Context, req types.FetchMembershipByTypeAndStatusRequest, opts ...Option) (types.FetchMembershipByTypeAndStatusResponse, error)

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
