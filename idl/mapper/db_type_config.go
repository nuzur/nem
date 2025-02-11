package mapper

import (
	main_entity "nem/core/entity/db_type_config"
	pb "nem/idl/gen"
)

func DbTypeConfigToProto(e main_entity.DbTypeConfig) *pb.DbTypeConfig {
	return &pb.DbTypeConfig{
		Postgres: DbTypePostgresConfigToProto(e.Postgres),
		Mysql:    DbTypeMysqlConfigToProto(e.Mysql),
	}
}

func DbTypeConfigSliceToProto(es []main_entity.DbTypeConfig) []*pb.DbTypeConfig {
	res := []*pb.DbTypeConfig{}
	for _, e := range es {
		res = append(res, DbTypeConfigToProto(e))
	}
	return res
}

func DbTypeConfigFromProto(m *pb.DbTypeConfig) main_entity.DbTypeConfig {
	if m == nil {
		return main_entity.DbTypeConfig{}
	}
	return main_entity.DbTypeConfig{
		Postgres: DbTypePostgresConfigFromProto(m.GetPostgres()),
		Mysql:    DbTypeMysqlConfigFromProto(m.GetMysql()),
	}
}

func DbTypeConfigSliceFromProto(es []*pb.DbTypeConfig) []main_entity.DbTypeConfig {
	if es == nil {
		return []main_entity.DbTypeConfig{}
	}
	res := []main_entity.DbTypeConfig{}
	for _, e := range es {
		res = append(res, DbTypeConfigFromProto(e))
	}
	return res
}
