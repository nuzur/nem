package server

import (
	"context"
	"errors"
	"github.com/nuzur/nem/core/module/extension_version/types"
	pb "github.com/nuzur/nem/idl/gen"
	pbmapper "github.com/nuzur/nem/idl/mapper"
	"github.com/nuzur/nem/monitoring"
	"go.einride.tech/aip/fieldmask"
	"strings"
)

func (s *server) UpdateExtensionVersion(ctx context.Context, req *pb.UpdateExtensionVersionRequest) (*pb.ExtensionVersion, error) {

	if req.ExtensionVersion.Uuid == "" {
		return nil, errors.New("please provide a valid UUID to update")
	}

	if req.ExtensionVersion.Version == 0 {
		return nil, errors.New("please provide a valid version to update")
	}

	err := fieldmask.Validate(req.UpdateMask, req)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "update_extension_version",
			Message:          "error validating fieldmask in UpdateExtensionVersion",
			EntityIdentifier: "extension_version",
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
		if !strings.Contains(req.UpdateMask.String(), "extension_version.uuid") {
			req.UpdateMask.Append(req, "extension_version.uuid")
		}

		if !strings.Contains(req.UpdateMask.String(), "extension_version.version") {
			req.UpdateMask.Append(req, "extension_version.version")
		}

		newReq := &pb.UpdateExtensionVersionRequest{}
		fieldmask.Update(req.UpdateMask, newReq, req)
		req = newReq
	}

	res, err := s.core.ExtensionVersion().Update(ctx, types.UpsertRequest{
		ExtensionVersion: pbmapper.ExtensionVersionFromProto(req.GetExtensionVersion()),
	}, !isFull)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "update_extension_version",
			Message:          "error calling core upsert in UpdateExtensionVersion",
			EntityIdentifier: "extension_version",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return nil, err
	}

	fetchRes, err := s.core.ExtensionVersion().FetchExtensionVersionByUUID(ctx, types.FetchExtensionVersionByUUIDRequest(res))
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "update_extension_version",
			Message:          "error fetching after upsert in UpdateExtensionVersion",
			EntityIdentifier: "extension_version",
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
			ActionIdentifier: "update_extension_version",
			Message:          "entity not found after upsert in UpdateExtensionVersion",
			EntityIdentifier: "extension_version",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return nil, err
	}

	s.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "update_extension_version",
		Message:          "successfully handled UpdateExtensionVersion",
		EntityIdentifier: "extension_version",
		Layer:            monitoring.ProtocolServiceLayer,
		LayerSubtype:     "protobuf",
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return pbmapper.ExtensionVersionToProto(fetchRes.Results[0]), nil
}
