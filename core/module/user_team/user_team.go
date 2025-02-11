package user_team

import (
	"context"
	"github.com/nuzur/nem/core/module/user_team/types"
	"github.com/nuzur/nem/core/repository"
	coretypes "github.com/nuzur/nem/core/types"
	"github.com/nuzur/nem/monitoring"
	"sync"

	"github.com/nuzur/nem/core/events"
)

type Module interface {
	FetchUserTeamByUUID(ctx context.Context, req types.FetchUserTeamByUUIDRequest, opts ...Option) (types.FetchUserTeamByUUIDResponse, error)
	FetchUserTeamByUserEmail(ctx context.Context, req types.FetchUserTeamByUserEmailRequest, opts ...Option) (types.FetchUserTeamByUserEmailResponse, error)
	FetchUserTeamByStatus(ctx context.Context, req types.FetchUserTeamByStatusRequest, opts ...Option) (types.FetchUserTeamByStatusResponse, error)
	FetchUserTeamByUserEmailAndStatus(ctx context.Context, req types.FetchUserTeamByUserEmailAndStatusRequest, opts ...Option) (types.FetchUserTeamByUserEmailAndStatusResponse, error)

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
