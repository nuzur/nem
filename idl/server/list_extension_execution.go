package server

import (
	"context"
	"fmt"
	//"encoding/json"

	"github.com/nuzur/nem/core/module/extension_execution/types"
	pb "github.com/nuzur/nem/idl/gen"
	pbmapper "github.com/nuzur/nem/idl/mapper"
	"github.com/nuzur/nem/monitoring"

	"go.einride.tech/aip/filtering"
	"go.einride.tech/aip/ordering"
	"go.einride.tech/aip/pagination"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func BuildListExtensionExecutionsRequest(ctx context.Context, request *pb.ListExtensionExecutionsRequest) (types.ListRequest, *pagination.PageToken, error) {
	err := validatePageSizeForListExtensionExecutions(request)
	if err != nil {
		return types.ListRequest{}, nil, err
	}

	// Use pagination.PageToken for offset-based page tokens.
	pageToken, err := pagination.ParsePageToken(request)
	if err != nil {
		return types.ListRequest{}, nil, status.Errorf(codes.InvalidArgument, "invalid page token")
	}

	// parse filters
	declarations := extension_executionDeclarations()
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

func (s *server) ListExtensionExecutions(ctx context.Context, request *pb.ListExtensionExecutionsRequest) (*pb.ListExtensionExecutionsResponse, error) {

	req, pageToken, err := BuildListExtensionExecutionsRequest(ctx, request)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "list_extension_execution_request",
			Message:          "error building ListExtensionExecutions request",
			EntityIdentifier: "extension_execution",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             request,
			Error:            err,
		})
		return nil, err
	}

	// Query the storage.
	result, err := s.core.ExtensionExecution().List(ctx, req)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "list_extension_execution",
			Message:          "error calling core in ListExtensionExecutions",
			EntityIdentifier: "extension_execution",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             request,
			Error:            err,
		})
		return nil, err
	}

	// Build the response.
	response := &pb.ListExtensionExecutionsResponse{
		ExtensionExecutions: pbmapper.ExtensionExecutionSliceToProto(result.ExtensionExecutions),
	}
	// Set the next page token.
	if result.HasNextPage {
		response.NextPageToken = pageToken.Next(request).String()
	}

	s.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "list_extension_execution",
		Message:          "successfully handled ListExtensionExecutions",
		EntityIdentifier: "extension_execution",
		Layer:            monitoring.ProtocolServiceLayer,
		LayerSubtype:     "protobuf",
		Type:             monitoring.EmitTypeSuccess,
		Data:             request,
	})
	// Respond.
	return response, nil
}

func validatePageSizeForListExtensionExecutions(request *pb.ListExtensionExecutionsRequest) error {
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

func extension_executionDeclarations() *filtering.Declarations {
	declarations, err := filtering.NewDeclarations(
		filtering.DeclareStandardFunctions(),
		// boolean values
		filtering.DeclareIdent("true", filtering.TypeBool),
		filtering.DeclareIdent("false", filtering.TypeBool),

		//extension_execution

		filtering.DeclareIdent("uuid", filtering.TypeString),

		filtering.DeclareIdent("extension_uuid", filtering.TypeString),

		filtering.DeclareIdent("extension_version_uuid", filtering.TypeString),

		filtering.DeclareIdent("project_extension_uuid", filtering.TypeString),

		filtering.DeclareIdent("project_uuid", filtering.TypeString),

		filtering.DeclareIdent("project_version_uuid", filtering.TypeString),

		filtering.DeclareIdent("executed_by_uuid", filtering.TypeString),

		filtering.DeclareIdent("metadata", filtering.TypeString),

		filtering.DeclareEnumIdent("status", pb.ExtensionExecutionStatus(0).Type()),

		filtering.DeclareIdent("status_msg", filtering.TypeString),

		filtering.DeclareIdent("created_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("updated_at", filtering.TypeTimestamp),
	)
	if err != nil {
		fmt.Printf("error creating declarions:%v", err)
	}
	return declarations
}
