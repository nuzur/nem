package user

import (
	"context"
	"github.com/nuzur/nem/core/module/user/types"
	"github.com/nuzur/nem/core/repository"
	coretypes "github.com/nuzur/nem/core/types"
	"github.com/nuzur/nem/monitoring"
	"sync"

	"github.com/nuzur/nem/core/events"
)

type Module interface {
	FetchUserByUUID(ctx context.Context, req types.FetchUserByUUIDRequest, opts ...Option) (types.FetchUserByUUIDResponse, error)
	FetchUserByIdentifier(ctx context.Context, req types.FetchUserByIdentifierRequest, opts ...Option) (types.FetchUserByIdentifierResponse, error)
	FetchUserByEmail(ctx context.Context, req types.FetchUserByEmailRequest, opts ...Option) (types.FetchUserByEmailResponse, error)
	FetchUserByIdentifierAndEmail(ctx context.Context, req types.FetchUserByIdentifierAndEmailRequest, opts ...Option) (types.FetchUserByIdentifierAndEmailResponse, error)
	FetchUserByStatus(ctx context.Context, req types.FetchUserByStatusRequest, opts ...Option) (types.FetchUserByStatusResponse, error)
	FetchUserByIdentifierAndStatus(ctx context.Context, req types.FetchUserByIdentifierAndStatusRequest, opts ...Option) (types.FetchUserByIdentifierAndStatusResponse, error)
	FetchUserByEmailAndStatus(ctx context.Context, req types.FetchUserByEmailAndStatusRequest, opts ...Option) (types.FetchUserByEmailAndStatusResponse, error)
	FetchUserByIdentifierAndEmailAndStatus(ctx context.Context, req types.FetchUserByIdentifierAndEmailAndStatusRequest, opts ...Option) (types.FetchUserByIdentifierAndEmailAndStatusResponse, error)

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
