package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/db_type_mysql_config"
	pb "github.com/nuzur/nem/idl/gen"
)

func DbTypeMysqlConfigToProto(e main_entity.DbTypeMysqlConfig) *pb.DbTypeMysqlConfig {
	return &pb.DbTypeMysqlConfig{
		Params: StringPtrToString(e.Params),
	}
}

func DbTypeMysqlConfigSliceToProto(es []main_entity.DbTypeMysqlConfig) []*pb.DbTypeMysqlConfig {
	res := []*pb.DbTypeMysqlConfig{}
	for _, e := range es {
		res = append(res, DbTypeMysqlConfigToProto(e))
	}
	return res
}

func DbTypeMysqlConfigFromProto(m *pb.DbTypeMysqlConfig) main_entity.DbTypeMysqlConfig {
	if m == nil {
		return main_entity.DbTypeMysqlConfig{}
	}
	return main_entity.DbTypeMysqlConfig{
		Params: &m.Params,
	}
}

func DbTypeMysqlConfigSliceFromProto(es []*pb.DbTypeMysqlConfig) []main_entity.DbTypeMysqlConfig {
	if es == nil {
		return []main_entity.DbTypeMysqlConfig{}
	}
	res := []main_entity.DbTypeMysqlConfig{}
	for _, e := range es {
		res = append(res, DbTypeMysqlConfigFromProto(e))
	}
	return res
}
