package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/change_request_metadata"
	pb "github.com/nuzur/nem/idl/gen"
)

func ChangeRequestMetadataToProto(e main_entity.ChangeRequestMetadata) *pb.ChangeRequestMetadata {
	return &pb.ChangeRequestMetadata{
		Data: ChangeRequestMetadataDataToProto(e.Data),
	}
}

func ChangeRequestMetadataSliceToProto(es []main_entity.ChangeRequestMetadata) []*pb.ChangeRequestMetadata {
	res := []*pb.ChangeRequestMetadata{}
	for _, e := range es {
		res = append(res, ChangeRequestMetadataToProto(e))
	}
	return res
}

func ChangeRequestMetadataFromProto(m *pb.ChangeRequestMetadata) main_entity.ChangeRequestMetadata {
	if m == nil {
		return main_entity.ChangeRequestMetadata{}
	}
	return main_entity.ChangeRequestMetadata{
		Data: ChangeRequestMetadataDataFromProto(m.GetData()),
	}
}

func ChangeRequestMetadataSliceFromProto(es []*pb.ChangeRequestMetadata) []main_entity.ChangeRequestMetadata {
	if es == nil {
		return []main_entity.ChangeRequestMetadata{}
	}
	res := []main_entity.ChangeRequestMetadata{}
	for _, e := range es {
		res = append(res, ChangeRequestMetadataFromProto(e))
	}
	return res
}
