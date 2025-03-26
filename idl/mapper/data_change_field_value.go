package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/data_change_field_value"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"
)

func DataChangeFieldValueToProto(e main_entity.DataChangeFieldValue) *pb.DataChangeFieldValue {
	return &pb.DataChangeFieldValue{
		FieldUuid: e.FieldUUID.String(),
		Value:     e.Value,
	}
}

func DataChangeFieldValueSliceToProto(es []main_entity.DataChangeFieldValue) []*pb.DataChangeFieldValue {
	res := []*pb.DataChangeFieldValue{}
	for _, e := range es {
		res = append(res, DataChangeFieldValueToProto(e))
	}
	return res
}

func DataChangeFieldValueFromProto(m *pb.DataChangeFieldValue) main_entity.DataChangeFieldValue {
	if m == nil {
		return main_entity.DataChangeFieldValue{}
	}
	return main_entity.DataChangeFieldValue{
		FieldUUID: uuid.FromStringOrNil(m.GetFieldUuid()),
		Value:     m.GetValue(),
	}
}

func DataChangeFieldValueSliceFromProto(es []*pb.DataChangeFieldValue) []main_entity.DataChangeFieldValue {
	if es == nil {
		return []main_entity.DataChangeFieldValue{}
	}
	res := []main_entity.DataChangeFieldValue{}
	for _, e := range es {
		res = append(res, DataChangeFieldValueFromProto(e))
	}
	return res
}
