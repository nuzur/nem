package server

import (
	"context"
	"errors"
	"github.com/nuzur/nem/core/module/user_project/types"
	pb "github.com/nuzur/nem/idl/gen"
	pbmapper "github.com/nuzur/nem/idl/mapper"
	"github.com/nuzur/nem/monitoring"
	"go.einride.tech/aip/fieldmask"
	"strings"
)

func (s *server) UpdateUserProject(ctx context.Context, req *pb.UpdateUserProjectRequest) (*pb.UserProject, error) {

	if req.UserProject.Uuid == "" {
		return nil, errors.New("please provide a valid UUID to update")
	}

	err := fieldmask.Validate(req.UpdateMask, req)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "update_user_project",
			Message:          "error validating fieldmask in UpdateUserProject",
			EntityIdentifier: "user_project",
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
		if !strings.Contains(req.UpdateMask.String(), "user_project.uuid") {
			req.UpdateMask.Append(req, "user_project.uuid")
		}

		newReq := &pb.UpdateUserProjectRequest{}
		fieldmask.Update(req.UpdateMask, newReq, req)
		req = newReq
	}

	res, err := s.core.UserProject().Update(ctx, types.UpsertRequest{
		UserProject: pbmapper.UserProjectFromProto(req.GetUserProject()),
	}, !isFull)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "update_user_project",
			Message:          "error calling core upsert in UpdateUserProject",
			EntityIdentifier: "user_project",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return nil, err
	}

	fetchRes, err := s.core.UserProject().FetchUserProjectByUUID(ctx, types.FetchUserProjectByUUIDRequest(res))
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "update_user_project",
			Message:          "error fetching after upsert in UpdateUserProject",
			EntityIdentifier: "user_project",
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
			ActionIdentifier: "update_user_project",
			Message:          "entity not found after upsert in UpdateUserProject",
			EntityIdentifier: "user_project",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return nil, err
	}

	s.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "update_user_project",
		Message:          "successfully handled UpdateUserProject",
		EntityIdentifier: "user_project",
		Layer:            monitoring.ProtocolServiceLayer,
		LayerSubtype:     "protobuf",
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return pbmapper.UserProjectToProto(fetchRes.Results[0]), nil
}
