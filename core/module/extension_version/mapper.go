package extension_version

import (
	main_entity "github.com/nuzur/nem/core/entity/extension_version"
	nemdb "github.com/nuzur/nem/core/repository/gen"

	"github.com/nuzur/nem/core/entity/mapper"

	"github.com/gofrs/uuid"
)

func mapModelsToEntities(models []nemdb.ExtensionVersion) []main_entity.ExtensionVersion {
	result := []main_entity.ExtensionVersion{}
	for _, p := range models {
		result = append(result, mapModelToEntity(p))
	}
	return result
}

func mapModelToEntity(model nemdb.ExtensionVersion) main_entity.ExtensionVersion {
	return main_entity.ExtensionVersion{
		UUID:                uuid.FromStringOrNil(model.UUID),
		Version:             model.Version,
		ExtensionUUID:       uuid.FromStringOrNil(model.ExtensionUUID),
		DisplayVersion:      mapper.SqlNullStringToStringPtr(model.DisplayVersion),
		Description:         mapper.SqlNullStringToStringPtr(model.Description),
		RepositoryTag:       model.RepositoryTag,
		ConfigurationEntity: mapper.StringToStringPtr(string(model.ConfigurationEntity)),
		ExecutionMode:       main_entity.JSONToExecutionModeSlice(model.ExecutionMode),
		ReviewStatus:        main_entity.ReviewStatus(model.ReviewStatus),
		Status:              main_entity.Status(model.Status),
		CreatedAt:           model.CreatedAt,
		UpdatedAt:           model.UpdatedAt,
		CreatedByUUID:       uuid.FromStringOrNil(model.CreatedByUUID),
		UpdatedByUUID:       uuid.FromStringOrNil(model.UpdatedByUUID),
	}
}
