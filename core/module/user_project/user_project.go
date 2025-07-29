package user_project

import (
	"context"
	"github.com/nuzur/nem/core/module/user_project/types"
	"github.com/nuzur/nem/core/repository"
	coretypes "github.com/nuzur/nem/core/types"
	"github.com/nuzur/nem/monitoring"
	"sync"

	"github.com/nuzur/nem/core/events"
)

type Module interface {
	FetchUserProjectByUUID(ctx context.Context, req types.FetchUserProjectByUUIDRequest, opts ...Option) (types.FetchUserProjectByUUIDResponse, error)
	FetchUserProjectByUserEmail(ctx context.Context, req types.FetchUserProjectByUserEmailRequest, opts ...Option) (types.FetchUserProjectByUserEmailResponse, error)
	FetchUserProjectByRole(ctx context.Context, req types.FetchUserProjectByRoleRequest, opts ...Option) (types.FetchUserProjectByRoleResponse, error)
	FetchUserProjectByUserEmailAndRole(ctx context.Context, req types.FetchUserProjectByUserEmailAndRoleRequest, opts ...Option) (types.FetchUserProjectByUserEmailAndRoleResponse, error)
	FetchUserProjectByStatus(ctx context.Context, req types.FetchUserProjectByStatusRequest, opts ...Option) (types.FetchUserProjectByStatusResponse, error)
	FetchUserProjectByUserEmailAndStatus(ctx context.Context, req types.FetchUserProjectByUserEmailAndStatusRequest, opts ...Option) (types.FetchUserProjectByUserEmailAndStatusResponse, error)
	FetchUserProjectByRoleAndStatus(ctx context.Context, req types.FetchUserProjectByRoleAndStatusRequest, opts ...Option) (types.FetchUserProjectByRoleAndStatusResponse, error)
	FetchUserProjectByUserEmailAndRoleAndStatus(ctx context.Context, req types.FetchUserProjectByUserEmailAndRoleAndStatusRequest, opts ...Option) (types.FetchUserProjectByUserEmailAndRoleAndStatusResponse, error)

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
