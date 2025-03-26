package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/change_request_data_change"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"
)

func ChangeRequestDataChangeToProto(e main_entity.ChangeRequestDataChange) *pb.ChangeRequestDataChange {
	return &pb.ChangeRequestDataChange{
		Uuid:                 e.UUID.String(),
		DataChangeType:       pb.ChangeRequestDataChangeDataChangeType(e.DataChangeType),
		DataChangeTypeConfig: DataChangeTypeConfigToProto(e.DataChangeTypeConfig),
		Order:                int64(e.Order),
	}
}

func ChangeRequestDataChangeSliceToProto(es []main_entity.ChangeRequestDataChange) []*pb.ChangeRequestDataChange {
	res := []*pb.ChangeRequestDataChange{}
	for _, e := range es {
		res = append(res, ChangeRequestDataChangeToProto(e))
	}
	return res
}

func ChangeRequestDataChangeFromProto(m *pb.ChangeRequestDataChange) main_entity.ChangeRequestDataChange {
	if m == nil {
		return main_entity.ChangeRequestDataChange{}
	}
	return main_entity.ChangeRequestDataChange{
		UUID:                 uuid.FromStringOrNil(m.GetUuid()),
		DataChangeType:       main_entity.DataChangeType(m.GetDataChangeType()),
		DataChangeTypeConfig: DataChangeTypeConfigFromProto(m.GetDataChangeTypeConfig()),
		Order:                int64(m.GetOrder()),
	}
}

func ChangeRequestDataChangeSliceFromProto(es []*pb.ChangeRequestDataChange) []main_entity.ChangeRequestDataChange {
	if es == nil {
		return []main_entity.ChangeRequestDataChange{}
	}
	res := []main_entity.ChangeRequestDataChange{}
	for _, e := range es {
		res = append(res, ChangeRequestDataChangeFromProto(e))
	}
	return res
}
