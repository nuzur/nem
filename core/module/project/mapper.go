package project

import (
	main_entity "github.com/nuzur/nem/core/entity/project"
	nemdb "github.com/nuzur/nem/core/repository/gen"

	"github.com/nuzur/nem/core/entity/mapper"

	"github.com/gofrs/uuid"
	"github.com/nuzur/nem/core/entity/project_extension"
)

func mapModelsToEntities(models []nemdb.Project) []main_entity.Project {
	result := []main_entity.Project{}
	for _, p := range models {
		result = append(result, mapModelToEntity(p))
	}
	return result
}

func mapModelToEntity(model nemdb.Project) main_entity.Project {
	return main_entity.Project{
		UUID:              uuid.FromStringOrNil(model.UUID),
		Version:           model.Version,
		Name:              model.Name,
		Description:       model.Description,
		Tags:              mapper.JSONToStringSlice(model.Tags),
		URL:               model.URL,
		OwnerUUID:         uuid.FromStringOrNil(model.OwnerUUID),
		TeamUUID:          uuid.FromStringOrNil(model.TeamUUID),
		ProjectExtensions: project_extension.ProjectExtensionSliceFromJSON(model.ProjectExtensions),
		Status:            main_entity.Status(model.Status),
		CreatedAt:         model.CreatedAt,
		UpdatedAt:         model.UpdatedAt,
		CreatedByUUID:     uuid.FromStringOrNil(model.CreatedByUUID),
		UpdatedByUUID:     uuid.FromStringOrNil(model.UpdatedByUUID),
	}
}
