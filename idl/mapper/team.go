package mapper

import (
	main_entity "nem/core/entity/team"
	pb "nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func TeamToProto(e main_entity.Team) *pb.Team {
	return &pb.Team{
		Uuid:             e.UUID.String(),
		Version:          int64(e.Version),
		Name:             e.Name,
		Enviorments:      EnviormentSliceToProto(e.Enviorments),
		ReviewConfigs:    ReviewConfigSliceToProto(e.ReviewConfigs),
		Memberships:      MembershipSliceToProto(e.Memberships),
		Stores:           StoreSliceToProto(e.Stores),
		Connections:      ConnectionSliceToProto(e.Connections),
		ObjectStores:     ObjectStoreSliceToProto(e.ObjectStores),
		OrganizationUuid: e.OrganizationUUID.String(),
		DefaultEntity:    EntityToProto(e.DefaultEntity),
		Status:           pb.TeamStatus(e.Status),
		CreatedAt:        timestamppb.New(e.CreatedAt),
		UpdatedAt:        timestamppb.New(e.UpdatedAt),
		CreatedByUuid:    e.CreatedByUUID.String(),
		UpdatedByUuid:    e.UpdatedByUUID.String(),
	}
}

func TeamSliceToProto(es []main_entity.Team) []*pb.Team {
	res := []*pb.Team{}
	for _, e := range es {
		res = append(res, TeamToProto(e))
	}
	return res
}

func TeamFromProto(m *pb.Team) main_entity.Team {
	if m == nil {
		return main_entity.Team{}
	}
	return main_entity.Team{
		UUID:             uuid.FromStringOrNil(m.GetUuid()),
		Version:          int64(m.GetVersion()),
		Name:             m.GetName(),
		Enviorments:      EnviormentSliceFromProto(m.GetEnviorments()),
		ReviewConfigs:    ReviewConfigSliceFromProto(m.GetReviewConfigs()),
		Memberships:      MembershipSliceFromProto(m.GetMemberships()),
		Stores:           StoreSliceFromProto(m.GetStores()),
		Connections:      ConnectionSliceFromProto(m.GetConnections()),
		ObjectStores:     ObjectStoreSliceFromProto(m.GetObjectStores()),
		OrganizationUUID: uuid.FromStringOrNil(m.GetOrganizationUuid()),
		DefaultEntity:    EntityFromProto(m.GetDefaultEntity()),
		Status:           main_entity.Status(m.GetStatus()),
		CreatedAt:        m.GetCreatedAt().AsTime(),
		UpdatedAt:        m.GetUpdatedAt().AsTime(),
		CreatedByUUID:    uuid.FromStringOrNil(m.GetCreatedByUuid()),
		UpdatedByUUID:    uuid.FromStringOrNil(m.GetUpdatedByUuid()),
	}
}

func TeamSliceFromProto(es []*pb.Team) []main_entity.Team {
	if es == nil {
		return []main_entity.Team{}
	}
	res := []main_entity.Team{}
	for _, e := range es {
		res = append(res, TeamFromProto(e))
	}
	return res
}
