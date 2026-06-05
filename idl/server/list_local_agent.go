package server

import (
	"context"
	"fmt"
	//"encoding/json"

	"github.com/nuzur/nem/core/module/local_agent/types"
	pb "github.com/nuzur/nem/idl/gen"
	pbmapper "github.com/nuzur/nem/idl/mapper"
	"github.com/nuzur/nem/monitoring"

	"go.einride.tech/aip/filtering"
	"go.einride.tech/aip/ordering"
	"go.einride.tech/aip/pagination"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func BuildListLocalAgentsRequest(ctx context.Context, request *pb.ListLocalAgentsRequest) (types.ListRequest, *pagination.PageToken, error) {
	err := validatePageSizeForListLocalAgents(request)
	if err != nil {
		return types.ListRequest{}, nil, err
	}

	// Use pagination.PageToken for offset-based page tokens.
	pageToken, err := pagination.ParsePageToken(request)
	if err != nil {
		return types.ListRequest{}, nil, status.Errorf(codes.InvalidArgument, "invalid page token")
	}

	// parse filters
	declarations := local_agentDeclarations()
	filter, err := filtering.ParseFilter(request, declarations)
	if err != nil {
		return types.ListRequest{}, nil, err
	}

	/* // enable for debugginh
	if filter.CheckedExpr != nil {
		b, _ := json.Marshal(filter.CheckedExpr.Expr)
		fmt.Printf("filtering: %v \n", string(b))
	}
	*/

	orderBy, err := ordering.ParseOrderBy(request)
	if err != nil {
		return types.ListRequest{}, nil, err
	}

	/* // enable for debugginh
	if orderBy.Fields != nil {
		b, _ := json.Marshal(orderBy.Fields)
		fmt.Printf("ordering: %v \n", string(b))
	}
	*/

	return types.ListRequest{
		Offset:                pageToken.Offset,
		PageSize:              request.GetPageSize(),
		Filter:                filter,
		FilteringDeclarations: declarations,
		OrderBy:               orderBy,
	}, &pageToken, nil
}

func (s *server) ListLocalAgents(ctx context.Context, request *pb.ListLocalAgentsRequest) (*pb.ListLocalAgentsResponse, error) {

	req, pageToken, err := BuildListLocalAgentsRequest(ctx, request)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "list_local_agent_request",
			Message:          "error building ListLocalAgents request",
			EntityIdentifier: "local_agent",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             request,
			Error:            err,
		})
		return nil, err
	}

	// Query the storage.
	result, err := s.core.LocalAgent().List(ctx, req)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "list_local_agent",
			Message:          "error calling core in ListLocalAgents",
			EntityIdentifier: "local_agent",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             request,
			Error:            err,
		})
		return nil, err
	}

	// Build the response.
	response := &pb.ListLocalAgentsResponse{
		LocalAgents: pbmapper.LocalAgentSliceToProto(result.LocalAgents),
	}
	// Set the next page token.
	if result.HasNextPage {
		response.NextPageToken = pageToken.Next(request).String()
	}

	s.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "list_local_agent",
		Message:          "successfully handled ListLocalAgents",
		EntityIdentifier: "local_agent",
		Layer:            monitoring.ProtocolServiceLayer,
		LayerSubtype:     "protobuf",
		Type:             monitoring.EmitTypeSuccess,
		Data:             request,
	})
	// Respond.
	return response, nil
}

func validatePageSizeForListLocalAgents(request *pb.ListLocalAgentsRequest) error {
	// Handle request constraints.
	const (
		maxPageSize     = 100
		defaultPageSize = 10
	)
	switch {
	case request.PageSize < 0:
		return status.Errorf(codes.InvalidArgument, "page size is negative")
	case request.PageSize == 0:
		request.PageSize = defaultPageSize
	case request.PageSize > maxPageSize:
		request.PageSize = maxPageSize
	}
	return nil
}

func local_agentDeclarations() *filtering.Declarations {
	declarations, err := filtering.NewDeclarations(
		filtering.DeclareStandardFunctions(),
		// boolean values
		filtering.DeclareIdent("true", filtering.TypeBool),
		filtering.DeclareIdent("false", filtering.TypeBool),

		//

		//local_agent

		filtering.DeclareIdent("uuid", filtering.TypeString),

		filtering.DeclareIdent("user_uuid", filtering.TypeString),

		filtering.DeclareIdent("token_hash", filtering.TypeString),

		filtering.DeclareIdent("machine_name", filtering.TypeString),

		filtering.DeclareIdent("os", filtering.TypeString),

		filtering.DeclareIdent("cli_version", filtering.TypeString),

		filtering.DeclareEnumIdent("status", pb.LocalAgentStatus(0).Type()),

		filtering.DeclareIdent("last_seen_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("revoked_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("created_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("updated_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("created_by_uuid", filtering.TypeString),

		filtering.DeclareIdent("updated_by_uuid", filtering.TypeString),
	)
	if err != nil {
		fmt.Printf("error creating declarions:%v", err)
	}
	return declarations
}
