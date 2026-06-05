package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/user_connection_local_config"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"
)

func UserConnectionLocalConfigToProto(e main_entity.UserConnectionLocalConfig) *pb.UserConnectionLocalConfig {
	return &pb.UserConnectionLocalConfig{
		LocalAgentUuid:           e.LocalAgentUUID.String(),
		LocalAgentConnectionUuid: e.LocalAgentConnectionUUID.String(),
	}
}

func UserConnectionLocalConfigSliceToProto(es []main_entity.UserConnectionLocalConfig) []*pb.UserConnectionLocalConfig {
	res := []*pb.UserConnectionLocalConfig{}
	for _, e := range es {
		res = append(res, UserConnectionLocalConfigToProto(e))
	}
	return res
}

func UserConnectionLocalConfigFromProto(m *pb.UserConnectionLocalConfig) main_entity.UserConnectionLocalConfig {
	if m == nil {
		return main_entity.UserConnectionLocalConfig{}
	}
	return main_entity.UserConnectionLocalConfig{
		LocalAgentUUID:           uuid.FromStringOrNil(m.GetLocalAgentUuid()),
		LocalAgentConnectionUUID: uuid.FromStringOrNil(m.GetLocalAgentConnectionUuid()),
	}
}

func UserConnectionLocalConfigSliceFromProto(es []*pb.UserConnectionLocalConfig) []main_entity.UserConnectionLocalConfig {
	if es == nil {
		return []main_entity.UserConnectionLocalConfig{}
	}
	res := []main_entity.UserConnectionLocalConfig{}
	for _, e := range es {
		res = append(res, UserConnectionLocalConfigFromProto(e))
	}
	return res
}
