package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/field_type_decimal_config"
	pb "github.com/nuzur/nem/idl/gen"
)

func FieldTypeDecimalConfigToProto(e main_entity.FieldTypeDecimalConfig) *pb.FieldTypeDecimalConfig {
	return &pb.FieldTypeDecimalConfig{
		MinValue:          e.MinValue,
		MinValueInclusive: e.MinValueInclusive,
		MaxValue:          e.MaxValue,
		MaxValueInclusive: e.MaxValueInclusive,
		AllowNegatives:    e.AllowNegatives,
		NumberOfDecimals:  int64(e.NumberOfDecimals),
		Separator:         pb.FieldTypeDecimalConfigSeparator(e.Separator),
		IsCurrency:        e.IsCurrency,
		CurrencyCode:      e.CurrencyCode,
		EnableLimits:      e.EnableLimits,
	}
}

func FieldTypeDecimalConfigSliceToProto(es []main_entity.FieldTypeDecimalConfig) []*pb.FieldTypeDecimalConfig {
	res := []*pb.FieldTypeDecimalConfig{}
	for _, e := range es {
		res = append(res, FieldTypeDecimalConfigToProto(e))
	}
	return res
}

func FieldTypeDecimalConfigFromProto(m *pb.FieldTypeDecimalConfig) main_entity.FieldTypeDecimalConfig {
	if m == nil {
		return main_entity.FieldTypeDecimalConfig{}
	}
	return main_entity.FieldTypeDecimalConfig{
		MinValue:          m.MinValue,
		MinValueInclusive: m.GetMinValueInclusive(),
		MaxValue:          m.MaxValue,
		MaxValueInclusive: m.GetMaxValueInclusive(),
		AllowNegatives:    m.GetAllowNegatives(),
		NumberOfDecimals:  int64(m.GetNumberOfDecimals()),
		Separator:         main_entity.Separator(m.GetSeparator()),
		IsCurrency:        m.GetIsCurrency(),
		CurrencyCode:      m.GetCurrencyCode(),
		EnableLimits:      m.GetEnableLimits(),
	}
}

func FieldTypeDecimalConfigSliceFromProto(es []*pb.FieldTypeDecimalConfig) []main_entity.FieldTypeDecimalConfig {
	if es == nil {
		return []main_entity.FieldTypeDecimalConfig{}
	}
	res := []main_entity.FieldTypeDecimalConfig{}
	for _, e := range es {
		res = append(res, FieldTypeDecimalConfigFromProto(e))
	}
	return res
}
