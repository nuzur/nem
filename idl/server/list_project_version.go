package server

import (
	"context"
	"fmt"
	//"encoding/json"

	"github.com/nuzur/nem/core/module/project_version/types"
	pb "github.com/nuzur/nem/idl/gen"
	pbmapper "github.com/nuzur/nem/idl/mapper"
	"github.com/nuzur/nem/monitoring"

	"go.einride.tech/aip/filtering"
	"go.einride.tech/aip/ordering"
	"go.einride.tech/aip/pagination"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func BuildListProjectVersionsRequest(ctx context.Context, request *pb.ListProjectVersionsRequest) (types.ListRequest, *pagination.PageToken, error) {
	err := validatePageSizeForListProjectVersions(request)
	if err != nil {
		return types.ListRequest{}, nil, err
	}

	// Use pagination.PageToken for offset-based page tokens.
	pageToken, err := pagination.ParsePageToken(request)
	if err != nil {
		return types.ListRequest{}, nil, status.Errorf(codes.InvalidArgument, "invalid page token")
	}

	// parse filters
	declarations := project_versionDeclarations()
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

func (s *server) ListProjectVersions(ctx context.Context, request *pb.ListProjectVersionsRequest) (*pb.ListProjectVersionsResponse, error) {

	req, pageToken, err := BuildListProjectVersionsRequest(ctx, request)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "list_project_version_request",
			Message:          "error building ListProjectVersions request",
			EntityIdentifier: "project_version",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             request,
			Error:            err,
		})
		return nil, err
	}

	// Query the storage.
	result, err := s.core.ProjectVersion().List(ctx, req)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "list_project_version",
			Message:          "error calling core in ListProjectVersions",
			EntityIdentifier: "project_version",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             request,
			Error:            err,
		})
		return nil, err
	}

	// Build the response.
	response := &pb.ListProjectVersionsResponse{
		ProjectVersions: pbmapper.ProjectVersionSliceToProto(result.ProjectVersions),
	}
	// Set the next page token.
	if result.HasNextPage {
		response.NextPageToken = pageToken.Next(request).String()
	}

	s.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "list_project_version",
		Message:          "successfully handled ListProjectVersions",
		EntityIdentifier: "project_version",
		Layer:            monitoring.ProtocolServiceLayer,
		LayerSubtype:     "protobuf",
		Type:             monitoring.EmitTypeSuccess,
		Data:             request,
	})
	// Respond.
	return response, nil
}

func validatePageSizeForListProjectVersions(request *pb.ListProjectVersionsRequest) error {
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

func project_versionDeclarations() *filtering.Declarations {
	declarations, err := filtering.NewDeclarations(
		filtering.DeclareStandardFunctions(),
		// boolean values
		filtering.DeclareIdent("true", filtering.TypeBool),
		filtering.DeclareIdent("false", filtering.TypeBool),

		//entity

		filtering.DeclareIdent("entity.uuid", filtering.TypeString),

		filtering.DeclareIdent("entity.version", filtering.TypeInt),

		filtering.DeclareIdent("entity.identifier", filtering.TypeString),

		filtering.DeclareIdent("entity.description", filtering.TypeString),

		filtering.DeclareEnumIdent("entity.type", pb.EntityType(0).Type()),

		filtering.DeclareEnumIdent("entity.status", pb.EntityStatus(0).Type()),

		filtering.DeclareIdent("entity.created_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("entity.updated_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("entity.created_by_uuid", filtering.TypeString),

		filtering.DeclareIdent("entity.updated_by_uuid", filtering.TypeString),

		//relationship

		filtering.DeclareIdent("relationship.uuid", filtering.TypeString),

		filtering.DeclareIdent("relationship.version", filtering.TypeInt),

		filtering.DeclareIdent("relationship.identifier", filtering.TypeString),

		filtering.DeclareIdent("relationship.description", filtering.TypeString),

		filtering.DeclareEnumIdent("relationship.cardinality", pb.RelationshipCardinality(0).Type()),

		filtering.DeclareIdent("relationship.use_foreign_key", filtering.TypeBool),

		filtering.DeclareEnumIdent("relationship.status", pb.RelationshipStatus(0).Type()),

		filtering.DeclareIdent("relationship.created_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("relationship.updated_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("relationship.created_by_uuid", filtering.TypeString),

		filtering.DeclareIdent("relationship.updated_by_uuid", filtering.TypeString),

		//enum

		filtering.DeclareIdent("enum.uuid", filtering.TypeString),

		filtering.DeclareIdent("enum.version", filtering.TypeInt),

		filtering.DeclareIdent("enum.identifier", filtering.TypeString),

		filtering.DeclareIdent("enum.remote_values", filtering.TypeBool),

		filtering.DeclareEnumIdent("enum.status", pb.EnumStatus(0).Type()),

		filtering.DeclareIdent("enum.created_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("enum.updated_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("enum.created_by_uuid", filtering.TypeString),

		filtering.DeclareIdent("enum.updated_by_uuid", filtering.TypeString),

		//service

		filtering.DeclareIdent("service.uuid", filtering.TypeString),

		filtering.DeclareIdent("service.version", filtering.TypeInt),

		filtering.DeclareIdent("service.identifier", filtering.TypeString),

		filtering.DeclareIdent("service.description", filtering.TypeString),

		filtering.DeclareEnumIdent("service.status", pb.ServiceStatus(0).Type()),

		filtering.DeclareIdent("service.created_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("service.updated_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("service.created_by_uuid", filtering.TypeString),

		filtering.DeclareIdent("service.updated_by_uuid", filtering.TypeString),

		//

		//

		//project_version

		filtering.DeclareIdent("uuid", filtering.TypeString),

		filtering.DeclareIdent("version", filtering.TypeInt),

		filtering.DeclareIdent("identifier", filtering.TypeString),

		filtering.DeclareIdent("description", filtering.TypeString),

		filtering.DeclareIdent("project_uuid", filtering.TypeString),

		filtering.DeclareIdent("base_version_uuid", filtering.TypeString),

		filtering.DeclareEnumIdent("review_status", pb.ProjectVersionReviewStatus(0).Type()),

		filtering.DeclareEnumIdent("status", pb.ProjectVersionStatus(0).Type()),

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
