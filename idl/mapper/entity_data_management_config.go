package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/entity_data_management_config"
	pb "github.com/nuzur/nem/idl/gen"
)

func EntityDataManagementConfigToProto(e main_entity.EntityDataManagementConfig) *pb.EntityDataManagementConfig {
	return &pb.EntityDataManagementConfig{
		ListDisplayFields: UUIDSliceToStringSlice(e.ListDisplayFields),
	}
}

func EntityDataManagementConfigSliceToProto(es []main_entity.EntityDataManagementConfig) []*pb.EntityDataManagementConfig {
	res := []*pb.EntityDataManagementConfig{}
	for _, e := range es {
		res = append(res, EntityDataManagementConfigToProto(e))
	}
	return res
}

func EntityDataManagementConfigFromProto(m *pb.EntityDataManagementConfig) main_entity.EntityDataManagementConfig {
	if m == nil {
		return main_entity.EntityDataManagementConfig{}
	}
	return main_entity.EntityDataManagementConfig{
		ListDisplayFields: StringSliceToUUIDSlice(m.GetListDisplayFields()),
	}
}

func EntityDataManagementConfigSliceFromProto(es []*pb.EntityDataManagementConfig) []main_entity.EntityDataManagementConfig {
	if es == nil {
		return []main_entity.EntityDataManagementConfig{}
	}
	res := []main_entity.EntityDataManagementConfig{}
	for _, e := range es {
		res = append(res, EntityDataManagementConfigFromProto(e))
	}
	return res
}
