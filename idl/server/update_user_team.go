package server

import (
	"context"
	"errors"
	"go.einride.tech/aip/fieldmask"
	"nem/core/module/user_team/types"
	pb "nem/idl/gen"
	pbmapper "nem/idl/mapper"
	"nem/monitoring"
	"strings"
)

func (s *server) UpdateUserTeam(ctx context.Context, req *pb.UpdateUserTeamRequest) (*pb.UserTeam, error) {

	if req.UserTeam.Uuid == "" {
		return nil, errors.New("please provide a valid UUID to update")
	}

	err := fieldmask.Validate(req.UpdateMask, req)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "update_user_team",
			Message:          "error validating fieldmask in UpdateUserTeam",
			EntityIdentifier: "user_team",
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
		if !strings.Contains(req.UpdateMask.String(), "user_team.uuid") {
			req.UpdateMask.Append(req, "user_team.uuid")
		}

		newReq := &pb.UpdateUserTeamRequest{}
		fieldmask.Update(req.UpdateMask, newReq, req)
		req = newReq
	}

	res, err := s.core.UserTeam().Update(ctx, types.UpsertRequest{
		UserTeam: pbmapper.UserTeamFromProto(req.GetUserTeam()),
	}, !isFull)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "update_user_team",
			Message:          "error calling core upsert in UpdateUserTeam",
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
			ActionIdentifier: "update_user_team",
			Message:          "error fetching after upsert in UpdateUserTeam",
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
			ActionIdentifier: "update_user_team",
			Message:          "entity not found after upsert in UpdateUserTeam",
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
		ActionIdentifier: "update_user_team",
		Message:          "successfully handled UpdateUserTeam",
		EntityIdentifier: "user_team",
		Layer:            monitoring.ProtocolServiceLayer,
		LayerSubtype:     "protobuf",
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return pbmapper.UserTeamToProto(fetchRes.Results[0]), nil
}
