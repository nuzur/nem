package server

import (
	"context"
	"fmt"
	//"encoding/json"

	"github.com/nuzur/nem/core/module/team/types"
	pb "github.com/nuzur/nem/idl/gen"
	pbmapper "github.com/nuzur/nem/idl/mapper"
	"github.com/nuzur/nem/monitoring"

	"go.einride.tech/aip/filtering"
	"go.einride.tech/aip/ordering"
	"go.einride.tech/aip/pagination"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func BuildListTeamsRequest(ctx context.Context, request *pb.ListTeamsRequest) (types.ListRequest, *pagination.PageToken, error) {
	err := validatePageSizeForListTeams(request)
	if err != nil {
		return types.ListRequest{}, nil, err
	}

	// Use pagination.PageToken for offset-based page tokens.
	pageToken, err := pagination.ParsePageToken(request)
	if err != nil {
		return types.ListRequest{}, nil, status.Errorf(codes.InvalidArgument, "invalid page token")
	}

	// parse filters
	declarations := teamDeclarations()
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

func (s *server) ListTeams(ctx context.Context, request *pb.ListTeamsRequest) (*pb.ListTeamsResponse, error) {

	req, pageToken, err := BuildListTeamsRequest(ctx, request)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "list_team_request",
			Message:          "error building ListTeams request",
			EntityIdentifier: "team",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             request,
			Error:            err,
		})
		return nil, err
	}

	// Query the storage.
	result, err := s.core.Team().List(ctx, req)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "list_team",
			Message:          "error calling core in ListTeams",
			EntityIdentifier: "team",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             request,
			Error:            err,
		})
		return nil, err
	}

	// Build the response.
	response := &pb.ListTeamsResponse{
		Teams: pbmapper.TeamSliceToProto(result.Teams),
	}
	// Set the next page token.
	if result.HasNextPage {
		response.NextPageToken = pageToken.Next(request).String()
	}

	s.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "list_team",
		Message:          "successfully handled ListTeams",
		EntityIdentifier: "team",
		Layer:            monitoring.ProtocolServiceLayer,
		LayerSubtype:     "protobuf",
		Type:             monitoring.EmitTypeSuccess,
		Data:             request,
	})
	// Respond.
	return response, nil
}

func validatePageSizeForListTeams(request *pb.ListTeamsRequest) error {
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

func teamDeclarations() *filtering.Declarations {
	declarations, err := filtering.NewDeclarations(
		filtering.DeclareStandardFunctions(),
		// boolean values
		filtering.DeclareIdent("true", filtering.TypeBool),
		filtering.DeclareIdent("false", filtering.TypeBool),

		//enviorment

		filtering.DeclareIdent("enviorment.uuid", filtering.TypeString),

		filtering.DeclareIdent("enviorment.identifier", filtering.TypeString),

		filtering.DeclareIdent("enviorment.critical", filtering.TypeBool),

		filtering.DeclareEnumIdent("enviorment.status", pb.EnviormentStatus(0).Type()),

		filtering.DeclareIdent("enviorment.created_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("enviorment.updated_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("enviorment.created_by_uuid", filtering.TypeString),

		filtering.DeclareIdent("enviorment.updated_by_uuid", filtering.TypeString),

		//review_config

		filtering.DeclareIdent("review_config.uuid", filtering.TypeString),

		filtering.DeclareEnumIdent("review_config.types", pb.ReviewConfigType(0).Type()),

		filtering.DeclareEnumIdent("review_config.user_roles", pb.ReviewConfigUserRole(0).Type()),

		filtering.DeclareIdent("review_config.min_reviews", filtering.TypeInt),

		filtering.DeclareEnumIdent("review_config.status", pb.ReviewConfigStatus(0).Type()),

		filtering.DeclareIdent("review_config.created_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("review_config.updated_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("review_config.created_by_uuid", filtering.TypeString),

		filtering.DeclareIdent("review_config.updated_by_uuid", filtering.TypeString),

		//membership

		filtering.DeclareIdent("membership.uuid", filtering.TypeString),

		filtering.DeclareIdent("membership.owner_uuid", filtering.TypeString),

		filtering.DeclareEnumIdent("membership.type", pb.MembershipType(0).Type()),

		filtering.DeclareIdent("membership.metadata", filtering.TypeString),

		filtering.DeclareEnumIdent("membership.status", pb.MembershipStatus(0).Type()),

		filtering.DeclareIdent("membership.created_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("membership.updated_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("membership.created_by_uuid", filtering.TypeString),

		filtering.DeclareIdent("membership.updated_by_uuid", filtering.TypeString),

		//store

		filtering.DeclareIdent("store.uuid", filtering.TypeString),

		filtering.DeclareIdent("store.identifier", filtering.TypeString),

		filtering.DeclareIdent("store.description", filtering.TypeString),

		filtering.DeclareEnumIdent("store.status", pb.StoreStatus(0).Type()),

		filtering.DeclareIdent("store.created_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("store.updated_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("store.created_by_uuid", filtering.TypeString),

		filtering.DeclareIdent("store.updated_by_uuid", filtering.TypeString),

		//connection

		filtering.DeclareIdent("connection.uuid", filtering.TypeString),

		filtering.DeclareIdent("connection.version", filtering.TypeInt),

		filtering.DeclareIdent("connection.store_uuid", filtering.TypeString),

		filtering.DeclareIdent("connection.enviorment_uuid", filtering.TypeString),

		filtering.DeclareIdent("connection.identifier", filtering.TypeString),

		filtering.DeclareEnumIdent("connection.db_type", pb.ConnectionDbType(0).Type()),

		filtering.DeclareIdent("connection.db_version", filtering.TypeString),

		filtering.DeclareEnumIdent("connection.type", pb.ConnectionType(0).Type()),

		filtering.DeclareEnumIdent("connection.status", pb.ConnectionStatus(0).Type()),

		filtering.DeclareIdent("connection.created_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("connection.updated_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("connection.created_by_uuid", filtering.TypeString),

		filtering.DeclareIdent("connection.updated_by_uuid", filtering.TypeString),

		//object_store

		filtering.DeclareIdent("object_store.uuid", filtering.TypeString),

		filtering.DeclareIdent("object_store.identifier", filtering.TypeString),

		filtering.DeclareEnumIdent("object_store.type", pb.ObjectStoreType(0).Type()),

		filtering.DeclareEnumIdent("object_store.status", pb.ObjectStoreStatus(0).Type()),

		filtering.DeclareIdent("object_store.created_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("object_store.updated_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("object_store.created_by_uuid", filtering.TypeString),

		filtering.DeclareIdent("object_store.updated_by_uuid", filtering.TypeString),

		//

		//team

		filtering.DeclareIdent("uuid", filtering.TypeString),

		filtering.DeclareIdent("version", filtering.TypeInt),

		filtering.DeclareIdent("name", filtering.TypeString),

		filtering.DeclareIdent("organization_uuid", filtering.TypeString),

		filtering.DeclareEnumIdent("status", pb.TeamStatus(0).Type()),

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
