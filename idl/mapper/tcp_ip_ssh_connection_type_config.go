package mapper

import (
	main_entity "nem/core/entity/tcp_ip_ssh_connection_type_config"
	pb "nem/idl/gen"
)

func TcpIpSshConnectionTypeConfigToProto(e main_entity.TcpIpSshConnectionTypeConfig) *pb.TcpIpSshConnectionTypeConfig {
	return &pb.TcpIpSshConnectionTypeConfig{
		SshHostname:   e.SshHostname,
		SshUsername:   e.SshUsername,
		SshPassword:   e.SshPassword,
		SshKeyFile:    e.SshKeyFile,
		UseSshKeyFile: e.UseSshKeyFile,
		Hostname:      e.Hostname,
		Port:          e.Port,
		Username:      e.Username,
		Password:      e.Password,
	}
}

func TcpIpSshConnectionTypeConfigSliceToProto(es []main_entity.TcpIpSshConnectionTypeConfig) []*pb.TcpIpSshConnectionTypeConfig {
	res := []*pb.TcpIpSshConnectionTypeConfig{}
	for _, e := range es {
		res = append(res, TcpIpSshConnectionTypeConfigToProto(e))
	}
	return res
}

func TcpIpSshConnectionTypeConfigFromProto(m *pb.TcpIpSshConnectionTypeConfig) main_entity.TcpIpSshConnectionTypeConfig {
	if m == nil {
		return main_entity.TcpIpSshConnectionTypeConfig{}
	}
	return main_entity.TcpIpSshConnectionTypeConfig{
		SshHostname:   m.GetSshHostname(),
		SshUsername:   m.GetSshUsername(),
		SshPassword:   m.GetSshPassword(),
		SshKeyFile:    m.GetSshKeyFile(),
		UseSshKeyFile: m.GetUseSshKeyFile(),
		Hostname:      m.GetHostname(),
		Port:          m.GetPort(),
		Username:      m.GetUsername(),
		Password:      m.GetPassword(),
	}
}

func TcpIpSshConnectionTypeConfigSliceFromProto(es []*pb.TcpIpSshConnectionTypeConfig) []main_entity.TcpIpSshConnectionTypeConfig {
	if es == nil {
		return []main_entity.TcpIpSshConnectionTypeConfig{}
	}
	res := []main_entity.TcpIpSshConnectionTypeConfig{}
	for _, e := range es {
		res = append(res, TcpIpSshConnectionTypeConfigFromProto(e))
	}
	return res
}
