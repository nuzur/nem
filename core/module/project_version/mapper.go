package project_version

import (
	main_entity "github.com/nuzur/nem/core/entity/project_version"
	nemdb "github.com/nuzur/nem/core/repository/gen"

	"github.com/nuzur/nem/core/entity/mapper"

	"github.com/gofrs/uuid"
	"github.com/nuzur/nem/core/entity/entity"
	"github.com/nuzur/nem/core/entity/enum"
	"github.com/nuzur/nem/core/entity/project_version_deployment"
	"github.com/nuzur/nem/core/entity/project_version_review"
	"github.com/nuzur/nem/core/entity/relationship"
	"github.com/nuzur/nem/core/entity/service"
)

func mapModelsToEntities(models []nemdb.ProjectVersion) []main_entity.ProjectVersion {
	result := []main_entity.ProjectVersion{}
	for _, p := range models {
		result = append(result, mapModelToEntity(p))
	}
	return result
}

func mapModelToEntity(model nemdb.ProjectVersion) main_entity.ProjectVersion {
	return main_entity.ProjectVersion{
		UUID:            uuid.FromStringOrNil(model.UUID),
		Version:         model.Version,
		Identifier:      model.Identifier,
		Description:     model.Description,
		ProjectUUID:     uuid.FromStringOrNil(model.ProjectUUID),
		Entities:        entity.EntitySliceFromJSON(model.Entities),
		Relationships:   relationship.RelationshipSliceFromJSON(model.Relationships),
		Enums:           enum.EnumSliceFromJSON(model.Enums),
		Services:        service.ServiceSliceFromJSON(model.Services),
		BaseVersionUUID: uuid.FromStringOrNil(mapper.SqlNullStringToString(model.BaseVersionUUID)),
		ReviewStatus:    main_entity.ReviewStatus(model.ReviewStatus),
		Reviews:         project_version_review.ProjectVersionReviewSliceFromJSON(model.Reviews),
		Deployments:     project_version_deployment.ProjectVersionDeploymentSliceFromJSON(model.Deployments),
		Status:          main_entity.Status(model.Status),
		CreatedAt:       model.CreatedAt,
		UpdatedAt:       model.UpdatedAt,
		CreatedByUUID:   uuid.FromStringOrNil(model.CreatedByUUID),
		UpdatedByUUID:   uuid.FromStringOrNil(model.UpdatedByUUID),
	}
}
