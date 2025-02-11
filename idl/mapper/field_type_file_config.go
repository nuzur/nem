package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/field_type_file_config"
	pb "github.com/nuzur/nem/idl/gen"
)

func FieldTypeFileConfigToProto(e main_entity.FieldTypeFileConfig) *pb.FieldTypeFileConfig {
	return &pb.FieldTypeFileConfig{
		StorageType:       pb.FieldTypeFileConfigStorageType(e.StorageType),
		StorageConfig:     FileStorageConfigToProto(e.StorageConfig),
		AllowedExtensions: e.AllowedExtensions,
		MaxSize:           int64(e.MaxSize),
		AllowMultiple:     e.AllowMultiple,
		MaxFiles:          int64(e.MaxFiles),
	}
}

func FieldTypeFileConfigSliceToProto(es []main_entity.FieldTypeFileConfig) []*pb.FieldTypeFileConfig {
	res := []*pb.FieldTypeFileConfig{}
	for _, e := range es {
		res = append(res, FieldTypeFileConfigToProto(e))
	}
	return res
}

func FieldTypeFileConfigFromProto(m *pb.FieldTypeFileConfig) main_entity.FieldTypeFileConfig {
	if m == nil {
		return main_entity.FieldTypeFileConfig{}
	}
	return main_entity.FieldTypeFileConfig{
		StorageType:       main_entity.StorageType(m.GetStorageType()),
		StorageConfig:     FileStorageConfigFromProto(m.GetStorageConfig()),
		AllowedExtensions: m.GetAllowedExtensions(),
		MaxSize:           int64(m.GetMaxSize()),
		AllowMultiple:     m.GetAllowMultiple(),
		MaxFiles:          int64(m.GetMaxFiles()),
	}
}

func FieldTypeFileConfigSliceFromProto(es []*pb.FieldTypeFileConfig) []main_entity.FieldTypeFileConfig {
	if es == nil {
		return []main_entity.FieldTypeFileConfig{}
	}
	res := []main_entity.FieldTypeFileConfig{}
	for _, e := range es {
		res = append(res, FieldTypeFileConfigFromProto(e))
	}
	return res
}
