package extension_execution

import (
	main_entity "nem/core/entity/extension_execution"
	nemdb "nem/core/repository/gen"

	"github.com/gofrs/uuid"
)

func mapModelsToEntities(models []nemdb.ExtensionExecution) []main_entity.ExtensionExecution {
	result := []main_entity.ExtensionExecution{}
	for _, p := range models {
		result = append(result, mapModelToEntity(p))
	}
	return result
}

func mapModelToEntity(model nemdb.ExtensionExecution) main_entity.ExtensionExecution {
	return main_entity.ExtensionExecution{
		UUID:                 uuid.FromStringOrNil(model.UUID),
		ExtensionUUID:        uuid.FromStringOrNil(model.ExtensionUUID),
		ExtensionVersionUUID: uuid.FromStringOrNil(model.ExtensionVersionUUID),
		ProjectExtensionUUID: uuid.FromStringOrNil(model.ProjectExtensionUUID),
		ProjectUUID:          uuid.FromStringOrNil(model.ProjectUUID),
		ProjectVersionUUID:   uuid.FromStringOrNil(model.ProjectVersionUUID),
		ExecutedByUUID:       uuid.FromStringOrNil(model.ExecutedByUUID),
		Metadata:             string(model.Metadata),
		Status:               main_entity.Status(model.Status),
		StatusMsg:            model.StatusMsg,
		CreatedAt:            model.CreatedAt,
		UpdatedAt:            model.UpdatedAt,
	}
}
