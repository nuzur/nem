package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/local_agent_connection"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"
)

func LocalAgentConnectionToProto(e main_entity.LocalAgentConnection) *pb.LocalAgentConnection {
	return &pb.LocalAgentConnection{
		Uuid:          e.UUID.String(),
		Name:          e.Name,
		DbType:        pb.LocalAgentConnectionDbType(e.DbType),
		DefaultSchema: e.DefaultSchema,
	}
}

func LocalAgentConnectionSliceToProto(es []main_entity.LocalAgentConnection) []*pb.LocalAgentConnection {
	res := []*pb.LocalAgentConnection{}
	for _, e := range es {
		res = append(res, LocalAgentConnectionToProto(e))
	}
	return res
}

func LocalAgentConnectionFromProto(m *pb.LocalAgentConnection) main_entity.LocalAgentConnection {
	if m == nil {
		return main_entity.LocalAgentConnection{}
	}
	return main_entity.LocalAgentConnection{
		UUID:          uuid.FromStringOrNil(m.GetUuid()),
		Name:          m.GetName(),
		DbType:        main_entity.DbType(m.GetDbType()),
		DefaultSchema: m.GetDefaultSchema(),
	}
}

func LocalAgentConnectionSliceFromProto(es []*pb.LocalAgentConnection) []main_entity.LocalAgentConnection {
	if es == nil {
		return []main_entity.LocalAgentConnection{}
	}
	res := []main_entity.LocalAgentConnection{}
	for _, e := range es {
		res = append(res, LocalAgentConnectionFromProto(e))
	}
	return res
}
