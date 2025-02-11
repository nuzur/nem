package server

import (
	"context"
	"fmt"
	//"encoding/json"

	"nem/core/module/extension/types"
	pb "nem/idl/gen"
	pbmapper "nem/idl/mapper"
	"nem/monitoring"

	"go.einride.tech/aip/filtering"
	"go.einride.tech/aip/ordering"
	"go.einride.tech/aip/pagination"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func BuildListExtensionsRequest(ctx context.Context, request *pb.ListExtensionsRequest) (types.ListRequest, *pagination.PageToken, error) {
	err := validatePageSizeForListExtensions(request)
	if err != nil {
		return types.ListRequest{}, nil, err
	}

	// Use pagination.PageToken for offset-based page tokens.
	pageToken, err := pagination.ParsePageToken(request)
	if err != nil {
		return types.ListRequest{}, nil, status.Errorf(codes.InvalidArgument, "invalid page token")
	}

	// parse filters
	declarations := extensionDeclarations()
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

func (s *server) ListExtensions(ctx context.Context, request *pb.ListExtensionsRequest) (*pb.ListExtensionsResponse, error) {

	req, pageToken, err := BuildListExtensionsRequest(ctx, request)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "list_extension_request",
			Message:          "error building ListExtensions request",
			EntityIdentifier: "extension",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             request,
			Error:            err,
		})
		return nil, err
	}

	// Query the storage.
	result, err := s.core.Extension().List(ctx, req)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "list_extension",
			Message:          "error calling core in ListExtensions",
			EntityIdentifier: "extension",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             request,
			Error:            err,
		})
		return nil, err
	}

	// Build the response.
	response := &pb.ListExtensionsResponse{
		Extensions: pbmapper.ExtensionSliceToProto(result.Extensions),
	}
	// Set the next page token.
	if result.HasNextPage {
		response.NextPageToken = pageToken.Next(request).String()
	}

	s.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "list_extension",
		Message:          "successfully handled ListExtensions",
		EntityIdentifier: "extension",
		Layer:            monitoring.ProtocolServiceLayer,
		LayerSubtype:     "protobuf",
		Type:             monitoring.EmitTypeSuccess,
		Data:             request,
	})
	// Respond.
	return response, nil
}

func validatePageSizeForListExtensions(request *pb.ListExtensionsRequest) error {
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

func extensionDeclarations() *filtering.Declarations {
	declarations, err := filtering.NewDeclarations(
		filtering.DeclareStandardFunctions(),
		// boolean values
		filtering.DeclareIdent("true", filtering.TypeBool),
		filtering.DeclareIdent("false", filtering.TypeBool),

		//visibility

		filtering.DeclareIdent("visibility.uuid", filtering.TypeString),

		filtering.DeclareIdent("visibility.description", filtering.TypeString),

		filtering.DeclareIdent("visibility.organization_uuids", filtering.TypeString),

		filtering.DeclareIdent("visibility.team_uuids", filtering.TypeString),

		filtering.DeclareIdent("visibility.user_uuids", filtering.TypeString),

		filtering.DeclareEnumIdent("visibility.roles", pb.VisibilityRole(0).Type()),

		filtering.DeclareIdent("visibility.created_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("visibility.updated_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("visibility.created_by_uuid", filtering.TypeString),

		filtering.DeclareIdent("visibility.updated_by_uuid", filtering.TypeString),

		//extension

		filtering.DeclareIdent("uuid", filtering.TypeString),

		filtering.DeclareIdent("version", filtering.TypeInt),

		filtering.DeclareIdent("identifier", filtering.TypeString),

		filtering.DeclareIdent("display_name", filtering.TypeString),

		filtering.DeclareIdent("display_author_name", filtering.TypeString),

		filtering.DeclareIdent("description", filtering.TypeString),

		filtering.DeclareIdent("url", filtering.TypeString),

		filtering.DeclareIdent("verfied", filtering.TypeBool),

		filtering.DeclareIdent("repository", filtering.TypeString),

		filtering.DeclareEnumIdent("extension_type", pb.ExtensionExtensionType(0).Type()),

		filtering.DeclareIdent("tags", filtering.TypeString),

		filtering.DeclareIdent("public", filtering.TypeBool),

		filtering.DeclareEnumIdent("status", pb.ExtensionStatus(0).Type()),

		filtering.DeclareIdent("owner_uuid", filtering.TypeString),

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
