package organization

import (
	"context"
	"github.com/nuzur/nem/core/module/organization/types"
	"github.com/nuzur/nem/core/repository"
	coretypes "github.com/nuzur/nem/core/types"
	"github.com/nuzur/nem/monitoring"
	"sync"

	"github.com/nuzur/nem/core/events"
)

type Module interface {
	FetchOrganizationByUUID(ctx context.Context, req types.FetchOrganizationByUUIDRequest, opts ...Option) (types.FetchOrganizationByUUIDResponse, error)
	FetchOrganizationByVersion(ctx context.Context, req types.FetchOrganizationByVersionRequest, opts ...Option) (types.FetchOrganizationByVersionResponse, error)
	FetchOrganizationByStatus(ctx context.Context, req types.FetchOrganizationByStatusRequest, opts ...Option) (types.FetchOrganizationByStatusResponse, error)
	FetchOrganizationByVersionAndStatus(ctx context.Context, req types.FetchOrganizationByVersionAndStatusRequest, opts ...Option) (types.FetchOrganizationByVersionAndStatusResponse, error)

	List(ctx context.Context, req types.ListRequest, opts ...Option) (types.ListResponse, error)

	Upsert(ctx context.Context, req types.UpsertRequest, partial bool, opts ...Option) (types.UpsertResponse, error)
	Insert(ctx context.Context, req types.UpsertRequest, opts ...Option) (types.UpsertResponse, error)
	Update(ctx context.Context, req types.UpsertRequest, partial bool, opts ...Option) (types.UpsertResponse, error)

	Search(ctx context.Context, query string, limit int32, offset int32) (types.SearchOrganizationResponse, error)
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
