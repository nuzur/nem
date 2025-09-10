package server

import (
	"context"
	"errors"
	"github.com/nuzur/nem/core/module/membership/types"
	pb "github.com/nuzur/nem/idl/gen"
	pbmapper "github.com/nuzur/nem/idl/mapper"
	"github.com/nuzur/nem/monitoring"
)

func (s *server) CreateMembership(ctx context.Context, req *pb.CreateMembershipRequest) (*pb.Membership, error) {
	res, err := s.core.Membership().Insert(ctx, types.UpsertRequest{
		Membership: pbmapper.MembershipFromProto(req.GetMembership()),
	})
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "create_membership",
			Message:          "error calling core upsert in CreateMembership",
			EntityIdentifier: "membership",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return nil, err
	}

	fetchRes, err := s.core.Membership().FetchMembershipByUUID(ctx, types.FetchMembershipByUUIDRequest(res))
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "create_membership",
			Message:          "error fetching after upsert in CreateMembership",
			EntityIdentifier: "membership",
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
			ActionIdentifier: "create_membership",
			Message:          "entity not found after upsert in CreateMembership",
			EntityIdentifier: "membership",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return nil, err
	}

	s.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "create_membership",
		Message:          "successfully handled CreateMembership",
		EntityIdentifier: "membership",
		Layer:            monitoring.ProtocolServiceLayer,
		LayerSubtype:     "protobuf",
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return pbmapper.MembershipToProto(fetchRes.Results[0]), nil
}
