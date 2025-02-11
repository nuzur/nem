package project_version

import (
	"context"
	"github.com/nuzur/nem/core/module/project_version/types"
	"github.com/nuzur/nem/core/repository"
	coretypes "github.com/nuzur/nem/core/types"
	"github.com/nuzur/nem/monitoring"
	"sync"

	"github.com/nuzur/nem/core/events"
)

type Module interface {
	FetchProjectVersionByUUID(ctx context.Context, req types.FetchProjectVersionByUUIDRequest, opts ...Option) (types.FetchProjectVersionByUUIDResponse, error)
	FetchProjectVersionByVersion(ctx context.Context, req types.FetchProjectVersionByVersionRequest, opts ...Option) (types.FetchProjectVersionByVersionResponse, error)
	FetchProjectVersionByReviewStatus(ctx context.Context, req types.FetchProjectVersionByReviewStatusRequest, opts ...Option) (types.FetchProjectVersionByReviewStatusResponse, error)
	FetchProjectVersionByVersionAndReviewStatus(ctx context.Context, req types.FetchProjectVersionByVersionAndReviewStatusRequest, opts ...Option) (types.FetchProjectVersionByVersionAndReviewStatusResponse, error)
	FetchProjectVersionByStatus(ctx context.Context, req types.FetchProjectVersionByStatusRequest, opts ...Option) (types.FetchProjectVersionByStatusResponse, error)
	FetchProjectVersionByVersionAndStatus(ctx context.Context, req types.FetchProjectVersionByVersionAndStatusRequest, opts ...Option) (types.FetchProjectVersionByVersionAndStatusResponse, error)
	FetchProjectVersionByReviewStatusAndStatus(ctx context.Context, req types.FetchProjectVersionByReviewStatusAndStatusRequest, opts ...Option) (types.FetchProjectVersionByReviewStatusAndStatusResponse, error)
	FetchProjectVersionByVersionAndReviewStatusAndStatus(ctx context.Context, req types.FetchProjectVersionByVersionAndReviewStatusAndStatusRequest, opts ...Option) (types.FetchProjectVersionByVersionAndReviewStatusAndStatusResponse, error)

	List(ctx context.Context, req types.ListRequest, opts ...Option) (types.ListResponse, error)

	Upsert(ctx context.Context, req types.UpsertRequest, partial bool, opts ...Option) (types.UpsertResponse, error)
	Insert(ctx context.Context, req types.UpsertRequest, opts ...Option) (types.UpsertResponse, error)
	Update(ctx context.Context, req types.UpsertRequest, partial bool, opts ...Option) (types.UpsertResponse, error)

	Search(ctx context.Context, query string, limit int32, offset int32) (types.SearchProjectVersionResponse, error)
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
