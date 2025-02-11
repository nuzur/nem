package server

import (
	"context"
	"errors"
	"github.com/nuzur/nem/core/module/extension_version/types"
	pb "github.com/nuzur/nem/idl/gen"
	pbmapper "github.com/nuzur/nem/idl/mapper"
	"github.com/nuzur/nem/monitoring"
)

func (s *server) CreateExtensionVersion(ctx context.Context, req *pb.CreateExtensionVersionRequest) (*pb.ExtensionVersion, error) {
	res, err := s.core.ExtensionVersion().Insert(ctx, types.UpsertRequest{
		ExtensionVersion: pbmapper.ExtensionVersionFromProto(req.GetExtensionVersion()),
	})
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "create_extension_version",
			Message:          "error calling core upsert in CreateExtensionVersion",
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
			ActionIdentifier: "create_extension_version",
			Message:          "error fetching after upsert in CreateExtensionVersion",
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
			ActionIdentifier: "create_extension_version",
			Message:          "entity not found after upsert in CreateExtensionVersion",
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
		ActionIdentifier: "create_extension_version",
		Message:          "successfully handled CreateExtensionVersion",
		EntityIdentifier: "extension_version",
		Layer:            monitoring.ProtocolServiceLayer,
		LayerSubtype:     "protobuf",
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return pbmapper.ExtensionVersionToProto(fetchRes.Results[0]), nil
}
