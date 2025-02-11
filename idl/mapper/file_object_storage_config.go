package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/file_object_storage_config"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"
)

func FileObjectStorageConfigToProto(e main_entity.FileObjectStorageConfig) *pb.FileObjectStorageConfig {
	return &pb.FileObjectStorageConfig{
		ObjectStoreUuid: e.ObjectStoreUUID.String(),
	}
}

func FileObjectStorageConfigSliceToProto(es []main_entity.FileObjectStorageConfig) []*pb.FileObjectStorageConfig {
	res := []*pb.FileObjectStorageConfig{}
	for _, e := range es {
		res = append(res, FileObjectStorageConfigToProto(e))
	}
	return res
}

func FileObjectStorageConfigFromProto(m *pb.FileObjectStorageConfig) main_entity.FileObjectStorageConfig {
	if m == nil {
		return main_entity.FileObjectStorageConfig{}
	}
	return main_entity.FileObjectStorageConfig{
		ObjectStoreUUID: uuid.FromStringOrNil(m.GetObjectStoreUuid()),
	}
}

func FileObjectStorageConfigSliceFromProto(es []*pb.FileObjectStorageConfig) []main_entity.FileObjectStorageConfig {
	if es == nil {
		return []main_entity.FileObjectStorageConfig{}
	}
	res := []main_entity.FileObjectStorageConfig{}
	for _, e := range es {
		res = append(res, FileObjectStorageConfigFromProto(e))
	}
	return res
}
