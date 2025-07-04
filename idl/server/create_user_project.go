package server

import (
	"context"
	"errors"
	"github.com/nuzur/nem/core/module/user_project/types"
	pb "github.com/nuzur/nem/idl/gen"
	pbmapper "github.com/nuzur/nem/idl/mapper"
	"github.com/nuzur/nem/monitoring"
)

func (s *server) CreateUserProject(ctx context.Context, req *pb.CreateUserProjectRequest) (*pb.UserProject, error) {
	res, err := s.core.UserProject().Insert(ctx, types.UpsertRequest{
		UserProject: pbmapper.UserProjectFromProto(req.GetUserProject()),
	})
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "create_user_project",
			Message:          "error calling core upsert in CreateUserProject",
			EntityIdentifier: "user_project",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return nil, err
	}

	fetchRes, err := s.core.UserProject().FetchUserProjectByUUID(ctx, types.FetchUserProjectByUUIDRequest(res))
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "create_user_project",
			Message:          "error fetching after upsert in CreateUserProject",
			EntityIdentifier: "user_project",
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
			ActionIdentifier: "create_user_project",
			Message:          "entity not found after upsert in CreateUserProject",
			EntityIdentifier: "user_project",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return nil, err
	}

	s.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "create_user_project",
		Message:          "successfully handled CreateUserProject",
		EntityIdentifier: "user_project",
		Layer:            monitoring.ProtocolServiceLayer,
		LayerSubtype:     "protobuf",
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return pbmapper.UserProjectToProto(fetchRes.Results[0]), nil
}
