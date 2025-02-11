package mapper

import (
	main_entity "nem/core/entity/index_field"
	pb "nem/idl/gen"

	"github.com/gofrs/uuid"
)

func IndexFieldToProto(e main_entity.IndexField) *pb.IndexField {
	return &pb.IndexField{
		FieldUuid: e.FieldUUID.String(),
		Priority:  int64(e.Priority),
		Order:     pb.IndexFieldOrder(e.Order),
		Length:    int64(e.Length),
	}
}

func IndexFieldSliceToProto(es []main_entity.IndexField) []*pb.IndexField {
	res := []*pb.IndexField{}
	for _, e := range es {
		res = append(res, IndexFieldToProto(e))
	}
	return res
}

func IndexFieldFromProto(m *pb.IndexField) main_entity.IndexField {
	if m == nil {
		return main_entity.IndexField{}
	}
	return main_entity.IndexField{
		FieldUUID: uuid.FromStringOrNil(m.GetFieldUuid()),
		Priority:  int64(m.GetPriority()),
		Order:     main_entity.Order(m.GetOrder()),
		Length:    int64(m.GetLength()),
	}
}

func IndexFieldSliceFromProto(es []*pb.IndexField) []main_entity.IndexField {
	if es == nil {
		return []main_entity.IndexField{}
	}
	res := []main_entity.IndexField{}
	for _, e := range es {
		res = append(res, IndexFieldFromProto(e))
	}
	return res
}
