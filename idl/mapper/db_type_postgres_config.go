package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/db_type_postgres_config"
	pb "github.com/nuzur/nem/idl/gen"
)

func DbTypePostgresConfigToProto(e main_entity.DbTypePostgresConfig) *pb.DbTypePostgresConfig {
	return &pb.DbTypePostgresConfig{
		Database: e.Database,
		Sslmode:  pb.DbTypePostgresConfigSslmode(e.Sslmode),
	}
}

func DbTypePostgresConfigSliceToProto(es []main_entity.DbTypePostgresConfig) []*pb.DbTypePostgresConfig {
	res := []*pb.DbTypePostgresConfig{}
	for _, e := range es {
		res = append(res, DbTypePostgresConfigToProto(e))
	}
	return res
}

func DbTypePostgresConfigFromProto(m *pb.DbTypePostgresConfig) main_entity.DbTypePostgresConfig {
	if m == nil {
		return main_entity.DbTypePostgresConfig{}
	}
	return main_entity.DbTypePostgresConfig{
		Database: m.GetDatabase(),
		Sslmode:  main_entity.Sslmode(m.GetSslmode()),
	}
}

func DbTypePostgresConfigSliceFromProto(es []*pb.DbTypePostgresConfig) []main_entity.DbTypePostgresConfig {
	if es == nil {
		return []main_entity.DbTypePostgresConfig{}
	}
	res := []main_entity.DbTypePostgresConfig{}
	for _, e := range es {
		res = append(res, DbTypePostgresConfigFromProto(e))
	}
	return res
}
