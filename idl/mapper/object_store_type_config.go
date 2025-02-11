package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/object_store_type_config"
	pb "github.com/nuzur/nem/idl/gen"
)

func ObjectStoreTypeConfigToProto(e main_entity.ObjectStoreTypeConfig) *pb.ObjectStoreTypeConfig {
	return &pb.ObjectStoreTypeConfig{
		S3: ObjectStoreS3TypeConfigToProto(e.S3),
	}
}

func ObjectStoreTypeConfigSliceToProto(es []main_entity.ObjectStoreTypeConfig) []*pb.ObjectStoreTypeConfig {
	res := []*pb.ObjectStoreTypeConfig{}
	for _, e := range es {
		res = append(res, ObjectStoreTypeConfigToProto(e))
	}
	return res
}

func ObjectStoreTypeConfigFromProto(m *pb.ObjectStoreTypeConfig) main_entity.ObjectStoreTypeConfig {
	if m == nil {
		return main_entity.ObjectStoreTypeConfig{}
	}
	return main_entity.ObjectStoreTypeConfig{
		S3: ObjectStoreS3TypeConfigFromProto(m.GetS3()),
	}
}

func ObjectStoreTypeConfigSliceFromProto(es []*pb.ObjectStoreTypeConfig) []main_entity.ObjectStoreTypeConfig {
	if es == nil {
		return []main_entity.ObjectStoreTypeConfig{}
	}
	res := []main_entity.ObjectStoreTypeConfig{}
	for _, e := range es {
		res = append(res, ObjectStoreTypeConfigFromProto(e))
	}
	return res
}
