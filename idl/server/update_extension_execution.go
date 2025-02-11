package server

import (
	"context"
	"errors"
	"go.einride.tech/aip/fieldmask"
	"nem/core/module/extension_execution/types"
	pb "nem/idl/gen"
	pbmapper "nem/idl/mapper"
	"nem/monitoring"
	"strings"
)

func (s *server) UpdateExtensionExecution(ctx context.Context, req *pb.UpdateExtensionExecutionRequest) (*pb.ExtensionExecution, error) {

	if req.ExtensionExecution.Uuid == "" {
		return nil, errors.New("please provide a valid UUID to update")
	}

	err := fieldmask.Validate(req.UpdateMask, req)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "update_extension_execution",
			Message:          "error validating fieldmask in UpdateExtensionExecution",
			EntityIdentifier: "extension_execution",
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
		if !strings.Contains(req.UpdateMask.String(), "extension_execution.uuid") {
			req.UpdateMask.Append(req, "extension_execution.uuid")
		}

		newReq := &pb.UpdateExtensionExecutionRequest{}
		fieldmask.Update(req.UpdateMask, newReq, req)
		req = newReq
	}

	res, err := s.core.ExtensionExecution().Update(ctx, types.UpsertRequest{
		ExtensionExecution: pbmapper.ExtensionExecutionFromProto(req.GetExtensionExecution()),
	}, !isFull)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "update_extension_execution",
			Message:          "error calling core upsert in UpdateExtensionExecution",
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
			ActionIdentifier: "update_extension_execution",
			Message:          "error fetching after upsert in UpdateExtensionExecution",
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
			ActionIdentifier: "update_extension_execution",
			Message:          "entity not found after upsert in UpdateExtensionExecution",
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
		ActionIdentifier: "update_extension_execution",
		Message:          "successfully handled UpdateExtensionExecution",
		EntityIdentifier: "extension_execution",
		Layer:            monitoring.ProtocolServiceLayer,
		LayerSubtype:     "protobuf",
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return pbmapper.ExtensionExecutionToProto(fetchRes.Results[0]), nil
}
