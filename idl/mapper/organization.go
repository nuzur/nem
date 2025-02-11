package mapper

import (
	main_entity "nem/core/entity/organization"
	pb "nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func OrganizationToProto(e main_entity.Organization) *pb.Organization {
	return &pb.Organization{
		Uuid:          e.UUID.String(),
		Version:       int64(e.Version),
		Name:          e.Name,
		Domains:       e.Domains,
		AdminUuids:    UUIDSliceToStringSlice(e.AdminUUIDs),
		Memberships:   MembershipSliceToProto(e.Memberships),
		Status:        pb.OrganizationStatus(e.Status),
		CreatedAt:     timestamppb.New(e.CreatedAt),
		UpdatedAt:     timestamppb.New(e.UpdatedAt),
		CreatedByUuid: e.CreatedByUUID.String(),
		UpdatedByUuid: e.UpdatedByUUID.String(),
	}
}

func OrganizationSliceToProto(es []main_entity.Organization) []*pb.Organization {
	res := []*pb.Organization{}
	for _, e := range es {
		res = append(res, OrganizationToProto(e))
	}
	return res
}

func OrganizationFromProto(m *pb.Organization) main_entity.Organization {
	if m == nil {
		return main_entity.Organization{}
	}
	return main_entity.Organization{
		UUID:          uuid.FromStringOrNil(m.GetUuid()),
		Version:       int64(m.GetVersion()),
		Name:          m.GetName(),
		Domains:       m.GetDomains(),
		AdminUUIDs:    StringSliceToUUIDSlice(m.GetAdminUuids()),
		Memberships:   MembershipSliceFromProto(m.GetMemberships()),
		Status:        main_entity.Status(m.GetStatus()),
		CreatedAt:     m.GetCreatedAt().AsTime(),
		UpdatedAt:     m.GetUpdatedAt().AsTime(),
		CreatedByUUID: uuid.FromStringOrNil(m.GetCreatedByUuid()),
		UpdatedByUUID: uuid.FromStringOrNil(m.GetUpdatedByUuid()),
	}
}

func OrganizationSliceFromProto(es []*pb.Organization) []main_entity.Organization {
	if es == nil {
		return []main_entity.Organization{}
	}
	res := []main_entity.Organization{}
	for _, e := range es {
		res = append(res, OrganizationFromProto(e))
	}
	return res
}
