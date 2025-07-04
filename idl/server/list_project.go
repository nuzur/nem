package server

import (
	"context"
	"fmt"
	//"encoding/json"

	"github.com/nuzur/nem/core/module/project/types"
	pb "github.com/nuzur/nem/idl/gen"
	pbmapper "github.com/nuzur/nem/idl/mapper"
	"github.com/nuzur/nem/monitoring"

	"go.einride.tech/aip/filtering"
	"go.einride.tech/aip/ordering"
	"go.einride.tech/aip/pagination"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func BuildListProjectsRequest(ctx context.Context, request *pb.ListProjectsRequest) (types.ListRequest, *pagination.PageToken, error) {
	err := validatePageSizeForListProjects(request)
	if err != nil {
		return types.ListRequest{}, nil, err
	}

	// Use pagination.PageToken for offset-based page tokens.
	pageToken, err := pagination.ParsePageToken(request)
	if err != nil {
		return types.ListRequest{}, nil, status.Errorf(codes.InvalidArgument, "invalid page token")
	}

	// parse filters
	declarations := projectDeclarations()
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

func (s *server) ListProjects(ctx context.Context, request *pb.ListProjectsRequest) (*pb.ListProjectsResponse, error) {

	req, pageToken, err := BuildListProjectsRequest(ctx, request)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "list_project_request",
			Message:          "error building ListProjects request",
			EntityIdentifier: "project",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             request,
			Error:            err,
		})
		return nil, err
	}

	// Query the storage.
	result, err := s.core.Project().List(ctx, req)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "list_project",
			Message:          "error calling core in ListProjects",
			EntityIdentifier: "project",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             request,
			Error:            err,
		})
		return nil, err
	}

	// Build the response.
	response := &pb.ListProjectsResponse{
		Projects: pbmapper.ProjectSliceToProto(result.Projects),
	}
	// Set the next page token.
	if result.HasNextPage {
		response.NextPageToken = pageToken.Next(request).String()
	}

	s.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "list_project",
		Message:          "successfully handled ListProjects",
		EntityIdentifier: "project",
		Layer:            monitoring.ProtocolServiceLayer,
		LayerSubtype:     "protobuf",
		Type:             monitoring.EmitTypeSuccess,
		Data:             request,
	})
	// Respond.
	return response, nil
}

func validatePageSizeForListProjects(request *pb.ListProjectsRequest) error {
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

func projectDeclarations() *filtering.Declarations {
	declarations, err := filtering.NewDeclarations(
		filtering.DeclareStandardFunctions(),
		// boolean values
		filtering.DeclareIdent("true", filtering.TypeBool),
		filtering.DeclareIdent("false", filtering.TypeBool),

		//project_extension

		filtering.DeclareIdent("project_extension.uuid", filtering.TypeString),

		filtering.DeclareIdent("project_extension.extension_uuid", filtering.TypeString),

		filtering.DeclareIdent("project_extension.extension_version_uuid", filtering.TypeString),

		filtering.DeclareIdent("project_extension.configuration_entity_values", filtering.TypeString),

		filtering.DeclareIdent("project_extension.blocking", filtering.TypeBool),

		filtering.DeclareEnumIdent("project_extension.status", pb.ProjectExtensionStatus(0).Type()),

		filtering.DeclareIdent("project_extension.created_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("project_extension.updated_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("project_extension.created_by_uuid", filtering.TypeString),

		filtering.DeclareIdent("project_extension.updated_by_uuid", filtering.TypeString),

		//project

		filtering.DeclareIdent("uuid", filtering.TypeString),

		filtering.DeclareIdent("version", filtering.TypeInt),

		filtering.DeclareIdent("name", filtering.TypeString),

		filtering.DeclareIdent("description", filtering.TypeString),

		filtering.DeclareIdent("tags", filtering.TypeString),

		filtering.DeclareIdent("url", filtering.TypeString),

		filtering.DeclareIdent("owner_uuid", filtering.TypeString),

		filtering.DeclareIdent("team_uuid", filtering.TypeString),

		filtering.DeclareEnumIdent("access_type", pb.ProjectAccessType(0).Type()),

		filtering.DeclareEnumIdent("status", pb.ProjectStatus(0).Type()),

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
