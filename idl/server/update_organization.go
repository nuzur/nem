package server

import (
	"context"
	"errors"
	"go.einride.tech/aip/fieldmask"
	"nem/core/module/organization/types"
	pb "nem/idl/gen"
	pbmapper "nem/idl/mapper"
	"nem/monitoring"
	"strings"
)

func (s *server) UpdateOrganization(ctx context.Context, req *pb.UpdateOrganizationRequest) (*pb.Organization, error) {

	if req.Organization.Uuid == "" {
		return nil, errors.New("please provide a valid UUID to update")
	}

	if req.Organization.Version == 0 {
		return nil, errors.New("please provide a valid version to update")
	}

	err := fieldmask.Validate(req.UpdateMask, req)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "update_organization",
			Message:          "error validating fieldmask in UpdateOrganization",
			EntityIdentifier: "organization",
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
		if !strings.Contains(req.UpdateMask.String(), "organization.uuid") {
			req.UpdateMask.Append(req, "organization.uuid")
		}

		if !strings.Contains(req.UpdateMask.String(), "organization.version") {
			req.UpdateMask.Append(req, "organization.version")
		}

		newReq := &pb.UpdateOrganizationRequest{}
		fieldmask.Update(req.UpdateMask, newReq, req)
		req = newReq
	}

	res, err := s.core.Organization().Update(ctx, types.UpsertRequest{
		Organization: pbmapper.OrganizationFromProto(req.GetOrganization()),
	}, !isFull)
	if err != nil {
		s.monitoring.Emit(monitoring.EmitRequest{
			ActionIdentifier: "update_organization",
			Message:          "error calling core upsert in UpdateOrganization",
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
			ActionIdentifier: "update_organization",
			Message:          "error fetching after upsert in UpdateOrganization",
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
			ActionIdentifier: "update_organization",
			Message:          "entity not found after upsert in UpdateOrganization",
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
		ActionIdentifier: "update_organization",
		Message:          "successfully handled UpdateOrganization",
		EntityIdentifier: "organization",
		Layer:            monitoring.ProtocolServiceLayer,
		LayerSubtype:     "protobuf",
		Type:             monitoring.EmitTypeSuccess,
		Data:             req,
	})
	return pbmapper.OrganizationToProto(fetchRes.Results[0]), nil
}
