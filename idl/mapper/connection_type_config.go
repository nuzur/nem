package mapper

import (
	main_entity "nem/core/entity/connection_type_config"
	pb "nem/idl/gen"
)

func ConnectionTypeConfigToProto(e main_entity.ConnectionTypeConfig) *pb.ConnectionTypeConfig {
	return &pb.ConnectionTypeConfig{
		TcpIp:    TcpIpConnectionTypeConfigToProto(e.TcpIp),
		TcpIpSsh: TcpIpSshConnectionTypeConfigToProto(e.TcpIpSsh),
	}
}

func ConnectionTypeConfigSliceToProto(es []main_entity.ConnectionTypeConfig) []*pb.ConnectionTypeConfig {
	res := []*pb.ConnectionTypeConfig{}
	for _, e := range es {
		res = append(res, ConnectionTypeConfigToProto(e))
	}
	return res
}

func ConnectionTypeConfigFromProto(m *pb.ConnectionTypeConfig) main_entity.ConnectionTypeConfig {
	if m == nil {
		return main_entity.ConnectionTypeConfig{}
	}
	return main_entity.ConnectionTypeConfig{
		TcpIp:    TcpIpConnectionTypeConfigFromProto(m.GetTcpIp()),
		TcpIpSsh: TcpIpSshConnectionTypeConfigFromProto(m.GetTcpIpSsh()),
	}
}

func ConnectionTypeConfigSliceFromProto(es []*pb.ConnectionTypeConfig) []main_entity.ConnectionTypeConfig {
	if es == nil {
		return []main_entity.ConnectionTypeConfig{}
	}
	res := []main_entity.ConnectionTypeConfig{}
	for _, e := range es {
		res = append(res, ConnectionTypeConfigFromProto(e))
	}
	return res
}
