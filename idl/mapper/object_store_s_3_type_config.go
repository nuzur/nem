package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/object_store_s3_type_config"
	pb "github.com/nuzur/nem/idl/gen"
)

func ObjectStoreS3TypeConfigToProto(e main_entity.ObjectStoreS3TypeConfig) *pb.ObjectStoreS3TypeConfig {
	return &pb.ObjectStoreS3TypeConfig{
		Region: e.Region,
		Key:    e.Key,
		Secret: e.Secret,
		Bucket: e.Bucket,
	}
}

func ObjectStoreS3TypeConfigSliceToProto(es []main_entity.ObjectStoreS3TypeConfig) []*pb.ObjectStoreS3TypeConfig {
	res := []*pb.ObjectStoreS3TypeConfig{}
	for _, e := range es {
		res = append(res, ObjectStoreS3TypeConfigToProto(e))
	}
	return res
}

func ObjectStoreS3TypeConfigFromProto(m *pb.ObjectStoreS3TypeConfig) main_entity.ObjectStoreS3TypeConfig {
	if m == nil {
		return main_entity.ObjectStoreS3TypeConfig{}
	}
	return main_entity.ObjectStoreS3TypeConfig{
		Region: m.GetRegion(),
		Key:    m.GetKey(),
		Secret: m.GetSecret(),
		Bucket: m.GetBucket(),
	}
}

func ObjectStoreS3TypeConfigSliceFromProto(es []*pb.ObjectStoreS3TypeConfig) []main_entity.ObjectStoreS3TypeConfig {
	if es == nil {
		return []main_entity.ObjectStoreS3TypeConfig{}
	}
	res := []main_entity.ObjectStoreS3TypeConfig{}
	for _, e := range es {
		res = append(res, ObjectStoreS3TypeConfigFromProto(e))
	}
	return res
}
