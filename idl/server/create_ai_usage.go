package server

import (
	"context"
	"errors"
	"github.com/nuzur/nem/core/module/ai_usage/types"
	pb "github.com/nuzur/nem/idl/gen"
	pbmapper "github.com/nuzur/nem/idl/mapper"
	"github.com/nuzur/nem/monitoring"
)

func (s *server) CreateAiUsage(ctx context.Context, req *pb.CreateAiUsageRequest) (*pb.AiUsage, error) {
	res, err := s.core.AiUsage().Insert(ctx, types.UpsertRequest{
		AiUsage: pbmapper.AiUsageFromProto(req.GetAiUsage()),
	})
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "create_ai_usage",
			Message:          "error calling core upsert in CreateAiUsage",
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
			ActionIdentifier: "create_ai_usage",
			Message:          "error fetching after upsert in CreateAiUsage",
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
			ActionIdentifier: "create_ai_usage",
			Message:          "entity not found after upsert in CreateAiUsage",
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
		ActionIdentifier: "create_ai_usage",
		Message:          "successfully handled CreateAiUsage",
		EntityIdentifier: "ai_usage",
		Layer:            monitoring.ProtocolServiceLayer,
		LayerSubtype:     "protobuf",
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return pbmapper.AiUsageToProto(fetchRes.Results[0]), nil
}
