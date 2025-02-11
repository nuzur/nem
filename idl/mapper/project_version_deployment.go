package mapper

import (
	main_entity "nem/core/entity/project_version_deployment"
	pb "nem/idl/gen"

	"github.com/gofrs/uuid"
)

func ProjectVersionDeploymentToProto(e main_entity.ProjectVersionDeployment) *pb.ProjectVersionDeployment {
	return &pb.ProjectVersionDeployment{
		EnviromentUuid: e.EnviromentUUID.String(),
		Status:         pb.ProjectVersionDeploymentStatus(e.Status),
	}
}

func ProjectVersionDeploymentSliceToProto(es []main_entity.ProjectVersionDeployment) []*pb.ProjectVersionDeployment {
	res := []*pb.ProjectVersionDeployment{}
	for _, e := range es {
		res = append(res, ProjectVersionDeploymentToProto(e))
	}
	return res
}

func ProjectVersionDeploymentFromProto(m *pb.ProjectVersionDeployment) main_entity.ProjectVersionDeployment {
	if m == nil {
		return main_entity.ProjectVersionDeployment{}
	}
	return main_entity.ProjectVersionDeployment{
		EnviromentUUID: uuid.FromStringOrNil(m.GetEnviromentUuid()),
		Status:         main_entity.Status(m.GetStatus()),
	}
}

func ProjectVersionDeploymentSliceFromProto(es []*pb.ProjectVersionDeployment) []main_entity.ProjectVersionDeployment {
	if es == nil {
		return []main_entity.ProjectVersionDeployment{}
	}
	res := []main_entity.ProjectVersionDeployment{}
	for _, e := range es {
		res = append(res, ProjectVersionDeploymentFromProto(e))
	}
	return res
}
