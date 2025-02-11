package server

import (
	"context"
	"errors"
	"github.com/nuzur/nem/core/module/user_connection/types"
	pb "github.com/nuzur/nem/idl/gen"
	pbmapper "github.com/nuzur/nem/idl/mapper"
	"github.com/nuzur/nem/monitoring"
	"go.einride.tech/aip/fieldmask"
	"strings"
)

func (s *server) UpdateUserConnection(ctx context.Context, req *pb.UpdateUserConnectionRequest) (*pb.UserConnection, error) {

	if req.UserConnection.Uuid == "" {
		return nil, errors.New("please provide a valid UUID to update")
	}

	err := fieldmask.Validate(req.UpdateMask, req)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "update_user_connection",
			Message:          "error validating fieldmask in UpdateUserConnection",
			EntityIdentifier: "user_connection",
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
		if !strings.Contains(req.UpdateMask.String(), "user_connection.uuid") {
			req.UpdateMask.Append(req, "user_connection.uuid")
		}

		newReq := &pb.UpdateUserConnectionRequest{}
		fieldmask.Update(req.UpdateMask, newReq, req)
		req = newReq
	}

	res, err := s.core.UserConnection().Update(ctx, types.UpsertRequest{
		UserConnection: pbmapper.UserConnectionFromProto(req.GetUserConnection()),
	}, !isFull)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "update_user_connection",
			Message:          "error calling core upsert in UpdateUserConnection",
			EntityIdentifier: "user_connection",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return nil, err
	}

	fetchRes, err := s.core.UserConnection().FetchUserConnectionByUUID(ctx, types.FetchUserConnectionByUUIDRequest(res))
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "update_user_connection",
			Message:          "error fetching after upsert in UpdateUserConnection",
			EntityIdentifier: "user_connection",
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
			ActionIdentifier: "update_user_connection",
			Message:          "entity not found after upsert in UpdateUserConnection",
			EntityIdentifier: "user_connection",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return nil, err
	}

	s.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "update_user_connection",
		Message:          "successfully handled UpdateUserConnection",
		EntityIdentifier: "user_connection",
		Layer:            monitoring.ProtocolServiceLayer,
		LayerSubtype:     "protobuf",
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return pbmapper.UserConnectionToProto(fetchRes.Results[0]), nil
}
