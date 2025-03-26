package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/data_change_type_config"
	pb "github.com/nuzur/nem/idl/gen"
)

func DataChangeTypeConfigToProto(e main_entity.DataChangeTypeConfig) *pb.DataChangeTypeConfig {
	return &pb.DataChangeTypeConfig{
		Update: DataChangeTypeConfigUpdateToProto(e.Update),
		Create: DataChangeTypeConfigCreateToProto(e.Create),
		Delete: DataChangeTypeConfigCreateToProto(e.Delete),
	}
}

func DataChangeTypeConfigSliceToProto(es []main_entity.DataChangeTypeConfig) []*pb.DataChangeTypeConfig {
	res := []*pb.DataChangeTypeConfig{}
	for _, e := range es {
		res = append(res, DataChangeTypeConfigToProto(e))
	}
	return res
}

func DataChangeTypeConfigFromProto(m *pb.DataChangeTypeConfig) main_entity.DataChangeTypeConfig {
	if m == nil {
		return main_entity.DataChangeTypeConfig{}
	}
	return main_entity.DataChangeTypeConfig{
		Update: DataChangeTypeConfigUpdateFromProto(m.GetUpdate()),
		Create: DataChangeTypeConfigCreateFromProto(m.GetCreate()),
		Delete: DataChangeTypeConfigCreateFromProto(m.GetDelete()),
	}
}

func DataChangeTypeConfigSliceFromProto(es []*pb.DataChangeTypeConfig) []main_entity.DataChangeTypeConfig {
	if es == nil {
		return []main_entity.DataChangeTypeConfig{}
	}
	res := []main_entity.DataChangeTypeConfig{}
	for _, e := range es {
		res = append(res, DataChangeTypeConfigFromProto(e))
	}
	return res
}
