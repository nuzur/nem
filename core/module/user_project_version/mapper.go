package user_project_version

import (
	main_entity "github.com/nuzur/nem/core/entity/user_project_version"
	nemdb "github.com/nuzur/nem/core/repository/gen"

	"github.com/nuzur/nem/core/entity/mapper"

	"github.com/gofrs/uuid"
)

func mapModelsToEntities(models []nemdb.UserProjectVersion) []main_entity.UserProjectVersion {
	result := []main_entity.UserProjectVersion{}
	for _, p := range models {
		result = append(result, mapModelToEntity(p))
	}
	return result
}

func mapModelToEntity(model nemdb.UserProjectVersion) main_entity.UserProjectVersion {
	return main_entity.UserProjectVersion{
		UUID:               uuid.FromStringOrNil(model.UUID),
		Version:            model.Version,
		ProjectVersionUUID: uuid.FromStringOrNil(model.ProjectVersionUUID),
		UserUUID:           uuid.FromStringOrNil(model.UserUUID),
		Data:               mapper.StringToStringPtr(string(model.Data)),
		Status:             main_entity.Status(model.Status),
		CreatedAt:          model.CreatedAt,
		UpdatedAt:          model.UpdatedAt,
		CreatedByUUID:      uuid.FromStringOrNil(model.CreatedByUUID),
		UpdatedByUUID:      uuid.FromStringOrNil(model.UpdatedByUUID),
	}
}
