package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/local_agent"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func LocalAgentToProto(e main_entity.LocalAgent) *pb.LocalAgent {
	return &pb.LocalAgent{
		Uuid:          e.UUID.String(),
		UserUuid:      e.UserUUID.String(),
		TokenHash:     StringPtrToString(e.TokenHash),
		MachineName:   StringPtrToString(e.MachineName),
		Os:            StringPtrToString(e.Os),
		CliVersion:    StringPtrToString(e.CliVersion),
		Connections:   LocalAgentConnectionSliceToProto(e.Connections),
		Status:        pb.LocalAgentStatus(e.Status),
		LastSeenAt:    TimePtrToTimestamppb(e.LastSeenAt),
		RevokedAt:     TimePtrToTimestamppb(e.RevokedAt),
		CreatedAt:     timestamppb.New(e.CreatedAt),
		UpdatedAt:     timestamppb.New(e.UpdatedAt),
		CreatedByUuid: e.CreatedByUUID.String(),
		UpdatedByUuid: e.UpdatedByUUID.String(),
	}
}

func LocalAgentSliceToProto(es []main_entity.LocalAgent) []*pb.LocalAgent {
	res := []*pb.LocalAgent{}
	for _, e := range es {
		res = append(res, LocalAgentToProto(e))
	}
	return res
}

func LocalAgentFromProto(m *pb.LocalAgent) main_entity.LocalAgent {
	if m == nil {
		return main_entity.LocalAgent{}
	}
	return main_entity.LocalAgent{
		UUID:          uuid.FromStringOrNil(m.GetUuid()),
		UserUUID:      uuid.FromStringOrNil(m.GetUserUuid()),
		TokenHash:     &m.TokenHash,
		MachineName:   &m.MachineName,
		Os:            &m.Os,
		CliVersion:    &m.CliVersion,
		Connections:   LocalAgentConnectionSliceFromProto(m.GetConnections()),
		Status:        main_entity.Status(m.GetStatus()),
		LastSeenAt:    TimestamppbToTimePtr(m.GetLastSeenAt()),
		RevokedAt:     TimestamppbToTimePtr(m.GetRevokedAt()),
		CreatedAt:     m.GetCreatedAt().AsTime(),
		UpdatedAt:     m.GetUpdatedAt().AsTime(),
		CreatedByUUID: uuid.FromStringOrNil(m.GetCreatedByUuid()),
		UpdatedByUUID: uuid.FromStringOrNil(m.GetUpdatedByUuid()),
	}
}

func LocalAgentSliceFromProto(es []*pb.LocalAgent) []main_entity.LocalAgent {
	if es == nil {
		return []main_entity.LocalAgent{}
	}
	res := []main_entity.LocalAgent{}
	for _, e := range es {
		res = append(res, LocalAgentFromProto(e))
	}
	return res
}
