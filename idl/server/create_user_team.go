package server

import (
	"context"
	"errors"
	"nem/core/module/user_team/types"
	pb "nem/idl/gen"
	pbmapper "nem/idl/mapper"
	"nem/monitoring"
)

func (s *server) CreateUserTeam(ctx context.Context, req *pb.CreateUserTeamRequest) (*pb.UserTeam, error) {
	res, err := s.core.UserTeam().Insert(ctx, types.UpsertRequest{
		UserTeam: pbmapper.UserTeamFromProto(req.GetUserTeam()),
	})
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "create_user_team",
			Message:          "error calling core upsert in CreateUserTeam",
			EntityIdentifier: "user_team",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return nil, err
	}

	fetchRes, err := s.core.UserTeam().FetchUserTeamByUUID(ctx, types.FetchUserTeamByUUIDRequest(res))
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "create_user_team",
			Message:          "error fetching after upsert in CreateUserTeam",
			EntityIdentifier: "user_team",
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
			ActionIdentifier: "create_user_team",
			Message:          "entity not found after upsert in CreateUserTeam",
			EntityIdentifier: "user_team",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return nil, err
	}

	s.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "create_user_team",
		Message:          "successfully handled CreateUserTeam",
		EntityIdentifier: "user_team",
		Layer:            monitoring.ProtocolServiceLayer,
		LayerSubtype:     "protobuf",
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return pbmapper.UserTeamToProto(fetchRes.Results[0]), nil
}
