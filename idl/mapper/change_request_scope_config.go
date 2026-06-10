package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/change_request_scope_config"
	pb "github.com/nuzur/nem/idl/gen"
)

func ChangeRequestScopeConfigToProto(e main_entity.ChangeRequestScopeConfig) *pb.ChangeRequestScopeConfig {
	return &pb.ChangeRequestScopeConfig{
		Local:  ChangeRequestScopeConfigLocalToProto(e.Local),
		Remote: ChangeRequestScopeConfigRemoteToProto(e.Remote),
	}
}

func ChangeRequestScopeConfigSliceToProto(es []main_entity.ChangeRequestScopeConfig) []*pb.ChangeRequestScopeConfig {
	res := []*pb.ChangeRequestScopeConfig{}
	for _, e := range es {
		res = append(res, ChangeRequestScopeConfigToProto(e))
	}
	return res
}

func ChangeRequestScopeConfigFromProto(m *pb.ChangeRequestScopeConfig) main_entity.ChangeRequestScopeConfig {
	if m == nil {
		return main_entity.ChangeRequestScopeConfig{}
	}
	return main_entity.ChangeRequestScopeConfig{
		Local:  ChangeRequestScopeConfigLocalFromProto(m.GetLocal()),
		Remote: ChangeRequestScopeConfigRemoteFromProto(m.GetRemote()),
	}
}

func ChangeRequestScopeConfigSliceFromProto(es []*pb.ChangeRequestScopeConfig) []main_entity.ChangeRequestScopeConfig {
	if es == nil {
		return []main_entity.ChangeRequestScopeConfig{}
	}
	res := []main_entity.ChangeRequestScopeConfig{}
	for _, e := range es {
		res = append(res, ChangeRequestScopeConfigFromProto(e))
	}
	return res
}
