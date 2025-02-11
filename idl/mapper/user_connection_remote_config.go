package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/user_connection_remote_config"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"
)

func UserConnectionRemoteConfigToProto(e main_entity.UserConnectionRemoteConfig) *pb.UserConnectionRemoteConfig {
	return &pb.UserConnectionRemoteConfig{
		TeamUuid:       e.TeamUUID.String(),
		StoreUuid:      e.StoreUUID.String(),
		ConnectionUuid: e.ConnectionUUID.String(),
	}
}

func UserConnectionRemoteConfigSliceToProto(es []main_entity.UserConnectionRemoteConfig) []*pb.UserConnectionRemoteConfig {
	res := []*pb.UserConnectionRemoteConfig{}
	for _, e := range es {
		res = append(res, UserConnectionRemoteConfigToProto(e))
	}
	return res
}

func UserConnectionRemoteConfigFromProto(m *pb.UserConnectionRemoteConfig) main_entity.UserConnectionRemoteConfig {
	if m == nil {
		return main_entity.UserConnectionRemoteConfig{}
	}
	return main_entity.UserConnectionRemoteConfig{
		TeamUUID:       uuid.FromStringOrNil(m.GetTeamUuid()),
		StoreUUID:      uuid.FromStringOrNil(m.GetStoreUuid()),
		ConnectionUUID: uuid.FromStringOrNil(m.GetConnectionUuid()),
	}
}

func UserConnectionRemoteConfigSliceFromProto(es []*pb.UserConnectionRemoteConfig) []main_entity.UserConnectionRemoteConfig {
	if es == nil {
		return []main_entity.UserConnectionRemoteConfig{}
	}
	res := []main_entity.UserConnectionRemoteConfig{}
	for _, e := range es {
		res = append(res, UserConnectionRemoteConfigFromProto(e))
	}
	return res
}
