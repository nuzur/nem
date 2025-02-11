package server

import (
	"context"
	"errors"
	"github.com/nuzur/nem/core/module/user_connection/types"
	pb "github.com/nuzur/nem/idl/gen"
	pbmapper "github.com/nuzur/nem/idl/mapper"
	"github.com/nuzur/nem/monitoring"
)

func (s *server) CreateUserConnection(ctx context.Context, req *pb.CreateUserConnectionRequest) (*pb.UserConnection, error) {
	res, err := s.core.UserConnection().Insert(ctx, types.UpsertRequest{
		UserConnection: pbmapper.UserConnectionFromProto(req.GetUserConnection()),
	})
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "create_user_connection",
			Message:          "error calling core upsert in CreateUserConnection",
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
			ActionIdentifier: "create_user_connection",
			Message:          "error fetching after upsert in CreateUserConnection",
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
			ActionIdentifier: "create_user_connection",
			Message:          "entity not found after upsert in CreateUserConnection",
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
		ActionIdentifier: "create_user_connection",
		Message:          "successfully handled CreateUserConnection",
		EntityIdentifier: "user_connection",
		Layer:            monitoring.ProtocolServiceLayer,
		LayerSubtype:     "protobuf",
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return pbmapper.UserConnectionToProto(fetchRes.Results[0]), nil
}
