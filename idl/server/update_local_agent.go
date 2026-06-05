package server

import (
	"context"
	"errors"
	"github.com/nuzur/nem/core/module/local_agent/types"
	pb "github.com/nuzur/nem/idl/gen"
	pbmapper "github.com/nuzur/nem/idl/mapper"
	"github.com/nuzur/nem/monitoring"
	"go.einride.tech/aip/fieldmask"
	"strings"
)

func (s *server) UpdateLocalAgent(ctx context.Context, req *pb.UpdateLocalAgentRequest) (*pb.LocalAgent, error) {

	if req.LocalAgent.Uuid == "" {
		return nil, errors.New("please provide a valid UUID to update")
	}

	err := fieldmask.Validate(req.UpdateMask, req)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "update_local_agent",
			Message:          "error validating fieldmask in UpdateLocalAgent",
			EntityIdentifier: "local_agent",
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
		if !strings.Contains(req.UpdateMask.String(), "local_agent.uuid") {
			req.UpdateMask.Append(req, "local_agent.uuid")
		}

		newReq := &pb.UpdateLocalAgentRequest{}
		fieldmask.Update(req.UpdateMask, newReq, req)
		req = newReq
	}

	res, err := s.core.LocalAgent().Update(ctx, types.UpsertRequest{
		LocalAgent: pbmapper.LocalAgentFromProto(req.GetLocalAgent()),
	}, !isFull)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "update_local_agent",
			Message:          "error calling core upsert in UpdateLocalAgent",
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
			ActionIdentifier: "update_local_agent",
			Message:          "error fetching after upsert in UpdateLocalAgent",
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
			ActionIdentifier: "update_local_agent",
			Message:          "entity not found after upsert in UpdateLocalAgent",
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
		ActionIdentifier: "update_local_agent",
		Message:          "successfully handled UpdateLocalAgent",
		EntityIdentifier: "local_agent",
		Layer:            monitoring.ProtocolServiceLayer,
		LayerSubtype:     "protobuf",
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return pbmapper.LocalAgentToProto(fetchRes.Results[0]), nil
}
