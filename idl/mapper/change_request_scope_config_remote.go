package mapper

import (
	main_entity "github.com/nuzur/nem/core/entity/change_request_scope_config_remote"
	pb "github.com/nuzur/nem/idl/gen"

	"github.com/gofrs/uuid"
)

func ChangeRequestScopeConfigRemoteToProto(e main_entity.ChangeRequestScopeConfigRemote) *pb.ChangeRequestScopeConfigRemote {
	return &pb.ChangeRequestScopeConfigRemote{
		TeamUuid:       e.TeamUUID.String(),
		StoreUuid:      e.StoreUUID.String(),
		ConnectionUuid: e.ConnectionUUID.String(),
	}
}

func ChangeRequestScopeConfigRemoteSliceToProto(es []main_entity.ChangeRequestScopeConfigRemote) []*pb.ChangeRequestScopeConfigRemote {
	res := []*pb.ChangeRequestScopeConfigRemote{}
	for _, e := range es {
		res = append(res, ChangeRequestScopeConfigRemoteToProto(e))
	}
	return res
}

func ChangeRequestScopeConfigRemoteFromProto(m *pb.ChangeRequestScopeConfigRemote) main_entity.ChangeRequestScopeConfigRemote {
	if m == nil {
		return main_entity.ChangeRequestScopeConfigRemote{}
	}
	return main_entity.ChangeRequestScopeConfigRemote{
		TeamUUID:       uuid.FromStringOrNil(m.GetTeamUuid()),
		StoreUUID:      uuid.FromStringOrNil(m.GetStoreUuid()),
		ConnectionUUID: uuid.FromStringOrNil(m.GetConnectionUuid()),
	}
}

func ChangeRequestScopeConfigRemoteSliceFromProto(es []*pb.ChangeRequestScopeConfigRemote) []main_entity.ChangeRequestScopeConfigRemote {
	if es == nil {
		return []main_entity.ChangeRequestScopeConfigRemote{}
	}
	res := []main_entity.ChangeRequestScopeConfigRemote{}
	for _, e := range es {
		res = append(res, ChangeRequestScopeConfigRemoteFromProto(e))
	}
	return res
}
