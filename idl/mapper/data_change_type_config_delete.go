package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/data_change_type_config_delete"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func DataChangeTypeConfigDeleteToProto(e main_entity.DataChangeTypeConfigDelete) *pb.DataChangeTypeConfigDelete {
	return &pb.DataChangeTypeConfigDelete{
		EntityUuid: e.EntityUUID.String(),
		Keys:       DataChangeFieldValueSliceToProto(e.Keys),
		CreatedAt:  timestamppb.New(e.CreatedAt),
	}
}

func DataChangeTypeConfigDeleteSliceToProto(es []main_entity.DataChangeTypeConfigDelete) []*pb.DataChangeTypeConfigDelete {
	res := []*pb.DataChangeTypeConfigDelete{}
	for _, e := range es {
		res = append(res, DataChangeTypeConfigDeleteToProto(e))
	}
	return res
}

func DataChangeTypeConfigDeleteFromProto(m *pb.DataChangeTypeConfigDelete) main_entity.DataChangeTypeConfigDelete {
	if m == nil {
		return main_entity.DataChangeTypeConfigDelete{}
	}
	return main_entity.DataChangeTypeConfigDelete{
		EntityUUID: uuid.FromStringOrNil(m.GetEntityUuid()),
		Keys:       DataChangeFieldValueSliceFromProto(m.GetKeys()),
		CreatedAt:  m.GetCreatedAt().AsTime(),
	}
}

func DataChangeTypeConfigDeleteSliceFromProto(es []*pb.DataChangeTypeConfigDelete) []main_entity.DataChangeTypeConfigDelete {
	if es == nil {
		return []main_entity.DataChangeTypeConfigDelete{}
	}
	res := []main_entity.DataChangeTypeConfigDelete{}
	for _, e := range es {
		res = append(res, DataChangeTypeConfigDeleteFromProto(e))
	}
	return res
}
