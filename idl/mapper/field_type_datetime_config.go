package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/field_type_datetime_config"
	pb "github.com/nuzur/nem/idl/gen"
)

func FieldTypeDatetimeConfigToProto(e main_entity.FieldTypeDatetimeConfig) *pb.FieldTypeDatetimeConfig {
	return &pb.FieldTypeDatetimeConfig{
		EnforceFuture:   e.EnforceFuture,
		EnforcePast:     e.EnforcePast,
		DisplayTimezone: StringPtrToString(e.DisplayTimezone),
		StorageTimezone: StringPtrToString(e.StorageTimezone),
	}
}

func FieldTypeDatetimeConfigSliceToProto(es []main_entity.FieldTypeDatetimeConfig) []*pb.FieldTypeDatetimeConfig {
	res := []*pb.FieldTypeDatetimeConfig{}
	for _, e := range es {
		res = append(res, FieldTypeDatetimeConfigToProto(e))
	}
	return res
}

func FieldTypeDatetimeConfigFromProto(m *pb.FieldTypeDatetimeConfig) main_entity.FieldTypeDatetimeConfig {
	if m == nil {
		return main_entity.FieldTypeDatetimeConfig{}
	}
	return main_entity.FieldTypeDatetimeConfig{
		EnforceFuture:   m.GetEnforceFuture(),
		EnforcePast:     m.GetEnforcePast(),
		DisplayTimezone: &m.DisplayTimezone,
		StorageTimezone: &m.StorageTimezone,
	}
}

func FieldTypeDatetimeConfigSliceFromProto(es []*pb.FieldTypeDatetimeConfig) []main_entity.FieldTypeDatetimeConfig {
	if es == nil {
		return []main_entity.FieldTypeDatetimeConfig{}
	}
	res := []main_entity.FieldTypeDatetimeConfig{}
	for _, e := range es {
		res = append(res, FieldTypeDatetimeConfigFromProto(e))
	}
	return res
}
