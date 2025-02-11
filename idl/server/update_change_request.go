package server

import (
	"context"
	"errors"
	"github.com/nuzur/nem/core/module/change_request/types"
	pb "github.com/nuzur/nem/idl/gen"
	pbmapper "github.com/nuzur/nem/idl/mapper"
	"github.com/nuzur/nem/monitoring"
	"go.einride.tech/aip/fieldmask"
	"strings"
)

func (s *server) UpdateChangeRequest(ctx context.Context, req *pb.UpdateChangeRequestRequest) (*pb.ChangeRequest, error) {

	if req.ChangeRequest.Uuid == "" {
		return nil, errors.New("please provide a valid UUID to update")
	}

	if req.ChangeRequest.Version == 0 {
		return nil, errors.New("please provide a valid version to update")
	}

	err := fieldmask.Validate(req.UpdateMask, req)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "update_change_request",
			Message:          "error validating fieldmask in UpdateChangeRequest",
			EntityIdentifier: "change_request",
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
		if !strings.Contains(req.UpdateMask.String(), "change_request.uuid") {
			req.UpdateMask.Append(req, "change_request.uuid")
		}

		if !strings.Contains(req.UpdateMask.String(), "change_request.version") {
			req.UpdateMask.Append(req, "change_request.version")
		}

		newReq := &pb.UpdateChangeRequestRequest{}
		fieldmask.Update(req.UpdateMask, newReq, req)
		req = newReq
	}

	res, err := s.core.ChangeRequest().Update(ctx, types.UpsertRequest{
		ChangeRequest: pbmapper.ChangeRequestFromProto(req.GetChangeRequest()),
	}, !isFull)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "update_change_request",
			Message:          "error calling core upsert in UpdateChangeRequest",
			EntityIdentifier: "change_request",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return nil, err
	}

	fetchRes, err := s.core.ChangeRequest().FetchChangeRequestByUUID(ctx, types.FetchChangeRequestByUUIDRequest(res))
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "update_change_request",
			Message:          "error fetching after upsert in UpdateChangeRequest",
			EntityIdentifier: "change_request",
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
			ActionIdentifier: "update_change_request",
			Message:          "entity not found after upsert in UpdateChangeRequest",
			EntityIdentifier: "change_request",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return nil, err
	}

	s.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "update_change_request",
		Message:          "successfully handled UpdateChangeRequest",
		EntityIdentifier: "change_request",
		Layer:            monitoring.ProtocolServiceLayer,
		LayerSubtype:     "protobuf",
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return pbmapper.ChangeRequestToProto(fetchRes.Results[0]), nil
}
