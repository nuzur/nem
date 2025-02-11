package server

import (
	"context"
	"errors"
	"nem/core/module/extension/types"
	pb "nem/idl/gen"
	pbmapper "nem/idl/mapper"
	"nem/monitoring"
)

func (s *server) CreateExtension(ctx context.Context, req *pb.CreateExtensionRequest) (*pb.Extension, error) {
	res, err := s.core.Extension().Insert(ctx, types.UpsertRequest{
		Extension: pbmapper.ExtensionFromProto(req.GetExtension()),
	})
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "create_extension",
			Message:          "error calling core upsert in CreateExtension",
			EntityIdentifier: "extension",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return nil, err
	}

	fetchRes, err := s.core.Extension().FetchExtensionByUUID(ctx, types.FetchExtensionByUUIDRequest(res))
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "create_extension",
			Message:          "error fetching after upsert in CreateExtension",
			EntityIdentifier: "extension",
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
			ActionIdentifier: "create_extension",
			Message:          "entity not found after upsert in CreateExtension",
			EntityIdentifier: "extension",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return nil, err
	}

	s.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "create_extension",
		Message:          "successfully handled CreateExtension",
		EntityIdentifier: "extension",
		Layer:            monitoring.ProtocolServiceLayer,
		LayerSubtype:     "protobuf",
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return pbmapper.ExtensionToProto(fetchRes.Results[0]), nil
}
