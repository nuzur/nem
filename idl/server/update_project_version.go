package server

import (
	"context"
	"errors"
	"github.com/nuzur/nem/core/module/project_version/types"
	pb "github.com/nuzur/nem/idl/gen"
	pbmapper "github.com/nuzur/nem/idl/mapper"
	"github.com/nuzur/nem/monitoring"
	"go.einride.tech/aip/fieldmask"
	"strings"
)

func (s *server) UpdateProjectVersion(ctx context.Context, req *pb.UpdateProjectVersionRequest) (*pb.ProjectVersion, error) {

	if req.ProjectVersion.Uuid == "" {
		return nil, errors.New("please provide a valid UUID to update")
	}

	if req.ProjectVersion.Version == 0 {
		return nil, errors.New("please provide a valid version to update")
	}

	err := fieldmask.Validate(req.UpdateMask, req)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "update_project_version",
			Message:          "error validating fieldmask in UpdateProjectVersion",
			EntityIdentifier: "project_version",
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
		if !strings.Contains(req.UpdateMask.String(), "project_version.uuid") {
			req.UpdateMask.Append(req, "project_version.uuid")
		}

		if !strings.Contains(req.UpdateMask.String(), "project_version.version") {
			req.UpdateMask.Append(req, "project_version.version")
		}

		newReq := &pb.UpdateProjectVersionRequest{}
		fieldmask.Update(req.UpdateMask, newReq, req)
		req = newReq
	}

	res, err := s.core.ProjectVersion().Update(ctx, types.UpsertRequest{
		ProjectVersion: pbmapper.ProjectVersionFromProto(req.GetProjectVersion()),
	}, !isFull)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "update_project_version",
			Message:          "error calling core upsert in UpdateProjectVersion",
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
			ActionIdentifier: "update_project_version",
			Message:          "error fetching after upsert in UpdateProjectVersion",
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
			ActionIdentifier: "update_project_version",
			Message:          "entity not found after upsert in UpdateProjectVersion",
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
		ActionIdentifier: "update_project_version",
		Message:          "successfully handled UpdateProjectVersion",
		EntityIdentifier: "project_version",
		Layer:            monitoring.ProtocolServiceLayer,
		LayerSubtype:     "protobuf",
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return pbmapper.ProjectVersionToProto(fetchRes.Results[0]), nil
}
