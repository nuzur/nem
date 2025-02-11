package server

import (
	"context"
	"errors"
	"github.com/nuzur/nem/core/module/team/types"
	pb "github.com/nuzur/nem/idl/gen"
	pbmapper "github.com/nuzur/nem/idl/mapper"
	"github.com/nuzur/nem/monitoring"
)

func (s *server) CreateTeam(ctx context.Context, req *pb.CreateTeamRequest) (*pb.Team, error) {
	res, err := s.core.Team().Insert(ctx, types.UpsertRequest{
		Team: pbmapper.TeamFromProto(req.GetTeam()),
	})
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "create_team",
			Message:          "error calling core upsert in CreateTeam",
			EntityIdentifier: "team",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return nil, err
	}

	fetchRes, err := s.core.Team().FetchTeamByUUID(ctx, types.FetchTeamByUUIDRequest(res))
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "create_team",
			Message:          "error fetching after upsert in CreateTeam",
			EntityIdentifier: "team",
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
			ActionIdentifier: "create_team",
			Message:          "entity not found after upsert in CreateTeam",
			EntityIdentifier: "team",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return nil, err
	}

	s.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "create_team",
		Message:          "successfully handled CreateTeam",
		EntityIdentifier: "team",
		Layer:            monitoring.ProtocolServiceLayer,
		LayerSubtype:     "protobuf",
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return pbmapper.TeamToProto(fetchRes.Results[0]), nil
}
