package server

import (
	"context"
	"errors"
	"github.com/nuzur/nem/core/module/user_project_version/types"
	pb "github.com/nuzur/nem/idl/gen"
	pbmapper "github.com/nuzur/nem/idl/mapper"
	"github.com/nuzur/nem/monitoring"
	"go.einride.tech/aip/fieldmask"
	"strings"
)

func (s *server) UpdateUserProjectVersion(ctx context.Context, req *pb.UpdateUserProjectVersionRequest) (*pb.UserProjectVersion, error) {

	if req.UserProjectVersion.Uuid == "" {
		return nil, errors.New("please provide a valid UUID to update")
	}

	if req.UserProjectVersion.Version == 0 {
		return nil, errors.New("please provide a valid version to update")
	}

	err := fieldmask.Validate(req.UpdateMask, req)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "update_user_project_version",
			Message:          "error validating fieldmask in UpdateUserProjectVersion",
			EntityIdentifier: "user_project_version",
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
		if !strings.Contains(req.UpdateMask.String(), "user_project_version.uuid") {
			req.UpdateMask.Append(req, "user_project_version.uuid")
		}

		if !strings.Contains(req.UpdateMask.String(), "user_project_version.version") {
			req.UpdateMask.Append(req, "user_project_version.version")
		}

		newReq := &pb.UpdateUserProjectVersionRequest{}
		fieldmask.Update(req.UpdateMask, newReq, req)
		req = newReq
	}

	res, err := s.core.UserProjectVersion().Update(ctx, types.UpsertRequest{
		UserProjectVersion: pbmapper.UserProjectVersionFromProto(req.GetUserProjectVersion()),
	}, !isFull)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "update_user_project_version",
			Message:          "error calling core upsert in UpdateUserProjectVersion",
			EntityIdentifier: "user_project_version",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return nil, err
	}

	fetchRes, err := s.core.UserProjectVersion().FetchUserProjectVersionByUUID(ctx, types.FetchUserProjectVersionByUUIDRequest(res))
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "update_user_project_version",
			Message:          "error fetching after upsert in UpdateUserProjectVersion",
			EntityIdentifier: "user_project_version",
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
			ActionIdentifier: "update_user_project_version",
			Message:          "entity not found after upsert in UpdateUserProjectVersion",
			EntityIdentifier: "user_project_version",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return nil, err
	}

	s.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "update_user_project_version",
		Message:          "successfully handled UpdateUserProjectVersion",
		EntityIdentifier: "user_project_version",
		Layer:            monitoring.ProtocolServiceLayer,
		LayerSubtype:     "protobuf",
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return pbmapper.UserProjectVersionToProto(fetchRes.Results[0]), nil
}
