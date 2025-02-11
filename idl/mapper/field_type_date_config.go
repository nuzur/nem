package mapper

import (
	main_entity "nem/core/entity/field_type_date_config"
	pb "nem/idl/gen"
)

func FieldTypeDateConfigToProto(e main_entity.FieldTypeDateConfig) *pb.FieldTypeDateConfig {
	return &pb.FieldTypeDateConfig{
		EnforceFuture:   e.EnforceFuture,
		EnforcePast:     e.EnforcePast,
		DisplayTimezone: e.DisplayTimezone,
		StorageTimezone: e.StorageTimezone,
	}
}

func FieldTypeDateConfigSliceToProto(es []main_entity.FieldTypeDateConfig) []*pb.FieldTypeDateConfig {
	res := []*pb.FieldTypeDateConfig{}
	for _, e := range es {
		res = append(res, FieldTypeDateConfigToProto(e))
	}
	return res
}

func FieldTypeDateConfigFromProto(m *pb.FieldTypeDateConfig) main_entity.FieldTypeDateConfig {
	if m == nil {
		return main_entity.FieldTypeDateConfig{}
	}
	return main_entity.FieldTypeDateConfig{
		EnforceFuture:   m.GetEnforceFuture(),
		EnforcePast:     m.GetEnforcePast(),
		DisplayTimezone: m.GetDisplayTimezone(),
		StorageTimezone: m.GetStorageTimezone(),
	}
}

func FieldTypeDateConfigSliceFromProto(es []*pb.FieldTypeDateConfig) []main_entity.FieldTypeDateConfig {
	if es == nil {
		return []main_entity.FieldTypeDateConfig{}
	}
	res := []main_entity.FieldTypeDateConfig{}
	for _, e := range es {
		res = append(res, FieldTypeDateConfigFromProto(e))
	}
	return res
}
