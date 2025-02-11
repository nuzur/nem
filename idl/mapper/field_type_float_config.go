package mapper

import (
	main_entity "nem/core/entity/field_type_float_config"
	pb "nem/idl/gen"
)

func FieldTypeFloatConfigToProto(e main_entity.FieldTypeFloatConfig) *pb.FieldTypeFloatConfig {
	return &pb.FieldTypeFloatConfig{
		MinValue:          e.MinValue,
		MinValueInclusive: e.MinValueInclusive,
		MaxValue:          e.MaxValue,
		MaxValueInclusive: e.MaxValueInclusive,
		AllowNegatives:    e.AllowNegatives,
		NumberOfDecimals:  int64(e.NumberOfDecimals),
		Separator:         pb.FieldTypeFloatConfigSeparator(e.Separator),
		EnableLimits:      e.EnableLimits,
	}
}

func FieldTypeFloatConfigSliceToProto(es []main_entity.FieldTypeFloatConfig) []*pb.FieldTypeFloatConfig {
	res := []*pb.FieldTypeFloatConfig{}
	for _, e := range es {
		res = append(res, FieldTypeFloatConfigToProto(e))
	}
	return res
}

func FieldTypeFloatConfigFromProto(m *pb.FieldTypeFloatConfig) main_entity.FieldTypeFloatConfig {
	if m == nil {
		return main_entity.FieldTypeFloatConfig{}
	}
	return main_entity.FieldTypeFloatConfig{
		MinValue:          m.MinValue,
		MinValueInclusive: m.GetMinValueInclusive(),
		MaxValue:          m.MaxValue,
		MaxValueInclusive: m.GetMaxValueInclusive(),
		AllowNegatives:    m.GetAllowNegatives(),
		NumberOfDecimals:  int64(m.GetNumberOfDecimals()),
		Separator:         main_entity.Separator(m.GetSeparator()),
		EnableLimits:      m.GetEnableLimits(),
	}
}

func FieldTypeFloatConfigSliceFromProto(es []*pb.FieldTypeFloatConfig) []main_entity.FieldTypeFloatConfig {
	if es == nil {
		return []main_entity.FieldTypeFloatConfig{}
	}
	res := []main_entity.FieldTypeFloatConfig{}
	for _, e := range es {
		res = append(res, FieldTypeFloatConfigFromProto(e))
	}
	return res
}
