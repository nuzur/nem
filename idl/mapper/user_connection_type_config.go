package mapper

import (
	main_entity "nem/core/entity/user_connection_type_config"
	pb "nem/idl/gen"
)

func UserConnectionTypeConfigToProto(e main_entity.UserConnectionTypeConfig) *pb.UserConnectionTypeConfig {
	return &pb.UserConnectionTypeConfig{
		Local:  UserConnectionLocalConfigToProto(e.Local),
		Remote: UserConnectionRemoteConfigToProto(e.Remote),
	}
}

func UserConnectionTypeConfigSliceToProto(es []main_entity.UserConnectionTypeConfig) []*pb.UserConnectionTypeConfig {
	res := []*pb.UserConnectionTypeConfig{}
	for _, e := range es {
		res = append(res, UserConnectionTypeConfigToProto(e))
	}
	return res
}

func UserConnectionTypeConfigFromProto(m *pb.UserConnectionTypeConfig) main_entity.UserConnectionTypeConfig {
	if m == nil {
		return main_entity.UserConnectionTypeConfig{}
	}
	return main_entity.UserConnectionTypeConfig{
		Local:  UserConnectionLocalConfigFromProto(m.GetLocal()),
		Remote: UserConnectionRemoteConfigFromProto(m.GetRemote()),
	}
}

func UserConnectionTypeConfigSliceFromProto(es []*pb.UserConnectionTypeConfig) []main_entity.UserConnectionTypeConfig {
	if es == nil {
		return []main_entity.UserConnectionTypeConfig{}
	}
	res := []main_entity.UserConnectionTypeConfig{}
	for _, e := range es {
		res = append(res, UserConnectionTypeConfigFromProto(e))
	}
	return res
}
