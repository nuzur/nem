package server

import (
	"context"
	"errors"
	"github.com/nuzur/nem/core/module/project/types"
	pb "github.com/nuzur/nem/idl/gen"
	pbmapper "github.com/nuzur/nem/idl/mapper"
	"github.com/nuzur/nem/monitoring"
)

func (s *server) CreateProject(ctx context.Context, req *pb.CreateProjectRequest) (*pb.Project, error) {
	res, err := s.core.Project().Insert(ctx, types.UpsertRequest{
		Project: pbmapper.ProjectFromProto(req.GetProject()),
	})
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "create_project",
			Message:          "error calling core upsert in CreateProject",
			EntityIdentifier: "project",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return nil, err
	}

	fetchRes, err := s.core.Project().FetchProjectByUUID(ctx, types.FetchProjectByUUIDRequest(res))
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "create_project",
			Message:          "error fetching after upsert in CreateProject",
			EntityIdentifier: "project",
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
			ActionIdentifier: "create_project",
			Message:          "entity not found after upsert in CreateProject",
			EntityIdentifier: "project",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return nil, err
	}

	s.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "create_project",
		Message:          "successfully handled CreateProject",
		EntityIdentifier: "project",
		Layer:            monitoring.ProtocolServiceLayer,
		LayerSubtype:     "protobuf",
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return pbmapper.ProjectToProto(fetchRes.Results[0]), nil
}
