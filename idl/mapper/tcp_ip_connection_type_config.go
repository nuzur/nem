package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/tcp_ip_connection_type_config"
	pb "github.com/nuzur/nem/idl/gen"
)

func TcpIpConnectionTypeConfigToProto(e main_entity.TcpIpConnectionTypeConfig) *pb.TcpIpConnectionTypeConfig {
	return &pb.TcpIpConnectionTypeConfig{
		Hostname: e.Hostname,
		Port:     e.Port,
		Username: e.Username,
		Password: e.Password,
	}
}

func TcpIpConnectionTypeConfigSliceToProto(es []main_entity.TcpIpConnectionTypeConfig) []*pb.TcpIpConnectionTypeConfig {
	res := []*pb.TcpIpConnectionTypeConfig{}
	for _, e := range es {
		res = append(res, TcpIpConnectionTypeConfigToProto(e))
	}
	return res
}

func TcpIpConnectionTypeConfigFromProto(m *pb.TcpIpConnectionTypeConfig) main_entity.TcpIpConnectionTypeConfig {
	if m == nil {
		return main_entity.TcpIpConnectionTypeConfig{}
	}
	return main_entity.TcpIpConnectionTypeConfig{
		Hostname: m.GetHostname(),
		Port:     m.GetPort(),
		Username: m.GetUsername(),
		Password: m.GetPassword(),
	}
}

func TcpIpConnectionTypeConfigSliceFromProto(es []*pb.TcpIpConnectionTypeConfig) []main_entity.TcpIpConnectionTypeConfig {
	if es == nil {
		return []main_entity.TcpIpConnectionTypeConfig{}
	}
	res := []main_entity.TcpIpConnectionTypeConfig{}
	for _, e := range es {
		res = append(res, TcpIpConnectionTypeConfigFromProto(e))
	}
	return res
}
