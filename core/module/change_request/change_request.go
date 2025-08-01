package change_request

import (
	"context"
	"github.com/nuzur/nem/core/module/change_request/types"
	"github.com/nuzur/nem/core/repository"
	coretypes "github.com/nuzur/nem/core/types"
	"github.com/nuzur/nem/monitoring"
	"sync"

	"github.com/nuzur/nem/core/events"
)

type Module interface {
	FetchChangeRequestByUUID(ctx context.Context, req types.FetchChangeRequestByUUIDRequest, opts ...Option) (types.FetchChangeRequestByUUIDResponse, error)
	FetchChangeRequestByVersion(ctx context.Context, req types.FetchChangeRequestByVersionRequest, opts ...Option) (types.FetchChangeRequestByVersionResponse, error)
	FetchChangeRequestByChangeType(ctx context.Context, req types.FetchChangeRequestByChangeTypeRequest, opts ...Option) (types.FetchChangeRequestByChangeTypeResponse, error)
	FetchChangeRequestByVersionAndChangeType(ctx context.Context, req types.FetchChangeRequestByVersionAndChangeTypeRequest, opts ...Option) (types.FetchChangeRequestByVersionAndChangeTypeResponse, error)
	FetchChangeRequestByReviewStatus(ctx context.Context, req types.FetchChangeRequestByReviewStatusRequest, opts ...Option) (types.FetchChangeRequestByReviewStatusResponse, error)
	FetchChangeRequestByVersionAndReviewStatus(ctx context.Context, req types.FetchChangeRequestByVersionAndReviewStatusRequest, opts ...Option) (types.FetchChangeRequestByVersionAndReviewStatusResponse, error)
	FetchChangeRequestByChangeTypeAndReviewStatus(ctx context.Context, req types.FetchChangeRequestByChangeTypeAndReviewStatusRequest, opts ...Option) (types.FetchChangeRequestByChangeTypeAndReviewStatusResponse, error)
	FetchChangeRequestByVersionAndChangeTypeAndReviewStatus(ctx context.Context, req types.FetchChangeRequestByVersionAndChangeTypeAndReviewStatusRequest, opts ...Option) (types.FetchChangeRequestByVersionAndChangeTypeAndReviewStatusResponse, error)
	FetchChangeRequestByStatus(ctx context.Context, req types.FetchChangeRequestByStatusRequest, opts ...Option) (types.FetchChangeRequestByStatusResponse, error)
	FetchChangeRequestByVersionAndStatus(ctx context.Context, req types.FetchChangeRequestByVersionAndStatusRequest, opts ...Option) (types.FetchChangeRequestByVersionAndStatusResponse, error)
	FetchChangeRequestByChangeTypeAndStatus(ctx context.Context, req types.FetchChangeRequestByChangeTypeAndStatusRequest, opts ...Option) (types.FetchChangeRequestByChangeTypeAndStatusResponse, error)
	FetchChangeRequestByVersionAndChangeTypeAndStatus(ctx context.Context, req types.FetchChangeRequestByVersionAndChangeTypeAndStatusRequest, opts ...Option) (types.FetchChangeRequestByVersionAndChangeTypeAndStatusResponse, error)
	FetchChangeRequestByReviewStatusAndStatus(ctx context.Context, req types.FetchChangeRequestByReviewStatusAndStatusRequest, opts ...Option) (types.FetchChangeRequestByReviewStatusAndStatusResponse, error)
	FetchChangeRequestByVersionAndReviewStatusAndStatus(ctx context.Context, req types.FetchChangeRequestByVersionAndReviewStatusAndStatusRequest, opts ...Option) (types.FetchChangeRequestByVersionAndReviewStatusAndStatusResponse, error)
	FetchChangeRequestByChangeTypeAndReviewStatusAndStatus(ctx context.Context, req types.FetchChangeRequestByChangeTypeAndReviewStatusAndStatusRequest, opts ...Option) (types.FetchChangeRequestByChangeTypeAndReviewStatusAndStatusResponse, error)
	FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndStatus(ctx context.Context, req types.FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndStatusRequest, opts ...Option) (types.FetchChangeRequestByVersionAndChangeTypeAndReviewStatusAndStatusResponse, error)

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
