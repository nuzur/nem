package server

import (
	"context"
	"errors"
	"github.com/nuzur/nem/core/module/local_agent/types"
	pb "github.com/nuzur/nem/idl/gen"
	pbmapper "github.com/nuzur/nem/idl/mapper"
	"github.com/nuzur/nem/monitoring"
)

func (s *server) CreateLocalAgent(ctx context.Context, req *pb.CreateLocalAgentRequest) (*pb.LocalAgent, error) {
	res, err := s.core.LocalAgent().Insert(ctx, types.UpsertRequest{
		LocalAgent: pbmapper.LocalAgentFromProto(req.GetLocalAgent()),
	})
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "create_local_agent",
			Message:          "error calling core upsert in CreateLocalAgent",
			EntityIdentifier: "local_agent",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return nil, err
	}

	fetchRes, err := s.core.LocalAgent().FetchLocalAgentByUUID(ctx, types.FetchLocalAgentByUUIDRequest(res))
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "create_local_agent",
			Message:          "error fetching after upsert in CreateLocalAgent",
			EntityIdentifier: "local_agent",
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
			ActionIdentifier: "create_local_agent",
			Message:          "entity not found after upsert in CreateLocalAgent",
			EntityIdentifier: "local_agent",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return nil, err
	}

	s.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "create_local_agent",
		Message:          "successfully handled CreateLocalAgent",
		EntityIdentifier: "local_agent",
		Layer:            monitoring.ProtocolServiceLayer,
		LayerSubtype:     "protobuf",
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return pbmapper.LocalAgentToProto(fetchRes.Results[0]), nil
}
