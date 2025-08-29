package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/change_request_metadata_data"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"
)

func ChangeRequestMetadataDataToProto(e main_entity.ChangeRequestMetadataData) *pb.ChangeRequestMetadataData {
	return &pb.ChangeRequestMetadataData{
		StoreUuid:      e.StoreUUID.String(),
		ConnectionUuid: e.ConnectionUUID.String(),
		Schema:         StringPtrToString(e.Schema),
	}
}

func ChangeRequestMetadataDataSliceToProto(es []main_entity.ChangeRequestMetadataData) []*pb.ChangeRequestMetadataData {
	res := []*pb.ChangeRequestMetadataData{}
	for _, e := range es {
		res = append(res, ChangeRequestMetadataDataToProto(e))
	}
	return res
}

func ChangeRequestMetadataDataFromProto(m *pb.ChangeRequestMetadataData) main_entity.ChangeRequestMetadataData {
	if m == nil {
		return main_entity.ChangeRequestMetadataData{}
	}
	return main_entity.ChangeRequestMetadataData{
		StoreUUID:      uuid.FromStringOrNil(m.GetStoreUuid()),
		ConnectionUUID: uuid.FromStringOrNil(m.GetConnectionUuid()),
		Schema:         &m.Schema,
	}
}

func ChangeRequestMetadataDataSliceFromProto(es []*pb.ChangeRequestMetadataData) []main_entity.ChangeRequestMetadataData {
	if es == nil {
		return []main_entity.ChangeRequestMetadataData{}
	}
	res := []main_entity.ChangeRequestMetadataData{}
	for _, e := range es {
		res = append(res, ChangeRequestMetadataDataFromProto(e))
	}
	return res
}
