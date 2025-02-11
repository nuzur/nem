package server

import (
	"context"
	"fmt"
	//"encoding/json"

	"github.com/nuzur/nem/core/module/user/types"
	pb "github.com/nuzur/nem/idl/gen"
	pbmapper "github.com/nuzur/nem/idl/mapper"
	"github.com/nuzur/nem/monitoring"

	"go.einride.tech/aip/filtering"
	"go.einride.tech/aip/ordering"
	"go.einride.tech/aip/pagination"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func BuildListUsersRequest(ctx context.Context, request *pb.ListUsersRequest) (types.ListRequest, *pagination.PageToken, error) {
	err := validatePageSizeForListUsers(request)
	if err != nil {
		return types.ListRequest{}, nil, err
	}

	// Use pagination.PageToken for offset-based page tokens.
	pageToken, err := pagination.ParsePageToken(request)
	if err != nil {
		return types.ListRequest{}, nil, status.Errorf(codes.InvalidArgument, "invalid page token")
	}

	// parse filters
	declarations := userDeclarations()
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

func (s *server) ListUsers(ctx context.Context, request *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {

	req, pageToken, err := BuildListUsersRequest(ctx, request)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "list_user_request",
			Message:          "error building ListUsers request",
			EntityIdentifier: "user",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             request,
			Error:            err,
		})
		return nil, err
	}

	// Query the storage.
	result, err := s.core.User().List(ctx, req)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "list_user",
			Message:          "error calling core in ListUsers",
			EntityIdentifier: "user",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             request,
			Error:            err,
		})
		return nil, err
	}

	// Build the response.
	response := &pb.ListUsersResponse{
		Users: pbmapper.UserSliceToProto(result.Users),
	}
	// Set the next page token.
	if result.HasNextPage {
		response.NextPageToken = pageToken.Next(request).String()
	}

	s.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "list_user",
		Message:          "successfully handled ListUsers",
		EntityIdentifier: "user",
		Layer:            monitoring.ProtocolServiceLayer,
		LayerSubtype:     "protobuf",
		Type:             monitoring.EmitTypeSuccess,
		Data:             request,
	})
	// Respond.
	return response, nil
}

func validatePageSizeForListUsers(request *pb.ListUsersRequest) error {
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

func userDeclarations() *filtering.Declarations {
	declarations, err := filtering.NewDeclarations(
		filtering.DeclareStandardFunctions(),
		// boolean values
		filtering.DeclareIdent("true", filtering.TypeBool),
		filtering.DeclareIdent("false", filtering.TypeBool),

		//user

		filtering.DeclareIdent("uuid", filtering.TypeString),

		filtering.DeclareIdent("identifier", filtering.TypeString),

		filtering.DeclareIdent("name", filtering.TypeString),

		filtering.DeclareIdent("last_name", filtering.TypeString),

		filtering.DeclareIdent("email", filtering.TypeString),

		filtering.DeclareEnumIdent("user_type", pb.UserUserType(0).Type()),

		filtering.DeclareIdent("country_ios2", filtering.TypeString),

		filtering.DeclareIdent("locale", filtering.TypeString),

		filtering.DeclareIdent("metadata", filtering.TypeString),

		filtering.DeclareEnumIdent("status", pb.UserStatus(0).Type()),

		filtering.DeclareIdent("created_at", filtering.TypeTimestamp),

		filtering.DeclareIdent("updated_at", filtering.TypeTimestamp),
	)
	if err != nil {
		fmt.Printf("error creating declarions:%v", err)
	}
	return declarations
}
