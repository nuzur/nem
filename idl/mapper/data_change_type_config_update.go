package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/data_change_type_config_update"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func DataChangeTypeConfigUpdateToProto(e main_entity.DataChangeTypeConfigUpdate) *pb.DataChangeTypeConfigUpdate {
	return &pb.DataChangeTypeConfigUpdate{
		EntityUuid:   e.EntityUUID.String(),
		FieldUuid:    e.FieldUUID.String(),
		CurrentValue: e.CurrentValue,
		NewValue:     e.NewValue,
		Keys:         DataChangeFieldValueSliceToProto(e.Keys),
		CreatedAt:    timestamppb.New(e.CreatedAt),
	}
}

func DataChangeTypeConfigUpdateSliceToProto(es []main_entity.DataChangeTypeConfigUpdate) []*pb.DataChangeTypeConfigUpdate {
	res := []*pb.DataChangeTypeConfigUpdate{}
	for _, e := range es {
		res = append(res, DataChangeTypeConfigUpdateToProto(e))
	}
	return res
}

func DataChangeTypeConfigUpdateFromProto(m *pb.DataChangeTypeConfigUpdate) main_entity.DataChangeTypeConfigUpdate {
	if m == nil {
		return main_entity.DataChangeTypeConfigUpdate{}
	}
	return main_entity.DataChangeTypeConfigUpdate{
		EntityUUID:   uuid.FromStringOrNil(m.GetEntityUuid()),
		FieldUUID:    uuid.FromStringOrNil(m.GetFieldUuid()),
		CurrentValue: m.GetCurrentValue(),
		NewValue:     m.GetNewValue(),
		Keys:         DataChangeFieldValueSliceFromProto(m.GetKeys()),
		CreatedAt:    m.GetCreatedAt().AsTime(),
	}
}

func DataChangeTypeConfigUpdateSliceFromProto(es []*pb.DataChangeTypeConfigUpdate) []main_entity.DataChangeTypeConfigUpdate {
	if es == nil {
		return []main_entity.DataChangeTypeConfigUpdate{}
	}
	res := []main_entity.DataChangeTypeConfigUpdate{}
	for _, e := range es {
		res = append(res, DataChangeTypeConfigUpdateFromProto(e))
	}
	return res
}
