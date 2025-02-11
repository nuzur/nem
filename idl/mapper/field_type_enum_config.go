package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/field_type_enum_config"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"
)

func FieldTypeEnumConfigToProto(e main_entity.FieldTypeEnumConfig) *pb.FieldTypeEnumConfig {
	return &pb.FieldTypeEnumConfig{
		EnumUuid:      e.EnumUUID.String(),
		AllowMultiple: e.AllowMultiple,
	}
}

func FieldTypeEnumConfigSliceToProto(es []main_entity.FieldTypeEnumConfig) []*pb.FieldTypeEnumConfig {
	res := []*pb.FieldTypeEnumConfig{}
	for _, e := range es {
		res = append(res, FieldTypeEnumConfigToProto(e))
	}
	return res
}

func FieldTypeEnumConfigFromProto(m *pb.FieldTypeEnumConfig) main_entity.FieldTypeEnumConfig {
	if m == nil {
		return main_entity.FieldTypeEnumConfig{}
	}
	return main_entity.FieldTypeEnumConfig{
		EnumUUID:      uuid.FromStringOrNil(m.GetEnumUuid()),
		AllowMultiple: m.GetAllowMultiple(),
	}
}

func FieldTypeEnumConfigSliceFromProto(es []*pb.FieldTypeEnumConfig) []main_entity.FieldTypeEnumConfig {
	if es == nil {
		return []main_entity.FieldTypeEnumConfig{}
	}
	res := []main_entity.FieldTypeEnumConfig{}
	for _, e := range es {
		res = append(res, FieldTypeEnumConfigFromProto(e))
	}
	return res
}
