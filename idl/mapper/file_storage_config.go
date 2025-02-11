package mapper

import (
	main_entity "nem/core/entity/file_storage_config"
	pb "nem/idl/gen"
)

func FileStorageConfigToProto(e main_entity.FileStorageConfig) *pb.FileStorageConfig {
	return &pb.FileStorageConfig{
		ObjectStore: FileObjectStorageConfigToProto(e.ObjectStore),
	}
}

func FileStorageConfigSliceToProto(es []main_entity.FileStorageConfig) []*pb.FileStorageConfig {
	res := []*pb.FileStorageConfig{}
	for _, e := range es {
		res = append(res, FileStorageConfigToProto(e))
	}
	return res
}

func FileStorageConfigFromProto(m *pb.FileStorageConfig) main_entity.FileStorageConfig {
	if m == nil {
		return main_entity.FileStorageConfig{}
	}
	return main_entity.FileStorageConfig{
		ObjectStore: FileObjectStorageConfigFromProto(m.GetObjectStore()),
	}
}

func FileStorageConfigSliceFromProto(es []*pb.FileStorageConfig) []main_entity.FileStorageConfig {
	if es == nil {
		return []main_entity.FileStorageConfig{}
	}
	res := []main_entity.FileStorageConfig{}
	for _, e := range es {
		res = append(res, FileStorageConfigFromProto(e))
	}
	return res
}
