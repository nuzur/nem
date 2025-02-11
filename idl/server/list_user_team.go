package server

import (
	"context"
	"fmt"
	//"encoding/json"

	"github.com/nuzur/nem/core/module/user_team/types"
	pb "github.com/nuzur/nem/idl/gen"
	pbmapper "github.com/nuzur/nem/idl/mapper"
	"github.com/nuzur/nem/monitoring"

	"go.einride.tech/aip/filtering"
	"go.einride.tech/aip/ordering"
	"go.einride.tech/aip/pagination"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func BuildListUserTeamsRequest(ctx context.Context, request *pb.ListUserTeamsRequest) (types.ListRequest, *pagination.PageToken, error) {
	err := validatePageSizeForListUserTeams(request)
	if err != nil {
		return types.ListRequest{}, nil, err
	}

	// Use pagination.PageToken for offset-based page tokens.
	pageToken, err := pagination.ParsePageToken(request)
	if err != nil {
		return types.ListRequest{}, nil, status.Errorf(codes.InvalidArgument, "invalid page token")
	}

	// parse filters
	declarations := user_teamDeclarations()
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

func (s *server) ListUserTeams(ctx context.Context, request *pb.ListUserTeamsRequest) (*pb.ListUserTeamsResponse, error) {

	req, pageToken, err := BuildListUserTeamsRequest(ctx, request)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "list_user_team_request",
			Message:          "error building ListUserTeams request",
			EntityIdentifier: "user_team",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             request,
			Error:            err,
		})
		return nil, err
	}

	// Query the storage.
	result, err := s.core.UserTeam().List(ctx, req)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "list_user_team",
			Message:          "error calling core in ListUserTeams",
			EntityIdentifier: "user_team",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             request,
			Error:            err,
		})
		return nil, err
	}

	// Build the response.
	response := &pb.ListUserTeamsResponse{
		UserTeams: pbmapper.UserTeamSliceToProto(result.UserTeams),
	}
	// Set the next page token.
	if result.HasNextPage {
		response.NextPageToken = pageToken.Next(request).String()
	}

	s.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "list_user_team",
		Message:          "successfully handled ListUserTeams",
		EntityIdentifier: "user_team",
		Layer:            monitoring.ProtocolServiceLayer,
		LayerSubtype:     "protobuf",
		Type:             monitoring.EmitTypeSuccess,
		Data:             request,
	})
	// Respond.
	return response, nil
}

func validatePageSizeForListUserTeams(request *pb.ListUserTeamsRequest) error {
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

func user_teamDeclarations() *filtering.Declarations {
	declarations, err := filtering.NewDeclarations(
		filtering.DeclareStandardFunctions(),
		// boolean values
		filtering.DeclareIdent("true", filtering.TypeBool),
		filtering.DeclareIdent("false", filtering.TypeBool),

		//user_team

		filtering.DeclareIdent("uuid", filtering.TypeString),

		filtering.DeclareIdent("user_uuid", filtering.TypeString),

		filtering.DeclareIdent("user_email", filtering.TypeString),

		filtering.DeclareIdent("team_uuid", filtering.TypeString),

		filtering.DeclareEnumIdent("roles", pb.UserTeamRole(0).Type()),

		filtering.DeclareEnumIdent("status", pb.UserTeamStatus(0).Type()),

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
