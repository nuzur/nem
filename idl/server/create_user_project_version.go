package server

import (
	"context"
	"errors"
	"github.com/nuzur/nem/core/module/user_project_version/types"
	pb "github.com/nuzur/nem/idl/gen"
	pbmapper "github.com/nuzur/nem/idl/mapper"
	"github.com/nuzur/nem/monitoring"
)

func (s *server) CreateUserProjectVersion(ctx context.Context, req *pb.CreateUserProjectVersionRequest) (*pb.UserProjectVersion, error) {
	res, err := s.core.UserProjectVersion().Insert(ctx, types.UpsertRequest{
		UserProjectVersion: pbmapper.UserProjectVersionFromProto(req.GetUserProjectVersion()),
	})
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "create_user_project_version",
			Message:          "error calling core upsert in CreateUserProjectVersion",
			EntityIdentifier: "user_project_version",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return nil, err
	}

	fetchRes, err := s.core.UserProjectVersion().FetchUserProjectVersionByUUID(ctx, types.FetchUserProjectVersionByUUIDRequest(res))
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "create_user_project_version",
			Message:          "error fetching after upsert in CreateUserProjectVersion",
			EntityIdentifier: "user_project_version",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return nil, err
	}

	if len(fetchRes.Results) == 0 {
		err := errors.New("error fetching entity")
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "create_user_project_version",
			Message:          "entity not found after upsert in CreateUserProjectVersion",
			EntityIdentifier: "user_project_version",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return nil, err
	}

	s.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "create_user_project_version",
		Message:          "successfully handled CreateUserProjectVersion",
		EntityIdentifier: "user_project_version",
		Layer:            monitoring.ProtocolServiceLayer,
		LayerSubtype:     "protobuf",
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return pbmapper.UserProjectVersionToProto(fetchRes.Results[0]), nil
}
