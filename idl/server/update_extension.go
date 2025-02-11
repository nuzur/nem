package server

import (
	"context"
	"errors"
	"github.com/nuzur/nem/core/module/extension/types"
	pb "github.com/nuzur/nem/idl/gen"
	pbmapper "github.com/nuzur/nem/idl/mapper"
	"github.com/nuzur/nem/monitoring"
	"go.einride.tech/aip/fieldmask"
	"strings"
)

func (s *server) UpdateExtension(ctx context.Context, req *pb.UpdateExtensionRequest) (*pb.Extension, error) {

	if req.Extension.Uuid == "" {
		return nil, errors.New("please provide a valid UUID to update")
	}

	if req.Extension.Version == 0 {
		return nil, errors.New("please provide a valid version to update")
	}

	err := fieldmask.Validate(req.UpdateMask, req)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "update_extension",
			Message:          "error validating fieldmask in UpdateExtension",
			EntityIdentifier: "extension",
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
		if !strings.Contains(req.UpdateMask.String(), "extension.uuid") {
			req.UpdateMask.Append(req, "extension.uuid")
		}

		if !strings.Contains(req.UpdateMask.String(), "extension.version") {
			req.UpdateMask.Append(req, "extension.version")
		}

		newReq := &pb.UpdateExtensionRequest{}
		fieldmask.Update(req.UpdateMask, newReq, req)
		req = newReq
	}

	res, err := s.core.Extension().Update(ctx, types.UpsertRequest{
		Extension: pbmapper.ExtensionFromProto(req.GetExtension()),
	}, !isFull)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "update_extension",
			Message:          "error calling core upsert in UpdateExtension",
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
			ActionIdentifier: "update_extension",
			Message:          "error fetching after upsert in UpdateExtension",
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
			ActionIdentifier: "update_extension",
			Message:          "entity not found after upsert in UpdateExtension",
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
		ActionIdentifier: "update_extension",
		Message:          "successfully handled UpdateExtension",
		EntityIdentifier: "extension",
		Layer:            monitoring.ProtocolServiceLayer,
		LayerSubtype:     "protobuf",
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return pbmapper.ExtensionToProto(fetchRes.Results[0]), nil
}
