package server

import (
	"context"
	"errors"
	"github.com/nuzur/nem/core/module/team/types"
	pb "github.com/nuzur/nem/idl/gen"
	pbmapper "github.com/nuzur/nem/idl/mapper"
	"github.com/nuzur/nem/monitoring"
	"go.einride.tech/aip/fieldmask"
	"strings"
)

func (s *server) UpdateTeam(ctx context.Context, req *pb.UpdateTeamRequest) (*pb.Team, error) {

	if req.Team.Uuid == "" {
		return nil, errors.New("please provide a valid UUID to update")
	}

	if req.Team.Version == 0 {
		return nil, errors.New("please provide a valid version to update")
	}

	err := fieldmask.Validate(req.UpdateMask, req)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "update_team",
			Message:          "error validating fieldmask in UpdateTeam",
			EntityIdentifier: "team",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return nil, err
	}

	isFull := fieldmask.IsFullReplacement(req.UpdateMask)

	if !isFull && req.UpdateMask != nil {
		if !strings.Contains(req.UpdateMask.String(), "team.uuid") {
			req.UpdateMask.Append(req, "team.uuid")
		}

		if !strings.Contains(req.UpdateMask.String(), "team.version") {
			req.UpdateMask.Append(req, "team.version")
		}

		newReq := &pb.UpdateTeamRequest{}
		fieldmask.Update(req.UpdateMask, newReq, req)
		req = newReq
	}

	res, err := s.core.Team().Update(ctx, types.UpsertRequest{
		Team: pbmapper.TeamFromProto(req.GetTeam()),
	}, !isFull)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "update_team",
			Message:          "error calling core upsert in UpdateTeam",
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
			ActionIdentifier: "update_team",
			Message:          "error fetching after upsert in UpdateTeam",
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
			ActionIdentifier: "update_team",
			Message:          "entity not found after upsert in UpdateTeam",
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
		ActionIdentifier: "update_team",
		Message:          "successfully handled UpdateTeam",
		EntityIdentifier: "team",
		Layer:            monitoring.ProtocolServiceLayer,
		LayerSubtype:     "protobuf",
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return pbmapper.TeamToProto(fetchRes.Results[0]), nil
}
