package server

import (
	"context"
	"errors"
	"github.com/nuzur/nem/core/module/project/types"
	pb "github.com/nuzur/nem/idl/gen"
	pbmapper "github.com/nuzur/nem/idl/mapper"
	"github.com/nuzur/nem/monitoring"
	"go.einride.tech/aip/fieldmask"
	"strings"
)

func (s *server) UpdateProject(ctx context.Context, req *pb.UpdateProjectRequest) (*pb.Project, error) {

	if req.Project.Uuid == "" {
		return nil, errors.New("please provide a valid UUID to update")
	}

	if req.Project.Version == 0 {
		return nil, errors.New("please provide a valid version to update")
	}

	err := fieldmask.Validate(req.UpdateMask, req)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "update_project",
			Message:          "error validating fieldmask in UpdateProject",
			EntityIdentifier: "project",
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
		if !strings.Contains(req.UpdateMask.String(), "project.uuid") {
			req.UpdateMask.Append(req, "project.uuid")
		}

		if !strings.Contains(req.UpdateMask.String(), "project.version") {
			req.UpdateMask.Append(req, "project.version")
		}

		newReq := &pb.UpdateProjectRequest{}
		fieldmask.Update(req.UpdateMask, newReq, req)
		req = newReq
	}

	res, err := s.core.Project().Update(ctx, types.UpsertRequest{
		Project: pbmapper.ProjectFromProto(req.GetProject()),
	}, !isFull)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "update_project",
			Message:          "error calling core upsert in UpdateProject",
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
			ActionIdentifier: "update_project",
			Message:          "error fetching after upsert in UpdateProject",
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
			ActionIdentifier: "update_project",
			Message:          "entity not found after upsert in UpdateProject",
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
		ActionIdentifier: "update_project",
		Message:          "successfully handled UpdateProject",
		EntityIdentifier: "project",
		Layer:            monitoring.ProtocolServiceLayer,
		LayerSubtype:     "protobuf",
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return pbmapper.ProjectToProto(fetchRes.Results[0]), nil
}
