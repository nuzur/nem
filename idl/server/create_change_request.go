package server

import (
	"context"
	"errors"
	"nem/core/module/change_request/types"
	pb "nem/idl/gen"
	pbmapper "nem/idl/mapper"
	"nem/monitoring"
)

func (s *server) CreateChangeRequest(ctx context.Context, req *pb.CreateChangeRequestRequest) (*pb.ChangeRequest, error) {
	res, err := s.core.ChangeRequest().Insert(ctx, types.UpsertRequest{
		ChangeRequest: pbmapper.ChangeRequestFromProto(req.GetChangeRequest()),
	})
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "create_change_request",
			Message:          "error calling core upsert in CreateChangeRequest",
			EntityIdentifier: "change_request",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return nil, err
	}

	fetchRes, err := s.core.ChangeRequest().FetchChangeRequestByUUID(ctx, types.FetchChangeRequestByUUIDRequest(res))
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "create_change_request",
			Message:          "error fetching after upsert in CreateChangeRequest",
			EntityIdentifier: "change_request",
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
			ActionIdentifier: "create_change_request",
			Message:          "entity not found after upsert in CreateChangeRequest",
			EntityIdentifier: "change_request",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return nil, err
	}

	s.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "create_change_request",
		Message:          "successfully handled CreateChangeRequest",
		EntityIdentifier: "change_request",
		Layer:            monitoring.ProtocolServiceLayer,
		LayerSubtype:     "protobuf",
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return pbmapper.ChangeRequestToProto(fetchRes.Results[0]), nil
}
