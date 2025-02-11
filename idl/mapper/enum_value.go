package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/enum_value"
	pb "github.com/nuzur/nem/idl/gen"
)

func EnumValueToProto(e main_entity.EnumValue) *pb.EnumValue {
	return &pb.EnumValue{
		Identifier:   e.Identifier,
		Display:      e.Display,
		NumericValue: int64(e.NumericValue),
	}
}

func EnumValueSliceToProto(es []main_entity.EnumValue) []*pb.EnumValue {
	res := []*pb.EnumValue{}
	for _, e := range es {
		res = append(res, EnumValueToProto(e))
	}
	return res
}

func EnumValueFromProto(m *pb.EnumValue) main_entity.EnumValue {
	if m == nil {
		return main_entity.EnumValue{}
	}
	return main_entity.EnumValue{
		Identifier:   m.GetIdentifier(),
		Display:      m.GetDisplay(),
		NumericValue: int64(m.GetNumericValue()),
	}
}

func EnumValueSliceFromProto(es []*pb.EnumValue) []main_entity.EnumValue {
	if es == nil {
		return []main_entity.EnumValue{}
	}
	res := []main_entity.EnumValue{}
	for _, e := range es {
		res = append(res, EnumValueFromProto(e))
	}
	return res
}
