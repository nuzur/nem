package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/data_change_type_config_update_field"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"
)

func DataChangeTypeConfigUpdateFieldToProto(e main_entity.DataChangeTypeConfigUpdateField) *pb.DataChangeTypeConfigUpdateField {
	return &pb.DataChangeTypeConfigUpdateField{
		FieldUuid:    e.FieldUUID.String(),
		CurrentValue: e.CurrentValue,
		NewValue:     e.NewValue,
	}
}

func DataChangeTypeConfigUpdateFieldSliceToProto(es []main_entity.DataChangeTypeConfigUpdateField) []*pb.DataChangeTypeConfigUpdateField {
	res := []*pb.DataChangeTypeConfigUpdateField{}
	for _, e := range es {
		res = append(res, DataChangeTypeConfigUpdateFieldToProto(e))
	}
	return res
}

func DataChangeTypeConfigUpdateFieldFromProto(m *pb.DataChangeTypeConfigUpdateField) main_entity.DataChangeTypeConfigUpdateField {
	if m == nil {
		return main_entity.DataChangeTypeConfigUpdateField{}
	}
	return main_entity.DataChangeTypeConfigUpdateField{
		FieldUUID:    uuid.FromStringOrNil(m.GetFieldUuid()),
		CurrentValue: m.GetCurrentValue(),
		NewValue:     m.GetNewValue(),
	}
}

func DataChangeTypeConfigUpdateFieldSliceFromProto(es []*pb.DataChangeTypeConfigUpdateField) []main_entity.DataChangeTypeConfigUpdateField {
	if es == nil {
		return []main_entity.DataChangeTypeConfigUpdateField{}
	}
	res := []main_entity.DataChangeTypeConfigUpdateField{}
	for _, e := range es {
		res = append(res, DataChangeTypeConfigUpdateFieldFromProto(e))
	}
	return res
}
