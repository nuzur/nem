package server

import (
	"context"
	"errors"
	"nem/core/module/extension_execution/types"
	pb "nem/idl/gen"
	pbmapper "nem/idl/mapper"
	"nem/monitoring"
)

func (s *server) CreateExtensionExecution(ctx context.Context, req *pb.CreateExtensionExecutionRequest) (*pb.ExtensionExecution, error) {
	res, err := s.core.ExtensionExecution().Insert(ctx, types.UpsertRequest{
		ExtensionExecution: pbmapper.ExtensionExecutionFromProto(req.GetExtensionExecution()),
	})
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "create_extension_execution",
			Message:          "error calling core upsert in CreateExtensionExecution",
			EntityIdentifier: "extension_execution",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return nil, err
	}

	fetchRes, err := s.core.ExtensionExecution().FetchExtensionExecutionByUUID(ctx, types.FetchExtensionExecutionByUUIDRequest(res))
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "create_extension_execution",
			Message:          "error fetching after upsert in CreateExtensionExecution",
			EntityIdentifier: "extension_execution",
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
			ActionIdentifier: "create_extension_execution",
			Message:          "entity not found after upsert in CreateExtensionExecution",
			EntityIdentifier: "extension_execution",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return nil, err
	}

	s.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "create_extension_execution",
		Message:          "successfully handled CreateExtensionExecution",
		EntityIdentifier: "extension_execution",
		Layer:            monitoring.ProtocolServiceLayer,
		LayerSubtype:     "protobuf",
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return pbmapper.ExtensionExecutionToProto(fetchRes.Results[0]), nil
}
