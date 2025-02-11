package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/user_connection_local_config"
	pb "github.com/nuzur/nem/idl/gen"
)

func UserConnectionLocalConfigToProto(e main_entity.UserConnectionLocalConfig) *pb.UserConnectionLocalConfig {
	return &pb.UserConnectionLocalConfig{
		IpAddress: e.IpAddress,
		DbType:    pb.UserConnectionLocalConfigDbType(e.DbType),
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
		IpAddress: m.GetIpAddress(),
		DbType:    main_entity.DbType(m.GetDbType()),
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
