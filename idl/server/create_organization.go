package server

import (
	"context"
	"errors"
	"nem/core/module/organization/types"
	pb "nem/idl/gen"
	pbmapper "nem/idl/mapper"
	"nem/monitoring"
)

func (s *server) CreateOrganization(ctx context.Context, req *pb.CreateOrganizationRequest) (*pb.Organization, error) {
	res, err := s.core.Organization().Insert(ctx, types.UpsertRequest{
		Organization: pbmapper.OrganizationFromProto(req.GetOrganization()),
	})
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "create_organization",
			Message:          "error calling core upsert in CreateOrganization",
			EntityIdentifier: "organization",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return nil, err
	}

	fetchRes, err := s.core.Organization().FetchOrganizationByUUID(ctx, types.FetchOrganizationByUUIDRequest(res))
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "create_organization",
			Message:          "error fetching after upsert in CreateOrganization",
			EntityIdentifier: "organization",
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
			ActionIdentifier: "create_organization",
			Message:          "entity not found after upsert in CreateOrganization",
			EntityIdentifier: "organization",
			Layer:            monitoring.ProtocolServiceLayer,
			LayerSubtype:     "protobuf",
			Type:             monitoring.EmitTypeError,
			Data:             req,
			Error:            err,
		})
		return nil, err
	}

	s.monitoring.Emit(monitoring.EmitRequest{
		ActionIdentifier: "create_organization",
		Message:          "successfully handled CreateOrganization",
		EntityIdentifier: "organization",
		Layer:            monitoring.ProtocolServiceLayer,
		LayerSubtype:     "protobuf",
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return pbmapper.OrganizationToProto(fetchRes.Results[0]), nil
}
