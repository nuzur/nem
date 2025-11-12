package server

import (
	"context"
	"errors"
	"github.com/nuzur/nem/core/module/ai_usage/types"
	pb "github.com/nuzur/nem/idl/gen"
	pbmapper "github.com/nuzur/nem/idl/mapper"
	"github.com/nuzur/nem/monitoring"
	"go.einride.tech/aip/fieldmask"
	"strings"
)

func (s *server) UpdateAiUsage(ctx context.Context, req *pb.UpdateAiUsageRequest) (*pb.AiUsage, error) {

	if req.AiUsage.Uuid == "" {
		return nil, errors.New("please provide a valid UUID to update")
	}

	err := fieldmask.Validate(req.UpdateMask, req)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "update_ai_usage",
			Message:          "error validating fieldmask in UpdateAiUsage",
			EntityIdentifier: "ai_usage",
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
		if !strings.Contains(req.UpdateMask.String(), "ai_usage.uuid") {
			req.UpdateMask.Append(req, "ai_usage.uuid")
		}

		newReq := &pb.UpdateAiUsageRequest{}
		fieldmask.Update(req.UpdateMask, newReq, req)
		req = newReq
	}

	res, err := s.core.AiUsage().Update(ctx, types.UpsertRequest{
		AiUsage: pbmapper.AiUsageFromProto(req.GetAiUsage()),
	}, !isFull)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "update_ai_usage",
			Message:          "error calling core upsert in UpdateAiUsage",
			EntityIdentifier: "ai_usage",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return nil, err
	}

	fetchRes, err := s.core.AiUsage().FetchAiUsageByUUID(ctx, types.FetchAiUsageByUUIDRequest(res))
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "update_ai_usage",
			Message:          "error fetching after upsert in UpdateAiUsage",
			EntityIdentifier: "ai_usage",
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
			ActionIdentifier: "update_ai_usage",
			Message:          "entity not found after upsert in UpdateAiUsage",
			EntityIdentifier: "ai_usage",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return nil, err
	}

	s.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "update_ai_usage",
		Message:          "successfully handled UpdateAiUsage",
		EntityIdentifier: "ai_usage",
		Layer:            monitoring.ProtocolServiceLayer,
		LayerSubtype:     "protobuf",
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return pbmapper.AiUsageToProto(fetchRes.Results[0]), nil
}
