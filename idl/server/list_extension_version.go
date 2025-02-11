package server

import (
	"context"
	"fmt"
	//"encoding/json"

	"nem/core/module/extension_version/types"
	pb "nem/idl/gen"
	pbmapper "nem/idl/mapper"
	"nem/monitoring"

	"go.einride.tech/aip/filtering"
	"go.einride.tech/aip/ordering"
	"go.einride.tech/aip/pagination"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func BuildListExtensionVersionsRequest(ctx context.Context, request *pb.ListExtensionVersionsRequest) (types.ListRequest, *pagination.PageToken, error) {
	err := validatePageSizeForListExtensionVersions(request)
	if err != nil {
		return types.ListRequest{}, nil, err
	}

	// Use pagination.PageToken for offset-based page tokens.
	pageToken, err := pagination.ParsePageToken(request)
	if err != nil {
		return types.ListRequest{}, nil, status.Errorf(codes.InvalidArgument, "invalid page token")
	}

	// parse filters
	declarations := extension_versionDeclarations()
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

func (s *server) ListExtensionVersions(ctx context.Context, request *pb.ListExtensionVersionsRequest) (*pb.ListExtensionVersionsResponse, error) {

	req, pageToken, err := BuildListExtensionVersionsRequest(ctx, request)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "list_extension_version_request",
			Message:          "error building ListExtensionVersions request",
			EntityIdentifier: "extension_version",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             request,
			Error:            err,
		})
		return nil, err
	}

	// Query the storage.
	result, err := s.core.ExtensionVersion().List(ctx, req)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "list_extension_version",
			Message:          "error calling core in ListExtensionVersions",
			EntityIdentifier: "extension_version",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             request,
			Error:            err,
		})
		return nil, err
	}

	// Build the response.
	response := &pb.ListExtensionVersionsResponse{
		ExtensionVersions: pbmapper.ExtensionVersionSliceToProto(result.ExtensionVersions),
	}
	// Set the next page token.
	if result.HasNextPage {
		response.NextPageToken = pageToken.Next(request).String()
	}

	s.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "list_extension_version",
		Message:          "successfully handled ListExtensionVersions",
		EntityIdentifier: "extension_version",
		Layer:            monitoring.ProtocolServiceLayer,
		LayerSubtype:     "protobuf",
		Type:             monitoring.EmitTypeSuccess,
		Data:             request,
	})
	// Respond.
	return response, nil
}

func validatePageSizeForListExtensionVersions(request *pb.ListExtensionVersionsRequest) error {
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

func extension_versionDeclarations() *filtering.Declarations {
	declarations, err := filtering.NewDeclarations(
		filtering.DeclareStandardFunctions(),
		// boolean values
		filtering.DeclareIdent("true", filtering.TypeBool),
		filtering.DeclareIdent("false", filtering.TypeBool),

		//extension_version

		filtering.DeclareIdent("uuid", filtering.TypeString),

		filtering.DeclareIdent("version", filtering.TypeInt),

		filtering.DeclareIdent("extension_uuid", filtering.TypeString),

		filtering.DeclareIdent("display_version", filtering.TypeString),

		filtering.DeclareIdent("description", filtering.TypeString),

		filtering.DeclareIdent("repository_tag", filtering.TypeString),

		filtering.DeclareIdent("configuration_entity", filtering.TypeString),

		filtering.DeclareEnumIdent("execution_mode", pb.ExtensionVersionExecutionMode(0).Type()),

		filtering.DeclareEnumIdent("review_status", pb.ExtensionVersionReviewStatus(0).Type()),

		filtering.DeclareEnumIdent("status", pb.ExtensionVersionStatus(0).Type()),

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
