package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/field_type_encrypted_config"
	pb "github.com/nuzur/nem/idl/gen"
)

func FieldTypeEncryptedConfigToProto(e main_entity.FieldTypeEncryptedConfig) *pb.FieldTypeEncryptedConfig {
	return &pb.FieldTypeEncryptedConfig{
		MinSize:         int64(e.MinSize),
		MaxSize:         int64(e.MaxSize),
		RegexValidation: e.RegexValidation,
		EncryptionType:  pb.FieldTypeEncryptedConfigEncryptionType(e.EncryptionType),
		UseSalt:         e.UseSalt,
		SaltFieldUuids:  UUIDSliceToStringSlice(e.SaltFieldUUIDs),
	}
}

func FieldTypeEncryptedConfigSliceToProto(es []main_entity.FieldTypeEncryptedConfig) []*pb.FieldTypeEncryptedConfig {
	res := []*pb.FieldTypeEncryptedConfig{}
	for _, e := range es {
		res = append(res, FieldTypeEncryptedConfigToProto(e))
	}
	return res
}

func FieldTypeEncryptedConfigFromProto(m *pb.FieldTypeEncryptedConfig) main_entity.FieldTypeEncryptedConfig {
	if m == nil {
		return main_entity.FieldTypeEncryptedConfig{}
	}
	return main_entity.FieldTypeEncryptedConfig{
		MinSize:         int64(m.GetMinSize()),
		MaxSize:         int64(m.GetMaxSize()),
		RegexValidation: m.GetRegexValidation(),
		EncryptionType:  main_entity.EncryptionType(m.GetEncryptionType()),
		UseSalt:         m.GetUseSalt(),
		SaltFieldUUIDs:  StringSliceToUUIDSlice(m.GetSaltFieldUuids()),
	}
}

func FieldTypeEncryptedConfigSliceFromProto(es []*pb.FieldTypeEncryptedConfig) []main_entity.FieldTypeEncryptedConfig {
	if es == nil {
		return []main_entity.FieldTypeEncryptedConfig{}
	}
	res := []main_entity.FieldTypeEncryptedConfig{}
	for _, e := range es {
		res = append(res, FieldTypeEncryptedConfigFromProto(e))
	}
	return res
}
