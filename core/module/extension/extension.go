package extension

import (
	"context"
	"nem/core/module/extension/types"
	"nem/core/repository"
	coretypes "nem/core/types"
	"nem/monitoring"
	"sync"

	"nem/core/events"
)

type Module interface {
	FetchExtensionByUUID(ctx context.Context, req types.FetchExtensionByUUIDRequest, opts ...Option) (types.FetchExtensionByUUIDResponse, error)
	FetchExtensionByVersion(ctx context.Context, req types.FetchExtensionByVersionRequest, opts ...Option) (types.FetchExtensionByVersionResponse, error)
	FetchExtensionByIdentifier(ctx context.Context, req types.FetchExtensionByIdentifierRequest, opts ...Option) (types.FetchExtensionByIdentifierResponse, error)
	FetchExtensionByVersionAndIdentifier(ctx context.Context, req types.FetchExtensionByVersionAndIdentifierRequest, opts ...Option) (types.FetchExtensionByVersionAndIdentifierResponse, error)
	FetchExtensionByRepository(ctx context.Context, req types.FetchExtensionByRepositoryRequest, opts ...Option) (types.FetchExtensionByRepositoryResponse, error)
	FetchExtensionByVersionAndRepository(ctx context.Context, req types.FetchExtensionByVersionAndRepositoryRequest, opts ...Option) (types.FetchExtensionByVersionAndRepositoryResponse, error)
	FetchExtensionByIdentifierAndRepository(ctx context.Context, req types.FetchExtensionByIdentifierAndRepositoryRequest, opts ...Option) (types.FetchExtensionByIdentifierAndRepositoryResponse, error)
	FetchExtensionByVersionAndIdentifierAndRepository(ctx context.Context, req types.FetchExtensionByVersionAndIdentifierAndRepositoryRequest, opts ...Option) (types.FetchExtensionByVersionAndIdentifierAndRepositoryResponse, error)
	FetchExtensionByExtensionType(ctx context.Context, req types.FetchExtensionByExtensionTypeRequest, opts ...Option) (types.FetchExtensionByExtensionTypeResponse, error)
	FetchExtensionByVersionAndExtensionType(ctx context.Context, req types.FetchExtensionByVersionAndExtensionTypeRequest, opts ...Option) (types.FetchExtensionByVersionAndExtensionTypeResponse, error)
	FetchExtensionByIdentifierAndExtensionType(ctx context.Context, req types.FetchExtensionByIdentifierAndExtensionTypeRequest, opts ...Option) (types.FetchExtensionByIdentifierAndExtensionTypeResponse, error)
	FetchExtensionByVersionAndIdentifierAndExtensionType(ctx context.Context, req types.FetchExtensionByVersionAndIdentifierAndExtensionTypeRequest, opts ...Option) (types.FetchExtensionByVersionAndIdentifierAndExtensionTypeResponse, error)
	FetchExtensionByRepositoryAndExtensionType(ctx context.Context, req types.FetchExtensionByRepositoryAndExtensionTypeRequest, opts ...Option) (types.FetchExtensionByRepositoryAndExtensionTypeResponse, error)
	FetchExtensionByVersionAndRepositoryAndExtensionType(ctx context.Context, req types.FetchExtensionByVersionAndRepositoryAndExtensionTypeRequest, opts ...Option) (types.FetchExtensionByVersionAndRepositoryAndExtensionTypeResponse, error)
	FetchExtensionByIdentifierAndRepositoryAndExtensionType(ctx context.Context, req types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeRequest, opts ...Option) (types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeResponse, error)
	FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionType(ctx context.Context, req types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeRequest, opts ...Option) (types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeResponse, error)
	FetchExtensionByPublic(ctx context.Context, req types.FetchExtensionByPublicRequest, opts ...Option) (types.FetchExtensionByPublicResponse, error)
	FetchExtensionByVersionAndPublic(ctx context.Context, req types.FetchExtensionByVersionAndPublicRequest, opts ...Option) (types.FetchExtensionByVersionAndPublicResponse, error)
	FetchExtensionByIdentifierAndPublic(ctx context.Context, req types.FetchExtensionByIdentifierAndPublicRequest, opts ...Option) (types.FetchExtensionByIdentifierAndPublicResponse, error)
	FetchExtensionByVersionAndIdentifierAndPublic(ctx context.Context, req types.FetchExtensionByVersionAndIdentifierAndPublicRequest, opts ...Option) (types.FetchExtensionByVersionAndIdentifierAndPublicResponse, error)
	FetchExtensionByRepositoryAndPublic(ctx context.Context, req types.FetchExtensionByRepositoryAndPublicRequest, opts ...Option) (types.FetchExtensionByRepositoryAndPublicResponse, error)
	FetchExtensionByVersionAndRepositoryAndPublic(ctx context.Context, req types.FetchExtensionByVersionAndRepositoryAndPublicRequest, opts ...Option) (types.FetchExtensionByVersionAndRepositoryAndPublicResponse, error)
	FetchExtensionByIdentifierAndRepositoryAndPublic(ctx context.Context, req types.FetchExtensionByIdentifierAndRepositoryAndPublicRequest, opts ...Option) (types.FetchExtensionByIdentifierAndRepositoryAndPublicResponse, error)
	FetchExtensionByVersionAndIdentifierAndRepositoryAndPublic(ctx context.Context, req types.FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicRequest, opts ...Option) (types.FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicResponse, error)
	FetchExtensionByExtensionTypeAndPublic(ctx context.Context, req types.FetchExtensionByExtensionTypeAndPublicRequest, opts ...Option) (types.FetchExtensionByExtensionTypeAndPublicResponse, error)
	FetchExtensionByVersionAndExtensionTypeAndPublic(ctx context.Context, req types.FetchExtensionByVersionAndExtensionTypeAndPublicRequest, opts ...Option) (types.FetchExtensionByVersionAndExtensionTypeAndPublicResponse, error)
	FetchExtensionByIdentifierAndExtensionTypeAndPublic(ctx context.Context, req types.FetchExtensionByIdentifierAndExtensionTypeAndPublicRequest, opts ...Option) (types.FetchExtensionByIdentifierAndExtensionTypeAndPublicResponse, error)
	FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublic(ctx context.Context, req types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicRequest, opts ...Option) (types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicResponse, error)
	FetchExtensionByRepositoryAndExtensionTypeAndPublic(ctx context.Context, req types.FetchExtensionByRepositoryAndExtensionTypeAndPublicRequest, opts ...Option) (types.FetchExtensionByRepositoryAndExtensionTypeAndPublicResponse, error)
	FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublic(ctx context.Context, req types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicRequest, opts ...Option) (types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicResponse, error)
	FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublic(ctx context.Context, req types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicRequest, opts ...Option) (types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicResponse, error)
	FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublic(ctx context.Context, req types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicRequest, opts ...Option) (types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicResponse, error)
	FetchExtensionByStatus(ctx context.Context, req types.FetchExtensionByStatusRequest, opts ...Option) (types.FetchExtensionByStatusResponse, error)
	FetchExtensionByVersionAndStatus(ctx context.Context, req types.FetchExtensionByVersionAndStatusRequest, opts ...Option) (types.FetchExtensionByVersionAndStatusResponse, error)
	FetchExtensionByIdentifierAndStatus(ctx context.Context, req types.FetchExtensionByIdentifierAndStatusRequest, opts ...Option) (types.FetchExtensionByIdentifierAndStatusResponse, error)
	FetchExtensionByVersionAndIdentifierAndStatus(ctx context.Context, req types.FetchExtensionByVersionAndIdentifierAndStatusRequest, opts ...Option) (types.FetchExtensionByVersionAndIdentifierAndStatusResponse, error)
	FetchExtensionByRepositoryAndStatus(ctx context.Context, req types.FetchExtensionByRepositoryAndStatusRequest, opts ...Option) (types.FetchExtensionByRepositoryAndStatusResponse, error)
	FetchExtensionByVersionAndRepositoryAndStatus(ctx context.Context, req types.FetchExtensionByVersionAndRepositoryAndStatusRequest, opts ...Option) (types.FetchExtensionByVersionAndRepositoryAndStatusResponse, error)
	FetchExtensionByIdentifierAndRepositoryAndStatus(ctx context.Context, req types.FetchExtensionByIdentifierAndRepositoryAndStatusRequest, opts ...Option) (types.FetchExtensionByIdentifierAndRepositoryAndStatusResponse, error)
	FetchExtensionByVersionAndIdentifierAndRepositoryAndStatus(ctx context.Context, req types.FetchExtensionByVersionAndIdentifierAndRepositoryAndStatusRequest, opts ...Option) (types.FetchExtensionByVersionAndIdentifierAndRepositoryAndStatusResponse, error)
	FetchExtensionByExtensionTypeAndStatus(ctx context.Context, req types.FetchExtensionByExtensionTypeAndStatusRequest, opts ...Option) (types.FetchExtensionByExtensionTypeAndStatusResponse, error)
	FetchExtensionByVersionAndExtensionTypeAndStatus(ctx context.Context, req types.FetchExtensionByVersionAndExtensionTypeAndStatusRequest, opts ...Option) (types.FetchExtensionByVersionAndExtensionTypeAndStatusResponse, error)
	FetchExtensionByIdentifierAndExtensionTypeAndStatus(ctx context.Context, req types.FetchExtensionByIdentifierAndExtensionTypeAndStatusRequest, opts ...Option) (types.FetchExtensionByIdentifierAndExtensionTypeAndStatusResponse, error)
	FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatus(ctx context.Context, req types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatusRequest, opts ...Option) (types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndStatusResponse, error)
	FetchExtensionByRepositoryAndExtensionTypeAndStatus(ctx context.Context, req types.FetchExtensionByRepositoryAndExtensionTypeAndStatusRequest, opts ...Option) (types.FetchExtensionByRepositoryAndExtensionTypeAndStatusResponse, error)
	FetchExtensionByVersionAndRepositoryAndExtensionTypeAndStatus(ctx context.Context, req types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndStatusRequest, opts ...Option) (types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndStatusResponse, error)
	FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatus(ctx context.Context, req types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatusRequest, opts ...Option) (types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndStatusResponse, error)
	FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatus(ctx context.Context, req types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatusRequest, opts ...Option) (types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndStatusResponse, error)
	FetchExtensionByPublicAndStatus(ctx context.Context, req types.FetchExtensionByPublicAndStatusRequest, opts ...Option) (types.FetchExtensionByPublicAndStatusResponse, error)
	FetchExtensionByVersionAndPublicAndStatus(ctx context.Context, req types.FetchExtensionByVersionAndPublicAndStatusRequest, opts ...Option) (types.FetchExtensionByVersionAndPublicAndStatusResponse, error)
	FetchExtensionByIdentifierAndPublicAndStatus(ctx context.Context, req types.FetchExtensionByIdentifierAndPublicAndStatusRequest, opts ...Option) (types.FetchExtensionByIdentifierAndPublicAndStatusResponse, error)
	FetchExtensionByVersionAndIdentifierAndPublicAndStatus(ctx context.Context, req types.FetchExtensionByVersionAndIdentifierAndPublicAndStatusRequest, opts ...Option) (types.FetchExtensionByVersionAndIdentifierAndPublicAndStatusResponse, error)
	FetchExtensionByRepositoryAndPublicAndStatus(ctx context.Context, req types.FetchExtensionByRepositoryAndPublicAndStatusRequest, opts ...Option) (types.FetchExtensionByRepositoryAndPublicAndStatusResponse, error)
	FetchExtensionByVersionAndRepositoryAndPublicAndStatus(ctx context.Context, req types.FetchExtensionByVersionAndRepositoryAndPublicAndStatusRequest, opts ...Option) (types.FetchExtensionByVersionAndRepositoryAndPublicAndStatusResponse, error)
	FetchExtensionByIdentifierAndRepositoryAndPublicAndStatus(ctx context.Context, req types.FetchExtensionByIdentifierAndRepositoryAndPublicAndStatusRequest, opts ...Option) (types.FetchExtensionByIdentifierAndRepositoryAndPublicAndStatusResponse, error)
	FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatus(ctx context.Context, req types.FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatusRequest, opts ...Option) (types.FetchExtensionByVersionAndIdentifierAndRepositoryAndPublicAndStatusResponse, error)
	FetchExtensionByExtensionTypeAndPublicAndStatus(ctx context.Context, req types.FetchExtensionByExtensionTypeAndPublicAndStatusRequest, opts ...Option) (types.FetchExtensionByExtensionTypeAndPublicAndStatusResponse, error)
	FetchExtensionByVersionAndExtensionTypeAndPublicAndStatus(ctx context.Context, req types.FetchExtensionByVersionAndExtensionTypeAndPublicAndStatusRequest, opts ...Option) (types.FetchExtensionByVersionAndExtensionTypeAndPublicAndStatusResponse, error)
	FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatus(ctx context.Context, req types.FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatusRequest, opts ...Option) (types.FetchExtensionByIdentifierAndExtensionTypeAndPublicAndStatusResponse, error)
	FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicAndStatus(ctx context.Context, req types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicAndStatusRequest, opts ...Option) (types.FetchExtensionByVersionAndIdentifierAndExtensionTypeAndPublicAndStatusResponse, error)
	FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatus(ctx context.Context, req types.FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatusRequest, opts ...Option) (types.FetchExtensionByRepositoryAndExtensionTypeAndPublicAndStatusResponse, error)
	FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatus(ctx context.Context, req types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatusRequest, opts ...Option) (types.FetchExtensionByVersionAndRepositoryAndExtensionTypeAndPublicAndStatusResponse, error)
	FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatus(ctx context.Context, req types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusRequest, opts ...Option) (types.FetchExtensionByIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusResponse, error)
	FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatus(ctx context.Context, req types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusRequest, opts ...Option) (types.FetchExtensionByVersionAndIdentifierAndRepositoryAndExtensionTypeAndPublicAndStatusResponse, error)

	List(ctx context.Context, req types.ListRequest, opts ...Option) (types.ListResponse, error)

	Upsert(ctx context.Context, req types.UpsertRequest, partial bool, opts ...Option) (types.UpsertResponse, error)
	Insert(ctx context.Context, req types.UpsertRequest, opts ...Option) (types.UpsertResponse, error)
	Update(ctx context.Context, req types.UpsertRequest, partial bool, opts ...Option) (types.UpsertResponse, error)

	Search(ctx context.Context, query string, limit int32, offset int32) (types.SearchExtensionResponse, error)
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
