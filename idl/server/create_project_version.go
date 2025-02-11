package server

import (
	"context"
	"errors"
	"github.com/nuzur/nem/core/module/project_version/types"
	pb "github.com/nuzur/nem/idl/gen"
	pbmapper "github.com/nuzur/nem/idl/mapper"
	"github.com/nuzur/nem/monitoring"
)

func (s *server) CreateProjectVersion(ctx context.Context, req *pb.CreateProjectVersionRequest) (*pb.ProjectVersion, error) {
	res, err := s.core.ProjectVersion().Insert(ctx, types.UpsertRequest{
		ProjectVersion: pbmapper.ProjectVersionFromProto(req.GetProjectVersion()),
	})
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "create_project_version",
			Message:          "error calling core upsert in CreateProjectVersion",
			EntityIdentifier: "project_version",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return nil, err
	}

	fetchRes, err := s.core.ProjectVersion().FetchProjectVersionByUUID(ctx, types.FetchProjectVersionByUUIDRequest(res))
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "create_project_version",
			Message:          "error fetching after upsert in CreateProjectVersion",
			EntityIdentifier: "project_version",
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
			ActionIdentifier: "create_project_version",
			Message:          "entity not found after upsert in CreateProjectVersion",
			EntityIdentifier: "project_version",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return nil, err
	}

	s.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "create_project_version",
		Message:          "successfully handled CreateProjectVersion",
		EntityIdentifier: "project_version",
		Layer:            monitoring.ProtocolServiceLayer,
		LayerSubtype:     "protobuf",
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return pbmapper.ProjectVersionToProto(fetchRes.Results[0]), nil
}
