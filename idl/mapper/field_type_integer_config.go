package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/field_type_integer_config"
	pb "github.com/nuzur/nem/idl/gen"
)

func FieldTypeIntegerConfigToProto(e main_entity.FieldTypeIntegerConfig) *pb.FieldTypeIntegerConfig {
	return &pb.FieldTypeIntegerConfig{
		Size:              pb.FieldTypeIntegerConfigSize(e.Size),
		MinValue:          int64(e.MinValue),
		MinValueInclusive: e.MinValueInclusive,
		MaxValue:          int64(e.MaxValue),
		MaxValueInclusive: e.MaxValueInclusive,
		AllowNegatives:    e.AllowNegatives,
		EnableLimits:      e.EnableLimits,
	}
}

func FieldTypeIntegerConfigSliceToProto(es []main_entity.FieldTypeIntegerConfig) []*pb.FieldTypeIntegerConfig {
	res := []*pb.FieldTypeIntegerConfig{}
	for _, e := range es {
		res = append(res, FieldTypeIntegerConfigToProto(e))
	}
	return res
}

func FieldTypeIntegerConfigFromProto(m *pb.FieldTypeIntegerConfig) main_entity.FieldTypeIntegerConfig {
	if m == nil {
		return main_entity.FieldTypeIntegerConfig{}
	}
	return main_entity.FieldTypeIntegerConfig{
		Size:              main_entity.Size(m.GetSize()),
		MinValue:          int64(m.GetMinValue()),
		MinValueInclusive: m.GetMinValueInclusive(),
		MaxValue:          int64(m.GetMaxValue()),
		MaxValueInclusive: m.GetMaxValueInclusive(),
		AllowNegatives:    m.GetAllowNegatives(),
		EnableLimits:      m.GetEnableLimits(),
	}
}

func FieldTypeIntegerConfigSliceFromProto(es []*pb.FieldTypeIntegerConfig) []main_entity.FieldTypeIntegerConfig {
	if es == nil {
		return []main_entity.FieldTypeIntegerConfig{}
	}
	res := []main_entity.FieldTypeIntegerConfig{}
	for _, e := range es {
		res = append(res, FieldTypeIntegerConfigFromProto(e))
	}
	return res
}
