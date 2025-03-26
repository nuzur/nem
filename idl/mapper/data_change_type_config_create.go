package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/data_change_type_config_create"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func DataChangeTypeConfigCreateToProto(e main_entity.DataChangeTypeConfigCreate) *pb.DataChangeTypeConfigCreate {
	return &pb.DataChangeTypeConfigCreate{
		EntityUuid: e.EntityUUID.String(),
		Keys:       DataChangeFieldValueSliceToProto(e.Keys),
		CreatedAt:  timestamppb.New(e.CreatedAt),
	}
}

func DataChangeTypeConfigCreateSliceToProto(es []main_entity.DataChangeTypeConfigCreate) []*pb.DataChangeTypeConfigCreate {
	res := []*pb.DataChangeTypeConfigCreate{}
	for _, e := range es {
		res = append(res, DataChangeTypeConfigCreateToProto(e))
	}
	return res
}

func DataChangeTypeConfigCreateFromProto(m *pb.DataChangeTypeConfigCreate) main_entity.DataChangeTypeConfigCreate {
	if m == nil {
		return main_entity.DataChangeTypeConfigCreate{}
	}
	return main_entity.DataChangeTypeConfigCreate{
		EntityUUID: uuid.FromStringOrNil(m.GetEntityUuid()),
		Keys:       DataChangeFieldValueSliceFromProto(m.GetKeys()),
		CreatedAt:  m.GetCreatedAt().AsTime(),
	}
}

func DataChangeTypeConfigCreateSliceFromProto(es []*pb.DataChangeTypeConfigCreate) []main_entity.DataChangeTypeConfigCreate {
	if es == nil {
		return []main_entity.DataChangeTypeConfigCreate{}
	}
	res := []main_entity.DataChangeTypeConfigCreate{}
	for _, e := range es {
		res = append(res, DataChangeTypeConfigCreateFromProto(e))
	}
	return res
}
